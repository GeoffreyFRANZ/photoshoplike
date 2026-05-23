#include "engine.h"
#include <stdio.h>
#include <stdlib.h>
#include "opencl_engine.h"

int max(int max, int value) {
    if (max < value) {
        return value;
    }
    return max;
}

int min(int min, int value) {
    if (min > value) {
        return value;
    }
    return min;
}

void get_img(unsigned char *pixels, int size) {
    int i = 0;
    int minR = 255;
    int minG = 255;
    int minB = 255;
    int maxR = 0;
    int maxG = 0;
    int maxB = 0;

    while (i < size) {
        if (i % 4 == 0) {
            minR = min(minR, pixels[i]);
            maxR = max(maxR, pixels[i]);
        }
        if (i % 4 == 1) {
            minG = min(minG, pixels[i]);
            maxG = max(maxG, pixels[i]);
        }
        if (i % 4 == 2) {
            minB = min(minB, pixels[i]);
            maxB = max(maxB, pixels[i]);
        }
        i++;
    }
    opencl_engine *engine = malloc(sizeof(opencl_engine));
    if (engine == NULL) {
        printf("Failed to allocate engine\n");
        return;
    }
    int err = init_gpu(engine);
    if (err != 0) printf("bad initializing GPU");
    err = process_pixels(engine, pixels, size);
    if (err != 0) printf("bad processing pixels");
    free(engine);
    fflush(stdout);
}
    