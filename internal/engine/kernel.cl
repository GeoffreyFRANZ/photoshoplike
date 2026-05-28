//
// Created by Franz on 27/05/2026.
//
__kernel void revert_pixels(__global unsigned char *pixels, int numElements){
    int iGID = get_global_id(0);
    if (iGID >= numElements) {
        return;
    }
    if (iGID % 4 != 3){
        pixels[iGID] = 255 - pixels[iGID];
    }
}