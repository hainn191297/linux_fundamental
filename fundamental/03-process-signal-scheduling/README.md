# Process, Signals & Scheduling (TLPI Ch.6, 20–22, 24–28, 35)

Mục tiêu: hiểu vòng đời tiến trình, tạo/thoát/wait, tín hiệu, và lịch trình.

## Chủ điểm
- `fork`/`vfork`/`clone` (khái niệm), `execve`, `wait`/`waitpid`.
- Bố cục bộ nhớ tiến trình, biến môi trường, `argv`.
- Tín hiệu: `SIGINT`, `SIGTERM`, `SIGCHLD`; handler và mask.
- Nhóm tiến trình, session, và job control (cơ bản).
- Lịch trình: `SCHED_OTHER`, `nice`/`setpriority` (khái niệm).

## Tham khảo
- TLPI: các chương đã liệt kê ở trên.  
- Go: `os/exec`, `os/signal`, `syscall`, `context`.

## Pet Project
Xem `pet_project-procmon` để chạy lệnh con, theo dõi PID/exit code, và xử lý tín hiệu dừng.

