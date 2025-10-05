# Threads, Synchronization & IPC (TLPI Ch.29–33, 45–48, 53–54)

Mục tiêu: hiểu threads, primitive đồng bộ (mutex/cond), và cơ chế IPC (pipe, FIFO, message queue, shared memory).

## Chủ điểm
- POSIX threads: tạo/join, contention và data race.
- Đồng bộ: mutex, condition variable, semaphore (khái niệm).
- IPC: pipe/unnamed pipe, named pipe (FIFO), message queue, shared memory.
- Trong Go: goroutines + `sync.Mutex`, `sync.WaitGroup`, `chan` và runtime `futex`.

## Tham khảo
- TLPI: các chương đã liệt kê.  
- Go: `sync`, `sync/atomic`, `os`, `net`, `syscall`.

## Pet Project
Xem `pet_project-ipc-chat` để tạo chat đơn giản qua Unix domain socket (local IPC).

