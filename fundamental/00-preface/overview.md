# Overview — Linux Fundamentals with Go

**Mục tiêu:**  
Nắm vững các khái niệm cốt lõi của Linux từ góc nhìn lập trình hệ thống, đối chiếu với Go runtime, và tổng hợp thành các ghi chú ngắn, dễ tra cứu, bám theo *The Linux Programming Interface (TLPI)*.

---

## 🔧 Nội dung chính

- [Tầng hệ sinh thái Linux](../01-fundamentals/linux-architecture.md): kernel vs user space, chuẩn POSIX/SUS.  
- [System calls](../01-fundamentals/system-calls.md): giao diện user ↔ kernel; giá trị trả về và `errno`.  
- [File / Dir / Permissions](../02-fundamentals/errno-glibc-posix.md): descriptor, cờ mở file, inode, quyền, link.  
- [Process / Signals / Scheduling](./02-processes/process-signals.md): vòng đời tiến trình, tín hiệu, ưu tiên.  
- [Threads / Sync / IPC](./03-concurrency/threads-sync-ipc.md): goroutines và primitives trong Go vs POSIX.  
- [Network / epoll](./04-networking/epoll.md): socket, non-blocking I/O, netpoll trong Go.  
- [Daemon, bảo mật & triển khai](./05-daemon-security/daemon-services.md): service, capability, hạn quyền.

---

## 🧭 Cách học & ghi chép

- Mỗi mục có file `.md` tóm lược khái niệm, liên kết chương TLPI và ví dụ ngắn.
- Ưu tiên ví dụ nhỏ, chạy nhanh; ghi lại lỗi và mapping sang Go.
- Sau mỗi cụm kiến thức, bổ sung 1 pet project ứng dụng nhỏ.

---

## 📚 Tài liệu tham khảo

- *The Linux Programming Interface (TLPI)* — Michael Kerrisk.  
- [man-pages tại man7.org/linux/man-pages](https://man7.org/linux/man-pages/) — tài liệu hệ thống POSIX chi tiết.  
- Go standard library: [`os`](https://pkg.go.dev/os), [`syscall`](https://pkg.go.dev/syscall), [`net`](https://pkg.go.dev/net), [`os/signal`](https://pkg.go.dev/os/signal), [`sync`](https://pkg.go.dev/sync).

---
