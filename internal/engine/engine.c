#include "engine.h"
#include <stdio.h>
#include <stdlib.h>
#include "opencl_engine.h"
#include "openvino_engine.h"
#include "openvino/c/openvino.h"

opencl_engine *create_engine(void) {
    opencl_engine *engine = malloc(sizeof(opencl_engine));
    if (engine == NULL) {
        printf("Failed to allocate engine\n");
        exit(-1);
    }
    int err = init_gpu(engine);
    if (err != 0) return NULL;
    return engine;
}

int revert_color(opencl_engine *engine, unsigned char *pixels, int size) {
    int err = send_pixels_gpu(engine, pixels, size);
    if (err != 0) {
        return 1;
    }
    return 0;
}

openvino_engine *create_openvino_engine(void) {
    int err;
    openvino_engine *engine = malloc(sizeof(openvino_engine));
    if (engine == NULL) {
        printf("Failed to allocate engine\n");
        exit(-1);
    }
    err = init_openvino(engine);
    if (err != 0) return NULL;
    err = load_model(engine);
    if (err != 0) return NULL;
    err = run_model(engine);
    if (err != 0) return NULL;
    err = inference(engine);
    if (err != 0) return NULL;

    return engine;
}

static float clamp01(float x) {
    if (x < 0) return 0;
    if (x > 1) return 1;
    return x;
}

// RGBA 8 bits entrelacé -> RGB float par couleur, normalisé 0-1
static void preprocess(unsigned char *pixels, float *input, int pixel_count) {
    for (int p = 0; p < pixel_count; p++) {
        input[0 * pixel_count + p] = pixels[p * 4 + 0] / 255.0f; // R
        input[1 * pixel_count + p] = pixels[p * 4 + 1] / 255.0f; // G
        input[2 * pixel_count + p] = pixels[p * 4 + 2] / 255.0f; // B
    }
}

// RGB float par couleur 0-1 -> RGBA 8 bits entrelacé, réécrit dans pixels
static void postprocess(float *output, unsigned char *pixels, int pixel_count) {
    for (int p = 0; p < pixel_count; p++) {
        pixels[p * 4 + 0] = (unsigned char)(clamp01(output[0 * pixel_count + p]) * 255.0f);
        pixels[p * 4 + 1] = (unsigned char)(clamp01(output[1 * pixel_count + p]) * 255.0f);
        pixels[p * 4 + 2] = (unsigned char)(clamp01(output[2 * pixel_count + p]) * 255.0f);
        // l'alpha (p*4+3) reste inchangé
    }
}

int contrast_color(openvino_engine *engine, unsigned char *pixels, int size, int height, int width) {
    int pixel_count = width * height;
    int rc = -1;
    int shape_created = 0;
    float *input = NULL;
    ov_tensor_t *input_tensor = NULL;
    ov_tensor_t *output_tensor = NULL;
    void *out_data = NULL;
    ov_shape_t shape;
    int64_t dims[] = {1, 3, height, width};

    if (ov_shape_create(4, dims, &shape) != 0) return -1;
    shape_created = 1;

    input = malloc(3 * pixel_count * sizeof(float));
    if (input == NULL) goto cleanup;
    preprocess(pixels, input, pixel_count);

    if (ov_tensor_create_from_host_ptr(F32, shape, input, &input_tensor) != 0) goto cleanup;
    if (ov_infer_request_set_input_tensor(engine->infer_request, input_tensor) != 0) goto cleanup;
    if (ov_infer_request_infer(engine->infer_request) != 0) goto cleanup;
    if (ov_infer_request_get_output_tensor(engine->infer_request, &output_tensor) != 0) goto cleanup;
    if (ov_tensor_data(output_tensor, &out_data) != 0) goto cleanup;

    postprocess((float *)out_data, pixels, pixel_count);
    rc = 0;

cleanup:
    if (output_tensor) ov_tensor_free(output_tensor);
    if (input_tensor) ov_tensor_free(input_tensor);
    if (shape_created) ov_shape_free(&shape);
    if (input) free(input);
    return rc;
}
