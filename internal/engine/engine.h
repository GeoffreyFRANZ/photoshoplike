#ifndef ENGINE_H
#define ENGINE_H
typedef struct opencl_engine opencl_engine;
opencl_engine *create_engine(void);
void revert_color(opencl_engine *engine, unsigned char *pixels, int size);
#endif