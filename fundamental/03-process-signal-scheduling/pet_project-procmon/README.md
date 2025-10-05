# Pet Project: procmon

Goal: khởi chạy một chương trình con (ví dụ `sleep 2`), in PID/PPID, chờ kết thúc và in exit status. Bắt `Ctrl+C` để dừng an toàn.

## Chạy thử

```bash
cd fundamental/03-process-signal-scheduling/pet_project-procmon
go run . -- sleep 1
```

## Gợi ý mở rộng

- Ghi thời gian bắt đầu/kết thúc, và duration.
- Log ra JSON line, thêm trường parent-child chain.
- Bắt `SIGCHLD` để log thời điểm child exit.
