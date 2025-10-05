# Network Programming & epoll (TLPI Ch.56–61, 63)

Socket API:
- Tạo socket: `socket`, `bind`, `listen`, `accept` (server) / `connect` (client).
- Địa chỉ: IPv4/IPv6, TCP/UDP; `getsockopt`/`setsockopt` để cấu hình.

Non-blocking & readiness:
- Non-blocking FD + readiness multiplexer: `select`, `poll`, `epoll` (Linux).
- epoll: chi phí O(1) khi số FD lớn; mô hình level-triggered/edge-triggered.

Trong Go:
- `net` gói trừu tượng socket; runtime tự đặt FD non-blocking và đăng ký với netpoller.
- Trên Linux, netpoller dùng epoll; trên BSD/macOS dùng kqueue.

Pattern cơ bản:
- Server concurrent: 1 goroutine/connection đơn giản; tối ưu hơn dùng worker pool, backpressure.
- Timeout/Deadline: `SetReadDeadline`/`SetWriteDeadline` để tránh treo vô hạn.

Tham khảo TLPI: các chương networking; `man 7 epoll`, `man 2 socket`.

