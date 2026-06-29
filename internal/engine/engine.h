#ifndef ENGINE_H
#define ENGINE_H
typedef struct opencl_engine opencl_engine;
typedef struct openvino_engine openvino_engine;
openvino_engine *create_openvino_engine(void);
opencl_engine *create_engine(void);
int revert_color(opencl_engine *engine, unsigned char *pixels, int size);
int contrast_color(openvino_engine *engine, unsigned char *pixels, int size, int height, int width);
#endif