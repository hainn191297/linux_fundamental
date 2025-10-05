# Toolchain — Môi trường và công cụ

Yêu cầu tối thiểu:

- Go 1.21+ (khuyến nghị): Sử dụng `go version` để kiểm tra.
- Hệ điều hành: Linux (ưu tiên) hoặc macOS cho phát triển.
- Công cụ quan sát syscall: `strace` (Linux), `dtruss`/`ktrace` (macOS).
- Tiện ích cơ bản: `git`, `make` (tuỳ chọn), `nc`/`netcat` để test mạng.

Cài đặt nhanh (Ubuntu/Debian):

```bash

sudo apt update
sudo apt install -y build-essential strace netcat-openbsd
# Cài Go từ tarball hoặc dùng gimme/gvm tuỳ môi trường
```

Gợi ý trên macOS:

- Cài Xcode CLT: `xcode-select --install`
- Homebrew: `brew install go` và `brew install netcat`
- Dùng `dtruss sudo -n -f -t openat ls` để xem syscall tương đương (yêu cầu quyền).

Lưu ý khác:

- `strace` hữu ích nhất trên Linux; hành vi Go runtime (epoll/futex) biểu hiện đầy đủ trên Linux.
- Một số ví dụ (epoll) chạy trên macOS vẫn OK vì Go dùng kqueue phía dưới, nhưng tên syscall khác.
