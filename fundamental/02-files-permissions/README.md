# Files, Directories & Permissions (TLPI Ch.4–5, 14–15, 18)

Mục tiêu: nắm khái niệm descriptor, cờ mở file, metadata (inode), quyền truy cập, và duyệt thư mục.

## Chủ điểm
- File descriptor table, `open`/`openat` flags (`O_RDONLY`, `O_CREAT`, `O_TRUNC`, `O_APPEND`).
- `read`, `write`, `close`, `lseek`, `stat`/`fstat`/`lstat`.
- Quyền truy cập: `chmod`, `umask`, `chown`; bit `suid/sgid/sticky` (cơ bản).
- Liên kết: hard link vs symlink; rủi ro path traversal.
- Duyệt thư mục: `opendir/readdir` và tương đương Go `filepath.WalkDir`.

## Tham khảo
- TLPI: Ch. 4–5 (File I/O), 14–15 (File attributes & permissions), 18 (Directory).  
- Go: `os`, `io`, `path/filepath`, `fs`.

## Pet Project
Xem `pet_project-filewalker` để duyệt cây thư mục, in inode, quyền và kích thước.

