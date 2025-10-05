# Reading Notes — Fundamentals Module

## Linux Architecture
- Linux kết hợp kernel của Linus Torvalds (1991) với công cụ GNU, kế thừa chuẩn POSIX/SUS để đảm bảo tính tương thích ở tầng API.
- Phân lớp hệ thống: kernel xử lý tài nguyên và syscall, user space chạy process/daemon, thư viện như glibc cầu nối, và shell/tiện ích hỗ trợ tương tác.
- Đối với dự án UEBA bằng Go, kernel là nguồn tín hiệu hành vi, trong khi collector chạy ở user space dùng API do glibc/Go cung cấp.

## System Calls
- Syscall là cơ chế chuyển quyền điều khiển có kiểm soát từ user space vào kernel thông qua wrapper libc và lệnh CPU `syscall`.
- Giá trị trả về âm (`-1`) biểu thị lỗi, với `errno` thread-local mô tả chi tiết; các lệnh quan sát như `strace -f -tt` giúp theo dõi chuỗi syscall.
- Go runtime bao bọc syscall qua package `syscall`/`os`, đồng thời sử dụng `clone`, `futex`, `epoll` nội bộ để lập lịch và multiplex I/O.

## errno, glibc, POSIX
- `errno` lưu lỗi mới nhất từ syscall thất bại; POSIX định nghĩa tập lỗi chuẩn như `ENOENT`, `EACCES`, `EINTR`, `EAGAIN`.
- Chiến lược xử lý cần phân biệt lỗi quyền, tồn tại, hay lỗi tạm thời; với `EINTR` nên retry nếu an toàn.
- Trong Go, lỗi unwrap về `syscall.Errno` hoặc `*os.PathError`; dùng `errors.Is`/`errors.As` để nhánh logic rõ ràng.

## Golang Mapping
- Go cung cấp lớp cao (`os`, `net`, `io`) ánh xạ trực tiếp các syscall `open/read/write/close`, `stat`, `socket`, `fork/execve/wait`.
- Runtime Go thiết lập FD non-blocking và sử dụng netpoll (epoll trên Linux) để xử lý I/O đồng thời cho goroutine.
- Khi cần kiểm soát thấp, có thể dùng `golang.org/x/sys/unix`; tuy nhiên đa số tác vụ nên tận dụng abstraction chuẩn của Go để giữ tính portable.
