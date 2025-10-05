# Daemon, Security & Advanced Topics

Daemonization:
- Đặc điểm: không TTY, chạy nền, quản lý bởi init/systemd; tách khỏi session, thiết lập umask, cwd.
- Trên hệ thống hiện đại, ưu tiên dùng systemd unit thay vì tự daemonize 2-fork.

Systemd basics:
- Unit file `.service`: `ExecStart`, `Restart=on-failure`, `User=...`, `Group=...`.
- Journal: log qua stdout/stderr; cấu hình `LimitNOFILE`, `EnvironmentFile`, `CapabilityBoundingSet`.

Least privilege:
- Chạy dưới user/group không phải root; `setcap` để cấp capability tối thiểu (ví dụ `cap_net_bind_service`).
- Chroot/namespace (nâng cao) để cô lập; kết hợp cgroups giới hạn tài nguyên.

Security hardening:
- Seccomp-bpf: lọc syscall cho process; giảm bề mặt tấn công.
- Read-only rootfs, `noexec` mount, `fs.protected_*` (sysctl) khi triển khai containerized.

Packaging & deploy:
- Binaries: `/usr/local/bin/`; cấu hình tại `/etc/<app>/config.yaml`.
- Dịch vụ: `systemctl enable --now <app>.service`; log tại `journalctl -u <app>`.

Tài liệu:
- TLPI: process credentials, capabilities; thực tiễn hiện đại bổ sung từ man-pages và systemd docs.

