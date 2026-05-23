# Photoshop-like — GPU Image Processor

Learning project building an image processing engine from scratch using **Go, C, and OpenCL**.  
Goal: understand low-level systems programming, GPU computing, and how image filters actually work.

> ⚠️ Work in progress — GPU pipeline is being implemented.

## Tech stack

| Layer | Tech | Role |
|---|---|---|
| Web server | Go (`net/http`) | HTTP server, image upload, response |
| Bridge | CGo + `unsafe.Pointer` | Zero-copy memory transfer between Go and C |
| Engine | C | Pixel manipulation, GPU orchestration |
| GPU | OpenCL 3.0 | Parallel pixel processing on GPU |

## Architecture

```
Browser (upload form)
    → Go HTTP server
    → CGo bridge (zero-copy via unsafe.Pointer)
    → C engine
    → OpenCL kernel (GPU)
    → processed pixels back to Go
    → PNG response to browser
```

## Current state

- [x] Go HTTP server with image upload
- [x] Image decoding (JPEG, PNG)
- [x] Pixel extraction to flat RGBA byte array
- [x] CGo bridge — Go → C memory transfer
- [x] OpenCL GPU init (platform, device, context, queue)
- [ ] GPU buffer + kernel execution
- [ ] Processed pixels back to browser

## Run

Requires OpenCL drivers installed (`clinfo` to check).

```bash
cd photoshop-like
go run .
```

Then open `http://localhost:8080`.

## What I'm learning

- How images live in memory (flat RGBA byte arrays)
- Manual memory management in C (`malloc`, `free`)
- CGo — bridging Go and C, ownership rules, GC interactions
- OpenCL — GPU parallelism, host/device memory model, kernels
- Why GPU parallel processing is fast for image operations
