#include "opencl_engine.h"
#include <stdio.h>
#include <stdlib.h>
#include <stddef.h>
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
int process_pixels(opencl_engine *engine, unsigned char *pixels, int size) {
    int err;
    engine->mem = clCreateBuffer(engine->context, CL_MEM_READ_WRITE | CL_MEM_COPY_HOST_PTR , size, pixels, &err);
    if (err != CL_SUCCESS) return -1;
    err = clEnqueueReadBuffer(engine->queue, engine->mem, CL_TRUE, 0, size, pixels, 0, 0, 0);
    if (err != CL_SUCCESS) return -1;
    return 0;
}

int send_pixels_gpu(opencl_engine *engine, unsigned char *pixels, int size) {
    int err;
    FILE *file = fopen("kernel.cl", "r");
    size_t global_size = size;
    if (file == NULL) return -1;
    fseek(file, 0, SEEK_END);
    int file_size = ftell(file);
    fseek(file, 0, SEEK_SET);
    char *buffer = malloc(file_size);
    fread(buffer, 1, file_size, file);
    cl_program program = clCreateProgramWithSource(engine->context, 1, (const char **)&buffer, 0, &err);
    if (err != CL_SUCCESS) return -1;
    cl_int building = clBuildProgram(program, 1, &engine->device_id, 0, 0, 0);
    cl_kernel kernel = clCreateKernel(program, "revert_pixels", &err);
    if (err != CL_SUCCESS) return -1;
    clSetKernelArg(kernel, 0, sizeof(engine->mem), &engine->mem);
    clSetKernelArg(kernel, 1, sizeof(int), &size);
    clEnqueueNDRangeKernel(engine->queue, kernel, 1, 0, &global_size, 0, 0, 0, 0);
    clEnqueueReadBuffer(engine->queue, engine->mem, CL_TRUE, 0, size, pixels, 0, 0, 0);
    fclose(file);
    free(buffer);
    return 0;
}