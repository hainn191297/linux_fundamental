
---

# System Calls — Giao diện User ↔ Kernel (TLPI Ch.3)

System call là điểm vào có kiểm soát từ user space xuống kernel để yêu cầu dịch vụ hệ thống.

Các syscall thường gặp: `open/openat`, `read`, `write`, `close`, `stat`, `mmap`, `fork/clone`, `execve`, `wait`, `socket`, `bind`, `connect`.

Luồng gọi (rút gọn):

```
User Program → libc wrapper → CPU instr (syscall) → Kernel handler → return
```

Key points:

- ABI & số syscall: phụ thuộc kiến trúc (x86_64, arm64). Trên Linux, wrapper (glibc) gọi đến instruction `syscall` (x86_64).
- Return value: thường là `-1` báo lỗi, `errno` chứa mã lỗi; trên thành công trả về giá trị không âm (fd, bytes, pid...).
- `errno` là thread-local; phải đọc ngay sau lời gọi thất bại (xem `errno-glibc-posix.md`).
- Blocking vs non-blocking: `read` trên socket có thể trả `EAGAIN` nếu non-blocking.
- Reentrancy: tránh dùng hàm không reentrant trong signal handler.

Thực hành quan sát:

- Linux: `strace -f -tt -s 80 <cmd>` để xem chuỗi syscall, thời gian, tham số.
- Go chương trình: `strace -f -tt go run main.go` sẽ thấy `execve`, open thư viện, `mmap`, `futex`, I/O...

Liên hệ Go:

- Go cung cấp package `syscall` và `os` làm wrapper; lỗi ánh xạ về `syscall.Errno` và `error` của Go.
- Scheduler của Go dùng `futex`, `clone`, `epoll` (Linux) ở tầng runtime để multiplex I/O và goroutine.
