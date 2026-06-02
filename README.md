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
- How to **industrialize** GPU inference behind a web service: create the GPU engine once, keep it resident, and reuse it across every request (dependency injection) instead of rebuilding it per call

## Architecture

```
                          ┌─────────────────────────────┐
  boot once ──▶ engine.New() → create OpenCL engine (GPU-resident,
                          │     injected into the HTTP server)
                          └──────────────┬──────────────┘
                                         │ reused every request
Browser upload                          │
    │                                    │
    ▼                                    ▼
Go HTTP server (net/http) ── Server{engine} (dependency injection)
    │  raw bytes
    ▼
CGo bridge ── unsafe.Pointer (zero-copy) ──▶ C engine (shared)
                                                  │
                                                  ▼
                                          OpenCL 3.0 kernel
                                          (parallel pixel ops on GPU)
                                                  │
                                                  ▼
                                          processed RGBA buffer
    ◀─────────────────────────────────────────────┘
    ▼
JPEG response → Browser
```

## Tech stack

| Layer | Tech | Why |
|---|---|---|
| HTTP server | Go `net/http` | Concurrency, goroutines for async uploads |
| Memory bridge | CGo + `unsafe.Pointer` | Zero-copy transfer — no malloc/free overhead at the boundary |
| Pixel engine | C (C11) | Manual memory management, direct hardware access |
| GPU compute | OpenCL 3.0 | Parallel pixel processing — each pixel is independent |
| Engine lifecycle | Dependency injection | GPU engine built once at boot, injected into the server, reused per request (GPU-resident) |

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

### GPU-resident engine (dependency injection)
The GPU engine is expensive to build (device enumeration, context, queue). Instead of
rebuilding it on every request, it is created **once** at boot and **injected** into the
HTTP server, then reused for the lifetime of the process.

```go
// boot — build the GPU engine once
type Server struct {
    engine *engine.Engine // injected dependency, GPU-resident
}

func main() {
    server := Server{engine.New()} // create_engine() runs a single time
    mux.HandleFunc("/reverse-color", server.reverseColorHandler)
    // ...
}

// per request — reuse the shared engine, no GPU re-setup
func (s *Server) reverseColorHandler(w http.ResponseWriter, r *http.Request) {
    s.engine.RevertColor(unsafe.Pointer(&pixels[0]), size)
}
```

The C side mirrors this: `create_engine()` allocates the OpenCL context, and the filter
functions receive that engine as a parameter instead of allocating/freeing one per call.

## Current state

- [x] Go HTTP server — image upload + response
- [x] Image decoding (JPEG) → flat RGBA byte array
- [x] CGo bridge — Go → C zero-copy via `unsafe.Pointer`
- [x] OpenCL init — platform, device, context, queue
- [x] GPU buffer allocation + kernel execution (invert filter)
- [x] Processed pixels → browser response (JPEG)
- [x] GPU-resident engine via dependency injection
- [ ] Thread-safe shared engine (mutex over the GPU)
- [ ] Upload hardening (size cap, decompression-bomb guard)
- [ ] Filter library (grayscale, brightness, blur, edge detection)
- [ ] AI-based light/contrast correction (OpenVINO, CPU/GPU/NPU) — v2

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
- **Dependency injection** — keeping a costly GPU engine resident and reusing it, instead of rebuilding it per request
- **Concurrency** — protecting a shared GPU engine against concurrent HTTP requests (goroutines + mutex)
- **Systems programming mindset** — thinking in bytes, pointers, and cache lines

## Related

- [AiLearning](https://github.com/GeoffreyFRANZ/AiLearning) — Neural network from scratch in Python/NumPy
