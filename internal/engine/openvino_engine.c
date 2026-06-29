#include <stdlib.h>

#include "openvino/c/openvino.h"
#include "openvino_engine.h"

int init_openvino(openvino_engine *engine) {
    int err = ov_core_create(&engine->core);
    if (err != 0) return -1;
    return 0;
}

int load_model(openvino_engine *engine) {
    if (ov_core_read_model(engine->core, "internal/model/zero_dce.onnx", NULL, &engine->model) != 0) return -1;
    return 0;
}
int run_model(openvino_engine *engine) {
    if (ov_core_compile_model(engine->core, engine->model, "AUTO", 0, &engine->compiled_model) != 0) return -1;
    return 0;
}
int inference(openvino_engine *engine) {
    if (ov_compiled_model_create_infer_request(engine->compiled_model, &engine->infer_request) != 0) return -1;
    return 0;
}