//
// Created by Franz on 17/05/2026.
//

#ifndef SENTINEL_GO_OPENCL_ENGINE_H
#define SENTINEL_GO_OPENCL_ENGINE_H
#include <CL/cl.h>

typedef struct opencl_engine {
    cl_platform_id platform_id;
    cl_device_id device_id;
    cl_context context;
}opencl_engine;

int init_gpu(opencl_engine *engine);
#endif //SENTINEL_GO_OPENCL_ENGINE_H

