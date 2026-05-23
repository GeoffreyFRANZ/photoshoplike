# Photoshop-like — Roadmap

## ✅ Done
- Go HTTP server + image upload
- Image decoding (JPEG/PNG) + pixel extraction
- CGo bridge — Go → C zero-copy memory
- OpenCL GPU init (platform, device, context, queue)
- GPU buffer creation + pixel transfer
- Pixels read back from GPU
- PNG response to browser
- Image displays in browser

## 🔧 In Progress
- HTML — fetch form submit (no page reload)
- Display result image in browser without redirect

## 📋 Next
1. Invert filter — OpenCL kernel (`255 - R, 255 - G, 255 - B`)
2. Toggle button (normal ↔ inverted)
3. New route `/reverse` in Go
4. release_gpu() — free OpenCL resources

## 🚀 Later
- More filters (grayscale, brightness, blur)
- AI model — highlight/shadow/contrast detection
- Clean UI
- Docker
- README update when POC done