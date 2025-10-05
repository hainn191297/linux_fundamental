# Study Guide — Lộ trình đọc *The Linux Programming Interface (TLPI)*

**Mục tiêu:**  
Đi từ nền tảng → thực hành → mở rộng, mỗi chặng có ghi chú `.md` và 1 pet project nhỏ để củng cố.

---

## 1. Khởi động

- Đọc [Overview](./overview.md) và [Toolchain](./toolchain.md) để chuẩn bị môi trường.
- TLPI Ch.1–2: bức tranh hệ thống, tiêu chuẩn POSIX/SUS.  
  ↳ [man7.org/linux/man-pages](https://man7.org/linux/man-pages/) để tra cứu chi tiết.

---

## 2. System calls & errno

- TLPI **Ch.3** (*System Calls*): cơ chế gọi vào kernel, ABI, giá trị trả về, `errno`.  
- Ghi chú:  
  - [`fundamental/01-fundamentals/system-calls.md`](../01-fundamentals/system-calls.md)  
  - [`errno-glibc-posix.md`](../fundamental/01-fundamentals/errno-glibc-posix.md)  
- Pet project: [`pet_project-syscall-inspect`](../pet_projects/pet_project-syscall-inspect/README.md)

---

## 📂 3. Files, Directories & Permissions

- TLPI **Ch.4–5**, **14–15**, **18**: descriptor, cờ mở, quyền truy cập, `inode`, `link`, `umask`.  
- Ghi chú: [`fundamental/02-files-permissions/files-permissions.md`](../02-files-permissions/files-permissions.md)  
- Pet project: [`pet_project-filewalker`](../02-files-permissions/pet_project-filewalker/README.md)

---

## ⚡ 4. Process, Signals & Scheduling

- TLPI **Ch.6**, **20–22**, **24–28**, **35**: vòng đời tiến trình, tín hiệu, ưu tiên, zombie/orphan.  
- Ghi chú: [`fundamental/03-process-signal-scheduling/process-signal-scheduling.md`](../fundamental/03-process-signal-scheduling/process-signal-scheduling.md)  
- Pet project: [`pet_project-procmon`](../pet_projects/pet_project-procmon/README.md)

---

## 🔄 5. Threads, Synchronization & IPC

- TLPI **Ch.29–33**, **45–48**, **53–54**: thread model, mutex, semaphore, pipe, message queue.  
- Ghi chú: [`fundamental/04-threads-sync-ipc/threads-sync-ipc.md`](../fundamental/04-threads-sync-ipc/threads-sync-ipc.md)  
- Pet project: [`pet_project-ipc-chat`](../pet_projects/pet_project-ipc-chat/README.md)

---

## 🌐 6. Network & epoll

- TLPI **Ch.56–61**, **63**: socket API, blocking/non-blocking I/O, `epoll`, `select`, `poll`.  
- Ghi chú: [`fundamental/05-network-epoll/network-epoll.md`](../fundamental/05-network-epoll/network-epoll.md)  
- Pet project: [`pet_project-echo-epoll`](../pet_projects/pet_project-echo-epoll/README.md)

---

## 🧠 7. Daemon, Security & Advanced Topics

- Chủ đề: daemon hoá, `systemd` integration, capability, hạn quyền, logging, packaging.  
- Ghi chú: [`fundamental/06-daemons-security-advanced/README.md`](../fundamental/06-daemons-security-advanced/README.md)  
- Pet project *(tùy chọn)*: service nhỏ chạy nền, ví dụ `pet_project-daemon-watcher`.

---

## 📘 Gợi ý học tập

- Đọc song song TLPI + `man` tương ứng (`man 2 open`, `man 7 signal`, …).  
- Dịch sang Go runtime tương ứng (`os`, `syscall`, `signal`, `netpoll`).  
- Ghi lại **mapping Linux ↔ Go** trong mỗi file `.md`.

---

> 🧩 Tip: mỗi “Pet Project” nên có thư mục riêng trong `pet_projects/` và ghi `README.md` mô tả mục tiêu, phạm vi, ví dụ chạy, và mapping syscall/Go API.
