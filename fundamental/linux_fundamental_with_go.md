# The Linux Programming Interface — Developer Path (for Golang Backend)

## Overall Goals

- Deeply understand how the **Linux kernel & syscalls** work  
- Learn how Linux manages **processes, threads, files, signals, and memory**  
- Understand mapping between **Linux API ↔ Go runtime** (`syscall`, `futex`, `epoll`, `clone`)  
- Build backend tools: daemon, epoll server, IPC chat service  

---

## Linux Fundamentals & System Calls

**TLPI Chapters:** 1–3  

### Core Concepts

- Kernel space vs user space  
- System call flow (`read`, `write`, `open`, `close`)  
- `errno`, `perror`, glibc, POSIX  
- Environment variables & process arguments  

### Mini Project: `syscall-inspect`

```bash
./syscall-inspect ls
```
> Display system calls executed by `ls` using `strace`.

### Go Mapping

- `syscall.Open`, `syscall.Read`, etc. ↔ POSIX API  
- Go runtime directly calls `clone`, `futex`, `epoll_wait`  
- Reference: `runtime/sys_linux_amd64.s`

---

## File, Directory & Permissions

**TLPI Chapters:** 4–5, 14–15, 18  

### Core Concepts

- File descriptor table, open flags, race condition  
- `readv`, `writev`, `pread`, `pwrite`  
- File permission bits (`chmod`, `umask`, `access`)  
- Directory traversal (`opendir`, `readdir`, `nftw`)  
- Symbolic vs hard links  

### Mini Project: `filewalker`

> Traverse `/tmp`, print inode, permissions, and size.  
> Implement parallel traversal with Goroutines + channels.

### Go Mapping

- `os.Open`, `os.Stat`, `filepath.WalkDir`  
- `io.Reader` ↔ `read(2)` syscall  
- Go runtime buffering = user-space stdio buffering  

---

## Process, Signal & Scheduling

**TLPI Chapters:** 6, 20–22, 24–28, 35  

### Core Concepts

- `fork`, `exec`, `wait`, `exit`  
- Process memory layout & stack frame  
- Signal handling (`SIGINT`, `SIGCHLD`, `SIGTERM`)  
- Process groups & sessions  
- Scheduling policies: `SCHED_OTHER`, `SCHED_RR`  

### Mini Project: `procmon`

```bash
./procmon sleep 10
```
> Show parent-child PIDs and log when `SIGCHLD` is received.

### Go Mapping

- `os.StartProcess()` internally invokes `clone` + `execve`  
- `os/signal.Notify()` wraps `sigaction`  
- Goroutine scheduler (M:N model) uses `futex`, `clone`, `sched_yield`  

---

## Threads, Synchronization & IPC

**TLPI Chapters:** 29–33, 45–48, 53–54  

### Core Concepts

- `pthread_create`, `pthread_join`, mutex, semaphore  
- System V vs POSIX shared memory (`shmget` vs `shm_open`)  
- Message queues, semaphores, and shared resources  
- Race condition prevention  

### Mini Project: IPC Chat

> Implement interprocess chat using shared memory + semaphore.

### Go Mapping

- Goroutine = user-level thread  
- `sync.Mutex`, `sync.WaitGroup`, `sync.Cond` similar to pthread primitives  
- Go runtime uses `futex` (`sys_futex`) for sleeping/waking goroutines  

---

## Network Programming & epoll

**TLPI Chapters:** 56–61, 63  

### 📖 Core Concepts

- `socket`, `bind`, `listen`, `accept`, `connect`  
- Non-blocking sockets (`fcntl(O_NONBLOCK)`)  
- `epoll_create`, `epoll_ctl`, `epoll_wait`  
- Level-triggered vs edge-triggered models  
- Performance comparison: select / poll / epoll  

### Mini Project: TCP Echo Server

```bash
go run epoll_server.go
```
> Handle thousands of concurrent connections using epoll.

### Go Mapping

- Go `netpoller` layer is built on top of `epoll`  
- Go runtime registers file descriptors with `epoll_ctl` and waits for I/O events via `epoll_wait`  
- Blocking I/O goroutines are parked by the runtime scheduler  

---

## Daemon, Security & Advanced Topics

**TLPI Chapters:** 37–39, 41–42, 49–50  

### Core Concepts

- Daemonization (detach from terminal, redirect fd)  
- Logging via `syslog`  
- Privilege dropping (SUID, Linux capabilities)  
- `mmap`, `mlock`, `mprotect`  
- Building shared libraries (`.so`)  

### Mini Project: `sysdaemon`

```bash
sudo ./sysdaemon
```

> Background service that logs I/O events into `/var/log/sysdaemon.log`.

### Go Mapping

- `go build -buildmode=c-shared` produces `.so` shared libraries  
- `syscall.Mmap` ↔ `mmap(2)` syscall  
- Capabilities managed via `setcap`  
- Go HTTP server can be daemonized via systemd or custom launcher  

---

## Summary

After mastering all sections, you will be able to:
- Understand **Linux syscalls ↔ Go runtime internals**  
- Build **daemon, IPC, and epoll-based servers**  
- Debug and tune performance using `strace`, `perf`, `top`, `lsof`  

---

## Recommended Reading

- [man7.org/linux/man-pages](https://man7.org/linux/man-pages)  
- *Operating Systems: Three Easy Pieces (OSTEP)*  
- *Linux Performance Tools* — Brendan Gregg  
- Go runtime source: `runtime/netpoll_epoll.go`, `runtime/proc.go`
