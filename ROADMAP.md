# Photoshop-like — Roadmap

## ✅ Done
- Go HTTP server + image upload
- Image decoding (JPEG) + pixel extraction
- CGo bridge — Go → C zero-copy memory
- OpenCL GPU init (platform, device, context, queue)
- GPU buffer creation + pixel transfer
- Pixels read back from GPU
- JPEG response to browser, image displays without page reload
- Invert filter — OpenCL kernel (`255 - R, 255 - G, 255 - B`) via `/reverse-color`
- GPU-resident engine refactor — dependency injection (engine built once, reused per request)

## 🔧 In Progress
- Thread-safety — protect the shared GPU engine against concurrent requests (`sync.Mutex`)

## 📋 Next
1. Upload hardening — `http.MaxBytesReader`, `image.DecodeConfig` dimension cap (decompression-bomb guard), empty-image guard
2. `create_engine` — return an error instead of `exit(-1)`
3. `release_gpu()` — free OpenCL resources on shutdown
4. Layered package layout (model / session / server)

## 🚀 Later
- More filters (grayscale, brightness, blur, edge detection)
- AI light/contrast correction — OpenVINO inference (CPU/GPU/NPU) via `extern "C"` (v2)
- Clean UI
- Docker
