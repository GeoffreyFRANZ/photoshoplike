//
// Created by Franz on 20/06/2026.
//

#ifndef SENTINEL_GO_OPENVINO_ENGINE_H
#define SENTINEL_GO_OPENVINO_ENGINE_H
#include "openvino/c/ov_core.h"
typedef struct openvino_engine {
    ov_core_t* core;
    ov_model_t* model;
    ov_compiled_model_t* compiled_model;
    ov_infer_request_t* infer_request;
    ov_tensor_t* tensor;
} openvino_engine;
int init_openvino(openvino_engine *engine);
int load_model(openvino_engine *engine);
int run_model(openvino_engine *engine);
int inference(openvino_engine *engine);
#endif //SENTINEL_GO_OPENVINO_ENGINE_H
