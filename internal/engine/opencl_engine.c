//
// Created by Franz on 17/05/2026.
//
#include <CL/opencl.h>
#include "opencl_engine.h"

int init_gpu(opencl_engine *engine) {
    if (engine == NULL) return -2;
    cl_uint num_platforms;
    cl_int err;

    err = clGetPlatformIDs(1, &(engine->platform_id), &num_platforms);
    if (err != CL_SUCCESS || num_platforms == 0) return -1;
    err = clGetDeviceIDs(engine->platform_id, CL_DEVICE_TYPE_GPU, 1, &(engine->device_id), NULL);
    if (err != CL_SUCCESS) return -1;
    return 0;
}

void create_buffer(opencl_engine *engine, int size) {
    cl_int err;
    cl_mem cl_buffer;
    cl_buffer = clCreateBuffer(engine->context, CL_MEM_READ_WRITE, size, NULL, &err)
}
int get_img_gpu(unsigned char *pixels, int size) {

    return 0;
}