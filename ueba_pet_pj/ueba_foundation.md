# UEBA Foundation — System-level Behavior Analytics on Linux (for Golang Developers)

## Purpose

This document defines the **foundation for developing UEBA (User and Entity Behavior Analytics)** modules on **Linux environments**, using low-level telemetry (process, file, and network activity) and **Golang-based collectors**.

Goal: enable developers to build lightweight UEBA agents that detect abnormal behavior directly from Linux system signals, processes, and network events — without relying on heavy SIEM tools.

---

## 1. Architecture Overview

### Layers

| Layer | Description | Implementation |
|-------|--------------|----------------|
| **System Source** | Linux kernel telemetry (syscalls, /proc, auditd, netlink) | C / Golang syscall / eBPF |
| **Collector Agent** | Go-based daemon collecting events (proc, file, network) | Go services (e.g., procwatch, netflow-agent) |
| **Stream / Queue** | Transfers normalized events | Kafka / NATS / HTTP ingestion |
| **Analytics Engine** | Statistical / rule-based anomaly detection | Go or Python models (thresholds, z-score, etc.) |
| **Alert Sink** | Exposes risk events to monitoring platform | REST API / WebSocket / Syslog / Email |

---

## 2. Core UEBA Modules

### 2.1 `procwatch` — Process Behavior Monitor
**Goal:** detect unexpected process creation or privilege escalation.

**Linux sources:** `/proc`, `fork`, `execve`, `setuid`  
**Collected fields:** pid, ppid, user, cmdline, exe, cwd, start_time

```go
func collectProc() []ProcInfo {
    entries, _ := os.ReadDir("/proc")
    var results []ProcInfo
    for _, e := range entries {
        if pid, err := strconv.Atoi(e.Name()); err == nil {
            info := readProc(pid)
            results = append(results, info)
        }
    }
    return results
}
```

**Detection example:**  
- New process under `/tmp` or `/dev/shm`
- Parent-child mismatch (e.g., `sshd` → `/bin/bash` → `nc`)

---

### 2.2 `netflow-agent` — Network Activity Tracker
**Goal:** monitor unusual outbound connections or local port scans.

**Linux sources:** `/proc/net/tcp`, `/proc/net/udp`, `netstat`, `epoll`  
**Collected fields:** local/remote IP, port, UID, inode, process name

```go
func readConnections() []ConnInfo {
    data, _ := os.ReadFile("/proc/net/tcp")
    // parse and convert hex addresses to human-readable IPs
}
```

**Detection example:**  
- Sudden increase in outbound connections  
- Communication with non-whitelisted IP ranges  

---

### 2.3 `auth-audit` — Authentication & Privilege Log Parser
**Goal:** parse `/var/log/auth.log` and `/var/log/secure` for login patterns.

**Linux sources:** PAM, sudo logs, sshd logs  
**Collected fields:** username, IP, timestamp, command, success/failure

```go
func tailAuthLog() {
    watcher, _ := fsnotify.NewWatcher()
    _ = watcher.Add("/var/log/auth.log")
    for e := range watcher.Events {
        go handleAuthEvent(e)
    }
}
```

**Detection example:**  
- Multiple failed logins within short intervals  
- `sudo` access from new IP address  

---

### 2.4 `resmon` — Resource Behavior Monitor
**Goal:** detect CPU/memory spikes per process.

**Linux sources:** `/proc/[pid]/stat`, `/proc/meminfo`, `/proc/loadavg`  
**Collected fields:** pid, cpu%, mem%, threads, io_read, io_write

```go
func collectStats(pid int) (cpu, mem float64) {
    // parse /proc/[pid]/stat for CPU and memory usage
}
```

**Detection example:**  
- High CPU usage from unexpected process  
- Memory leaks or zombie processes  

---

## 3. Event Pipeline & Normalization

### 3.1 Event Format (JSON)

```json
{
  "timestamp": 1733201123,
  "host": "edge-node-01",
  "module": "procwatch",
  "event": {
    "pid": 4231,
    "ppid": 1,
    "exe": "/tmp/runme",
    "user": "www-data",
    "risk": "high"
  }
}
```

### 3.2 Stream Pipeline
- Collectors push to Kafka/NATS topic `ueba.events`
- Processor subscribes, enriches (hostname, user, risk-score)
- Analytics Engine performs statistical / rule-based scoring
- Alerts published to `ueba.alerts`

---

## 4. Detection Techniques

| Technique | Description | Example |
|------------|-------------|----------|
| **Baseline Profiling** | Build normal model for each user/process | “User A normally runs 3 processes, now runs 10” |
| **Threshold Rules** | Fixed limits for CPU, I/O, login failures | CPU > 90%, >5 failed SSH in 1min |
| **Sequence Anomalies** | Unusual process chains | `/usr/bin/sshd` → `/bin/bash` → `curl` → `sh` |
| **Statistical Scoring** | Z-score, moving average | deviation > 2.5σ from baseline |
| **Temporal Correlation** | Time-based context | Off-hours login from new IP |

---

## 5. Implementation Example: UEBA Daemon

```go
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    go startProcWatch(ctx)
    go startNetflow(ctx)
    go tailAuthLog()

    // graceful shutdown
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
    <-sig
    cancel()
}
```

Each sub-module publishes structured events to a message broker or local file.  
These can later be analyzed or visualized by the central UEBA dashboard.

---

## 6. Integration & Deployment

| Component | Deployment | Note |
|------------|-------------|------|
| **Agent binaries** | `/usr/local/bin/ueba-*` | Run as systemd daemons |
| **Config** | `/etc/ueba/config.yaml` | Module enable/disable |
| **Logs** | `/var/log/ueba/*.log` | Rotated daily |
| **Security** | drop privilege via `setcap`, `nosuid` | Run as non-root where possible |
| **Packaging** | Docker / DEB / RPM | for production rollout |

---

## 7. Next Steps

1. Extend collectors with **eBPF-based hooks** for richer telemetry  
2. Integrate **risk scoring and correlation engine** (Go + SQLite or Redis)  
3. Expose REST API for dashboard / alert integrations  
4. Deploy on existing Linux-based clusters (CNSX / Workforce nodes)  

---

## References
- *The Linux Programming Interface* — Michael Kerrisk  
- *Linux Observability with BPF* — David Calavera, Lorenzo Fontana  
- [man7.org/linux/man-pages](https://man7.org/linux/man-pages)  
- Go standard library: `syscall`, `os/signal`, `runtime/pprof`, `net`, `io`
