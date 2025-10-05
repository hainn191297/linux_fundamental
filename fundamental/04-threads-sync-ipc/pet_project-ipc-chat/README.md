# Pet Project: ipc-chat (Unix domain socket)

Goal: server và client giao tiếp qua Unix domain socket tại `/tmp/ipc-chat.sock`.

## Chạy thử
```bash
cd fundamental/04-threads-sync-ipc/pet_project-ipc-chat

# terminal 1: chạy server
go run . server

# terminal 2: gửi tin nhắn từ client
go run . client "hello"
```

## Gợi ý mở rộng
- Hỗ trợ nhiều client (goroutine per connection).
- Gửi sự kiện có timestamp và tên client.
- Thêm xử lý signal để cleanup socket file.

