# Network Programming & epoll (TLPI Ch.56–61, 63)

Mục tiêu: hiểu cơ chế socket, non-blocking I/O, và epoll. Trong Go, runtime netpoll dựa trên epoll (Linux) hoặc kqueue (BSD/macOS).

## Chủ điểm
- Tạo socket: `socket`, `bind`, `listen`, `accept`, `connect`.
- Non-blocking: `fcntl(O_NONBLOCK)`, readiness model.
- `epoll_create`, `epoll_ctl`, `epoll_wait` (Linux); so sánh với select/poll.
- Trong Go: package `net` dùng netpoller của runtime (epoll/kqueue) phía dưới.

## Tham khảo
- TLPI: các chương đã liệt kê.  
- Go: `net`, `syscall` (nâng cao), runtime `netpoll_*` (tham khảo).

## Pet Project
Xem `pet_project-echo-epoll` để chạy TCP echo server đơn giản (1 goroutine/connection). Trên Linux, Go runtime sẽ sử dụng epoll phía dưới.

