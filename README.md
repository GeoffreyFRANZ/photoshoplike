# photoshoplike — GPU Image Processing Engine

> Low-level image processing engine built from scratch in **Go + C + OpenCL**.  
> Zero-copy memory pipeline between Go, C, and GPU — no image library, every pixel handled manually.

![Go](https://img.shields.io/badge/Go-1.21-00ADD8?style=flat&logo=go&logoColor=white)
![C](https://img.shields.io/badge/C-C11-A8B9CC?style=flat&logo=c&logoColor=white)
![OpenCL](https://img.shields.io/badge/OpenCL-3.0-ED1C24?style=flat)
![CGo](https://img.shields.io/badge/CGo-bridge-00ADD8?style=flat)
![Status](https://img.shields.io/badge/status-in%20progress-yellow?style=flat)

## Why this project

Most image processing libraries abstract everything away. I wanted to understand what actually happens when you apply a filter to an image:
- How pixels live in memory (flat RGBA byte arrays)
- How to move data between Go and C without copying (zero-copy via `unsafe.Pointer`)
- How GPU parallelism works at the OpenCL kernel level
- What the host/device memory model looks like in practice

## Architecture

```
Browser upload
    │
    ▼
Go HTTP server (net/http)
    │  raw bytes
    ▼
CGo bridge ── unsafe.Pointer (zero-copy) ──▶ C engine
                                                  │
                                                  ▼
                                          OpenCL 3.0 kernel
                                          (parallel pixel ops on GPU)
                                                  │
                                                  ▼
                                          processed RGBA buffer
    ◀─────────────────────────────────────────────┘
    ▼
PNG response → Browser
```

## Tech stack

| Layer | Tech | Why |
|---|---|---|
| HTTP server | Go `net/http` | Concurrency, goroutines for async uploads |
| Memory bridge | CGo + `unsafe.Pointer` | Zero-copy transfer — no malloc/free overhead at the boundary |
| Pixel engine | C (C11) | Manual memory management, direct hardware access |
| GPU compute | OpenCL 3.0 | Parallel pixel processing — each pixel is independent |

## Key concepts implemented

### Zero-copy CGo bridge
```go
// Go side — pass slice directly to C without copying
ptr := unsafe.Pointer(&pixels[0])
C.process_image((*C.uchar)(ptr), C.int(width), C.int(height))
```

```c
// C side — work directly on Go's memory
void process_image(unsigned char *pixels, int width, int height) {
    // manipulate RGBA pixels in place
}
```

### Image decoding pipeline
- JPEG and PNG decoded to flat `[]byte` (RGBA, 4 bytes/pixel)
- Pixel at `(x, y)` → `pixels[y*width*4 + x*4 : y*width*4 + x*4 + 4]`
- Passed to C engine as raw pointer — no intermediate copy

### OpenCL GPU setup
- Platform and device enumeration
- Context, command queue, and program compilation
- Kernel execution on device memory

## Current state

- [x] Go HTTP server — image upload + response
- [x] Image decoding (JPEG, PNG) → flat RGBA byte array
- [x] CGo bridge — Go → C zero-copy via `unsafe.Pointer`
- [x] OpenCL init — platform, device, context, queue
- [ ] GPU buffer allocation + kernel execution
- [ ] Processed pixels → browser response
- [ ] Filter library (grayscale, blur, edge detection)

## Run locally

Requires OpenCL drivers (`clinfo` to verify).

```bash
git clone https://github.com/GeoffreyFRANZ/photoshoplike
cd photoshoplike
go run .
# → http://localhost:8080
```

## What I'm learning

- **CGo ownership rules** — when Go GC can move memory and when it can't (`runtime.Pinner`)
- **OpenCL memory model** — host vs device buffers, `clEnqueueWriteBuffer`, `clEnqueueNDRangeKernel`
- **GPU parallelism** — why image filters are embarrassingly parallel (each pixel independent)
- **C memory discipline** — `malloc`/`free` lifecycle, avoiding leaks across the CGo boundary
- **Systems programming mindset** — thinking in bytes, pointers, and cache lines

## Related

- [AiLearning](https://github.com/GeoffreyFRANZ/AiLearning) — Neural network from scratch in Python/NumPy
