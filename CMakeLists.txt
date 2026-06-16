cmake_minimum_required(VERSION 3.20)
project(sentinel_go C)

set(CMAKE_C_STANDARD 11)

find_package(OpenCL REQUIRED)

include_directories(internal/engine)

add_library(engine STATIC
        internal/engine/engine.c
        internal/engine/engine.h
        internal/engine/opencl_engine.c
        internal/engine/opencl_engine.h)

target_link_libraries(engine PRIVATE OpenCL::OpenCL)
