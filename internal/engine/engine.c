#include "engine.h"
#include <stdio.h>
#include <stdlib.h>
#include "opencl_engine.h"

opencl_engine *create_engine(void) {
     opencl_engine *engine = malloc(sizeof(opencl_engine));
    if (engine == NULL) {
        printf("Failed to allocate engine\n");
        exit(-1);
    }
    int err = init_gpu(engine);
    if (err != 0) exit(-1);
    return engine;

}
 void revert_color(opencl_engine *engine, unsigned char *pixels, int size) {
    int err = send_pixels_gpu(engine, pixels, size);
    if (err != 0) printf("bad sending pixels");
}