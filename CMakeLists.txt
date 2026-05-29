cmake_minimum_required(VERSION 4.2)
project(sentinel_go C)

set(CMAKE_C_STANDARD 11)

include_directories(internal/engine)

add_executable(sentinel_go
        internal/engine/engine.c
        internal/engine/engine.h
        internal/engine/opencl_engine.c
        internal/engine/opencl_engine.h
        internal/engine/kernel.c)
