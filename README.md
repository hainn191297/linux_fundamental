**Mục Tiêu**
- Nắm chắc Linux fundamentals từ góc nhìn lập trình hệ thống (process, syscall, errno, kiến trúc kernel/userspace).
- Dùng Go để thực nghiệm các khái niệm (map syscall sang Go, viết ví dụ nhỏ).
- Xây dựng một UEBA POC tối thiểu dựa trên kiến thức hệ thống đã học.

**Cấu Trúc Thư Mục**
- `fundamental/00-preface` — phần dạo đầu và cách học: `overview.md`, `toolchain.md`, `study-guide.md`.
- `fundamental/01-fundamentals` — nội dung chính:
  - `linux-architecture.md` — kiến trúc Linux (user/kernel, process, memory, FS...).
  - `system-calls.md` — cơ chế system call, vòng đời, ví dụ minh hoạ.
  - `errno-glibc-posix.md` — lỗi chuẩn, cách tra cứu và xử lý.
  - `golang-mapping.md` — ánh xạ khái niệm hệ thống sang Go.
  - `main.go` — mã ví dụ Go để thực nghiệm.
- `fundamental/01-fundamentals/pet_project-syscall-inspect` — bài thực hành quan sát syscalls.
- `fundamental/02-files-permissions` — file, quyền truy cập, và duyệt thư mục.
- `fundamental/03-process-signal-scheduling` — tiến trình, tín hiệu, lịch trình.
- `fundamental/04-threads-sync-ipc` — threads, đồng bộ, và IPC.
- `fundamental/05-network-epoll` — socket, non-blocking I/O, epoll (Go netpoll).
- `fundamental/linux_fundamental_with_go.md` — ghi chú tổng hợp Linux + Go.
- `fundamental/The Linux Programming Interface.pdf` — tài liệu tham khảo chính (TLPI).
- `ueba_pet_pj/ueba_foundation.md` — nền tảng và phạm vi POC UEBA.

**Yêu Cầu Môi Trường**
- Go đã cài đặt (khuyến nghị bản mới, ví dụ Go 1.21+).
- Linux hoặc macOS với công cụ dòng lệnh cơ bản. Nếu dùng Linux, có thể cài thêm `strace` để quan sát syscall.

**Chạy Ví Dụ Go**
- `cd fundamental/01-fundamentals`
- `go run main.go`
- Tuỳ chọn build: `go build -o demo && ./demo`

**Lộ Trình Học Đề Xuất**
- Khởi động: đọc `fundamental/00-preface/overview.md`, `toolchain.md`, `study-guide.md` để nắm tổng quan và cách tổ chức ghi chép.
- Nền tảng Linux:
  - Đọc `fundamental/01-fundamentals/linux-architecture.md` để có khung kiến trúc.
  - Nghiên cứu `system-calls.md` và thực nghiệm trong `main.go` (mở, đọc, ghi file; tạo process; xử lý lỗi với errno).
  - Tham chiếu `errno-glibc-posix.md` khi bắt gặp lỗi; đối chiếu với Go error.
  - Đọc `golang-mapping.md` để map từ khái niệm hệ thống sang Go (syscall/os, unsafe, v.v.).
- Tổng hợp: bổ sung ghi chú vào `fundamental/linux_fundamental_with_go.md` với ví dụ code và lệnh minh hoạ.

**UEBA POC (Kế Hoạch Tối Thiểu)**
- Mục tiêu: phát hiện hành vi bất thường của user/entity ở mức cơ bản (ví dụ: trình tự lệnh, truy cập file bất thường, tần suất/giờ bất thường).
- Phạm vi dữ liệu: bắt đầu với log hệ thống/local (shell history, tiến trình đang chạy, truy cập file) để đơn giản hoá.
- Pipeline tối thiểu:
  - Thu thập: script/Go CLI đọc số liệu định kỳ (process list, file access, auth log nếu có).
  - Đặc trưng: trích xuất feature đơn giản (tần suất lệnh, entropy, outlier theo giờ/ngày).
  - Phát hiện: baseline thống kê (z-score/IQR) hoặc rule-based trước khi nâng cấp ML.
  - Cảnh báo: in ra console hoặc ghi `jsonl` để trực quan hóa.
- Tài liệu: theo dõi ý tưởng và phạm vi tại `ueba_pet_pj/ueba_foundation.md`.

**Tiến Độ & TODO**
- Preface: đọc và ghi chú (đang tiến hành).
- Fundamentals: hoàn thiện system calls + errno (kế tiếp).
- Demo Go: cập nhật `main.go` với ví dụ mở/ghi file và xử lý lỗi (kế tiếp).
- UEBA POC: chốt dữ liệu đầu vào và chỉ số baseline (sau khi xong fundamentals).

**Cách Ghi Chép Hiệu Quả**
- Mỗi mục nội dung nên có ví dụ lệnh/call thực thi kèm đầu ra ngắn gọn.
- Khi gặp lỗi, ghi lại errno, ngữ cảnh, cách khắc phục và mapping sang Go.
- Ưu tiên ví dụ nhỏ, độc lập và có thể chạy nhanh từ dòng lệnh.

**Gợi Ý Đọc**
- `fundamental/The Linux Programming Interface.pdf` (TLPI) — tra cứu chi tiết về API và hành vi hệ thống.

Nếu bạn muốn, tôi có thể bổ sung ví dụ syscall trong `fundamental/01-fundamentals/main.go` và tạo skeleton cho collector UEBA ban đầu.
