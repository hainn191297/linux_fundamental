# 📘 The Linux Programming Interface — Developer Path (for Golang Backend)

## 🧭 Mục tiêu tổng
- Hiểu rõ cơ chế hoạt động của **Linux kernel & syscalls**
- Nắm cách Linux quản lý **process, thread, file, signal, memory**
- Biết mapping giữa **Linux API ↔ Go runtime** (`syscall`, `futex`, `epoll`, `clone`)
- Tự viết các tool backend: daemon, epoll server, IPC service

---

## 1️⃣ Linux Fundamentals & System Calls
**Chương TLPI:** 1–3  

### 📖 Nội dung chính
- Kernel space vs user space  
- System call flow (`read`, `write`, `open`, `close`)  
- `errno`, `perror`, glibc, POSIX  
- Environment variables & process arguments  

### 🧩 Mini Project: `syscall-inspect`
```bash
./syscall-inspect ls
```
> In ra danh sách system call khi chạy lệnh `ls` (sử dụng `strace`).

### 🧠 Go Mapping
- `syscall.Open`, `syscall.Read`, ... ↔ POSIX API  
- Go runtime trực tiếp gọi `clone`, `futex`, `epoll_wait`  
- Xem `runtime/sys_linux_amd64.s`

---

## 2️⃣ File, Directory & Permissions
**Chương TLPI:** 4–5, 14–15, 18  

### 📖 Nội dung
- File descriptor table, open flags, race condition  
- `readv`, `writev`, `pread`, `pwrite`  
- File permission bits (`chmod`, `umask`, `access`)  
- Directory traversal (`opendir`, `readdir`, `nftw`)  
- Symbolic vs hard links  

### 🧩 Mini Project: `filewalker`
> Duyệt `/tmp`, in inode, quyền truy cập, kích thước — song song bằng Goroutine + channel.

### 🧠 Go Mapping
- `os.Open`, `os.Stat`, `filepath.WalkDir`  
- `io.Reader` ↔ system call `read`  
- Go runtime buffer = user-space buffering của `stdio`

---

## 3️⃣ Process, Signal & Scheduling
**Chương TLPI:** 6, 20–22, 24–28, 35  

### 📖 Nội dung
- `fork`, `exec`, `wait`, `exit`  
- Process memory layout, stack frame  
- Signal (`SIGINT`, `SIGCHLD`, `SIGTERM`)  
- Process group, session  
- Scheduling policy: `SCHED_OTHER`, `SCHED_RR`  

### 🧩 Mini Project: `procmon`
```bash
./procmon sleep 10
```
> Hiển thị PID cha–con và log khi nhận `SIGCHLD`.

### 🧠 Go Mapping
- `os.StartProcess()` nội bộ gọi `clone` + `execve`  
- `os/signal.Notify()` wrap kernel `sigaction`  
- Goroutine scheduler (M:N) dùng `futex`, `clone`, `sched_yield`

---

## 4️⃣ Threads, Synchronization & IPC
**Chương TLPI:** 29–33, 45–48, 53–54  

### 📖 Nội dung
- `pthread_create`, `pthread_join`, mutex, semaphore  
- System V vs POSIX shared memory (`shmget` vs `shm_open`)  
- Message queue, semaphore, shared resource  
- Race condition prevention  

### 🧩 Mini Project: IPC Chat
> Chat giữa 2 process qua shared memory + semaphore.

### 🧠 Go Mapping
- Goroutine = user-level thread  
- `sync.Mutex`, `sync.WaitGroup`, `sync.Cond` tương tự pthread  
- Runtime sử dụng futex (`sys_futex`) để sleep/wake goroutine

---

## 5️⃣ Network Programming & epoll
**Chương TLPI:** 56–61, 63  

### 📖 Nội dung
- `socket`, `bind`, `listen`, `accept`, `connect`  
- Non-blocking socket, `fcntl(O_NONBLOCK)`  
- `epoll_create`, `epoll_ctl`, `epoll_wait`  
- Level-triggered vs edge-triggered  
- select/poll/epoll performance  

### 🧩 Mini Project: TCP Echo Server
```bash
go run epoll_server.go
```
> Xử lý hàng ngàn connection đồng thời.

### 🧠 Go Mapping
- Go `netpoller` layer = wrapper quanh `epoll`  
- Go runtime đăng ký FD với `epoll_ctl`, chờ sự kiện I/O qua `epoll_wait`  
- Goroutine block I/O → runtime park thread

---

## 6️⃣ Daemon, Security & Advanced Topics
**Chương TLPI:** 37–39, 41–42, 49–50  

### 📖 Nội dung
- Daemonization (detach from terminal, redirect fd)  
- Logging bằng `syslog`  
- Privilege dropping (SUID, capabilities)  
- `mmap`, `mlock`, `mprotect`  
- Tạo `.so` shared library  

### 🧩 Mini Project: `sysdaemon`
```bash
sudo ./sysdaemon
```
> Ghi log I/O event vào `/var/log/sysdaemon.log` chạy ngầm.

### 🧠 Go Mapping
- `go build -buildmode=c-shared` tạo `.so`  
- `syscall.Mmap` = wrapper của `mmap(2)`  
- `setcap` để cấp quyền network port thấp  
- Go HTTP server chạy như daemon service

---

## 📈 Tổng kết
- Hiểu rõ **Linux syscall → Go runtime**
- Viết được **daemon, IPC, epoll server**
- Debug performance bằng `strace`, `perf`, `top`, `lsof`

---

## 📚 Tài liệu khuyến nghị
- [man7.org/linux/man-pages](https://man7.org/linux/man-pages)  
- *Operating Systems: Three Easy Pieces* (OSTEP)  
- *Linux Performance Tools* – Brendan Gregg  
- Go source: `runtime/netpoll_epoll.go`, `runtime/proc.go`
