#include "opencl_engine.h"
#include <CL/cl.h>
int init_gpu(opencl_engine *engine) {
    if (engine == 0) return -1;
    cl_int err = clGetPlatformIDs(1, &engine->platform_id, 0);
    if (err != CL_SUCCESS) return -1;
    err = clGetDeviceIDs(engine->platform_id, CL_DEVICE_TYPE_GPU, 1, &engine->device_id, 0);
    if (err != CL_SUCCESS) return -1;
     engine->context = clCreateContext(0, 1, &engine->device_id, 0, 0, &err);
    if (err != CL_SUCCESS) return -1;
    engine->queue = clCreateCommandQueueWithProperties(engine->context, engine->device_id, 0, &err);
    if (err != CL_SUCCESS) return -1;
    return 0;
}
int process_pixels() {
    return 0;
}