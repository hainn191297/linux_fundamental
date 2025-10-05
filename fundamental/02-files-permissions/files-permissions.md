# Files, Directories & Permissions (TLPI Ch.4–5, 14–15, 18)

Khái niệm chính:

- File descriptor (FD): số nguyên tham chiếu tới entry trong bảng mở file của process.
- Cờ mở file: `O_RDONLY`, `O_WRONLY`, `O_RDWR`, `O_CREAT`, `O_TRUNC`, `O_APPEND`, `O_EXCL`.
- Syscalls cốt lõi: `open/openat`, `read`, `write`, `close`, `lseek`, `stat/fstat/lstat`.
- Metadata & inode: kích thước, mode, uid/gid, thời gian (`atime`, `mtime`, `ctime`).
- Quyền: `rwx` cho user/group/other; `umask` ảnh hưởng quyền mặc định khi tạo file.
- Link: hard link chia sẻ inode; symlink trỏ tới đường dẫn; khác biệt khi xoá/di chuyển.
- Duyệt thư mục: `opendir/readdir/closedir`; trong Go dùng `os.ReadDir`/`filepath.WalkDir`.

Mẫu sử dụng an toàn:

- Tạo file mới: dùng `O_CREAT|O_EXCL` để tránh race điều kiện.
- Kiểm tra quyền: thông qua thao tác thật (open) hơn là `access(2)` (TOCTOU).
- Làm việc với đường dẫn: cân nhắc `openat` + `O_NOFOLLOW` để giảm rủi ro symlink.

Trong Go:

- `os.Open`, `os.Create`, `os.OpenFile` (hỗ trợ cờ và mode), `File.Stat`, `os.Lstat`.
- `fs.FileInfo` cung cấp `Mode()` (bao gồm bit loại file) và có thể lấy `*syscall.Stat_t` qua `Sys()`.
- Duyệt: `os.ReadDir` (không stat từng file) hoặc `filepath.WalkDir` (linh hoạt hơn).

Tham khảo TLPI:

- Ch.4–5: File I/O basics; Ch.14–15: File attributes & permissions; Ch.18: Directory & links.
