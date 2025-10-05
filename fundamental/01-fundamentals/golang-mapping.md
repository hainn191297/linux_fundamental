# Golang Mapping — Linux APIs ↔ Go Runtime

Mục tiêu: nắm cách Go ánh xạ các khái niệm Linux ở tầng API và runtime.

Tổng quan:

- Syscalls: Go cung cấp `syscall` (thấp) và `os`/`io` (cao) để thao tác file, process, socket.
- Error: lỗi từ kernel ánh xạ về `syscall.Errno` và được trình bày dưới dạng `error`.
- Process/Exec: `os/exec` gói gọi `execve`/`clone` phía dưới để tạo process mới.
- Scheduler: runtime dùng `futex`, `clone`, `sched_yield` để lập lịch M:N goroutine↔OS threads.
- Netpoll: Go sử dụng epoll (Linux) hoặc kqueue (BSD/macOS) để multiplex I/O cho `net`.

Bảng ánh xạ (khái quát):

- File I/O: Linux `open/read/write/close` ↔ Go `os.Open`, `File.Read/Write`, `File.Close`.
- Metadata: `stat/lstat` ↔ `os.Stat`, `Lstat` (qua `FileInfo` và `Sys()` chứa `Stat_t`).
- Process: `fork/execve/wait` ↔ `os.StartProcess`/`exec.Command`, `Cmd.Wait`.
- Signals: `sigaction/sigprocmask` ↔ `os/signal.Notify`, `signal.Ignore`.
- Sockets: `socket/bind/listen/accept` ↔ `net.Listen`, `net.Dial` (runtime netpoll xử lý non-blocking).
- Non-blocking: `fcntl(O_NONBLOCK)` ↔ `net` đặt non-blocking FD và đăng ký vào netpoller.

Lưu ý thực tiễn:

- Ưu tiên dùng `os`/`net` trừ khi cần hành vi rất đặc thù thì mới dùng `syscall`/`unix`.
- Unwrap lỗi: `errors.Is`/`errors.As` để lấy `syscall.Errno` hoặc `*os.PathError`.
- Trên Linux, `strace go run ...` sẽ thấy runtime dùng `mmap`, `futex`, `epoll_wait` khi xử lý mạng.
