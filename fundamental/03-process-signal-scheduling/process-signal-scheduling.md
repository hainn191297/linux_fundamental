# Process, Signals & Scheduling (TLPI Ch.6, 20–22, 24–28, 35)

Process lifecycle:
- Tạo tiến trình: lịch sử `fork`/`vfork`; trên Linux hiện đại thường dùng `clone` với flags.
- Thay thế ảnh: `execve` tải chương trình mới; kế thừa FD và môi trường.
- Chờ con: `wait`/`waitpid` trả về `status` (mã thoát hoặc tín hiệu kết thúc).

Môi trường & bộ nhớ:
- Bố cục: text, data, heap, stack, mmapped regions.
- Biến môi trường `environ`, đối số chương trình `argv`; kế thừa qua `execve`.

Tín hiệu:
- Phổ biến: `SIGINT`, `SIGTERM`, `SIGKILL`, `SIGCHLD`, `SIGHUP`.
- Handler & mask: `sigaction` thiết lập hành vi; tránh gọi hàm không reentrant trong handler.
- Ứng dụng: dừng mềm, reload cấu hình, giám sát child.

Scheduling (khái niệm):
- Chính sách: `SCHED_OTHER` (CFS), `SCHED_FIFO`, `SCHED_RR`.
- Ưu tiên, `nice`, `setpriority` ảnh hưởng phân chia CPU.

Trong Go:
- Tạo process: `exec.Command`, `Cmd.Start/Wait`, cấu hình `SysProcAttr` khi cần.
- Bắt tín hiệu: `os/signal.Notify` để xử lý `SIGINT`, `SIGTERM` cho shutdown an toàn.
- Goroutine scheduling độc lập với OS scheduling nhưng phụ thuộc tài nguyên thread/CPU.

Tham khảo TLPI: các chương nêu trên; `man 2 fork`, `man 2 execve`, `man 2 waitpid`, `man 7 signal`.

