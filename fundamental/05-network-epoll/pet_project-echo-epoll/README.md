# Pet Project: TCP echo server

Goal: server TCP lắng nghe trên `:8080`, echo lại dữ liệu nhận được.

## Chạy thử
```bash
cd fundamental/05-network-epoll/pet_project-echo-epoll
go run .

# Terminal khác
nc localhost 8080
hello
hello
```

## Gợi ý mở rộng
- In thêm remote addr và thời gian kết nối.
- Thử giới hạn số kết nối đồng thời; đo thông lượng.
- Thêm tuỳ chọn cờ `-reuseport` (Linux) bằng syscall nâng cao (tuỳ hệ).

