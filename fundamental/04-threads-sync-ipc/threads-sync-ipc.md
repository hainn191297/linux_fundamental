# Threads, Synchronization & IPC (TLPI Ch.29–33, 45–48, 53–54)

Threads:
- POSIX threads: `pthread_create`, `pthread_join`; chia sẻ địa chỉ, cần đồng bộ tránh race.
- Trong Go: goroutines nhẹ hơn, lập lịch bởi runtime; vẫn cần đồng bộ khi truy cập chung.

Đồng bộ:
- Mutex, RWMutex, Cond, Barrier (khái niệm); tránh deadlock qua order lock nhất quán.
- Atomic operations: `sync/atomic` cho biến nguyên thuỷ.

IPC:
- Pipe/unnamed pipe, FIFO (named pipe) cho streaming giữa process.
- Message queue, shared memory: tốc độ cao nhưng phức tạp đồng bộ.
- Unix domain socket: IPC linh hoạt, hỗ trợ datagram/stream, truyền FD (nâng cao).

Trong Go:
- `sync.Mutex`, `sync.RWMutex`, `sync.WaitGroup`, `chan` (blocking queue built-in).
- IPC: `os.Pipe`, `net.UnixConn`, `net.Pipe` (in-memory), `syscall`/`unix` cho tính năng đặc thù.

Tham khảo TLPI: các chương threads/IPC; `man 7 pipe`, `man 7 unix`.

