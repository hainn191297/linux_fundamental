# errno, glibc, POSIX — Xử lý lỗi hệ thống

Khái niệm:

- `errno`: biến thread-local do libc đặt khi syscall thất bại (trả `-1`).
- POSIX định nghĩa tập lỗi chuẩn (EACCES, ENOENT, EAGAIN, EINTR...).
- glibc cung cấp wrapper; kiểm tra lỗi ngay sau lời gọi vì `errno` có thể bị ghi đè.

Lỗi thường gặp (tóm tắt):

- `ENOENT` (2): không tồn tại file/đường dẫn.
- `EACCES` (13): không đủ quyền truy cập.
- `EEXIST` (17): đã tồn tại (ví dụ khi dùng `O_CREAT|O_EXCL`).
- `EISDIR` (21): đối tượng là thư mục nhưng mong đợi file thường.
- `EINVAL` (22): tham số không hợp lệ.
- `EAGAIN` (11): tài nguyên tạm thời không sẵn (non-blocking I/O).
- `EINTR` (4): lời gọi bị ngắt bởi tín hiệu (cần retry nếu an toàn).

Chiến lược xử lý:

- Kiểm tra đặc thù: ví dụ nếu `EINTR` → retry; nếu `EAGAIN` → thử lại sau hoặc dùng epoll.
- Khi tạo file: dùng `O_CREAT|O_EXCL` để tránh race “check-then-create”.
- Phân biệt lỗi quyền (`EACCES`) với “không tồn tại” (`ENOENT`) để báo lỗi chính xác.

Trong Go:

- Hầu hết hàm trả `(T, error)`. Lỗi có thể unwrap về `syscall.Errno`.
- So sánh bằng `errors.Is(err, fs.ErrNotExist)` hoặc `errors.Is(err, os.ErrPermission)`; hoặc ép kiểu `errors.As(&pathError)`.
- Non-blocking I/O: các gói mạng xử lý nội bộ; với `syscall.*` có thể thấy `EAGAIN`.

Ví dụ (Go) — phân biệt lỗi không tồn tại và quyền:

```go
f, err := os.Open("/root/secret.txt")
if err != nil {
    if errors.Is(err, fs.ErrNotExist) {
        // file không tồn tại
    } else if errors.Is(err, os.ErrPermission) {
        // không đủ quyền
    } else if errno, ok := err.(syscall.Errno); ok && errno == syscall.EINTR {
        // retry nếu phù hợp
    }
}
_ = f
```

Tài liệu:

- TLPI: chương lỗi và conventions trong các chương đầu (tham chiếu xuyên suốt).
- `man 3 errno`, `man 2 open`, `man 2 read`.
