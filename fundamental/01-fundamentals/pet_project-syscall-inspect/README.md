# Pet Project: syscall-inspect

Goal: quan sát hệ thống lời gọi (syscalls) của một chương trình để hiểu rõ user space ↔ kernel space.

TLPI: Ch. 1–3 (giới thiệu, tiêu chuẩn, system call interface)

## Ý tưởng

- Dùng `strace` (Linux) hoặc `dtruss` (macOS) để theo dõi syscalls.
- So sánh syscalls giữa một lệnh hệ thống (`ls`) và một chương trình Go đơn giản.

## Cách chạy

```bash
# 1) Quan sát syscalls của lệnh ls
strace -o trace_ls.txt -f -tt -s 80 ls

# 2) Quan sát syscalls của chương trình Go
go run ./sample
strace -o trace_go.txt -f -tt -s 80 go run ./sample

# 3) So sánh
diff -u trace_ls.txt trace_go.txt | less
```

## Bài tập

- Ghi lại chuỗi syscalls khởi động của Go (`execve`, `openat`, `mmap`, `futex`, v.v.).
- Tìm các lỗi `ENOENT`, `EACCES` nếu có và giải thích nguyên nhân.
- Liên hệ với `fundamental/01-fundamentals/system-calls.md` và `golang-mapping.md`.

