# Day 3 - Manual Application Deployment on VM

---

## 🎯 Objectives Achieved

- [x] Installed Go on Ubuntu VM
- [x] Cloned application repository
- [x] Compiled Go application
- [x] Created systemd service
- [x] Debugged and resolved deployment issues
- [x] Application running in production

---

## 🔧 Installation Steps Performed

### 1. Go Installation
**Method:** snap
```bash
sudo snap install go --classic
go version
```

### 2. Repository Clone
[Vos commandes exactes]

### 3. Application Compilation
```bash
cd job-listing
go build
```

### 4. systemd Service Creation
**File:** `/etc/systemd/system/job-listing.service`
```ini

[Unit]
Description=Job Listing Application - Go Web Server
Documentation=https://github.com/LionelPinheiroDuarte/job-listing
After=network.target
# After=network.target means: start only after network is ready

[Service]
Type=simple
# Type=simple: the process runs in foreground

User=appuser
Group=appuser
# User/Group: run as non-root user for security

WorkingDirectory=/opt/job-listing
# WorkingDirectory: where the binary and data.json are located

Environment="PORT=8000"
Environment="GIN_MODE=release"
# Environment: set environment variables for the app

ExecStart=/opt/job-listing/main
# ExecStart: command to start the application (the compiled Go binary)

Restart=always
# Restart=always: if the app crashes, systemd automatically restarts it

RestartSec=10
# RestartSec=10: wait 10 seconds before restarting

StandardOutput=journal
StandardError=journal
# StandardOutput/Error: send logs to journald (view with journalctl)

[Install]
WantedBy=multi-user.target
# WantedBy=multi-user.target: start this service at normal system boot


# COMPARISON WITH DOCKER
# 
# Docker equivalent:
#   restart: always           → Restart=always
#   ports: "8000:8000"        → Not needed in systemd (direct port access)
#   environment: PORT=8000    → Environment="PORT=8000"
#   user: appuser             → User=appuser
#   healthcheck:              → No built-in equivalent (need custom script)
#
# What systemd does that Docker doesn't:
#   - Integrated with system boot (WantedBy)
#   - Automatic logging to journald
#   - Process supervision at OS level
#
# What Docker does that systemd doesn't:
#   - Dependency isolation (libraries, etc.)
#   - Port mapping abstraction
#   - Built-in health checks

```

---

## 🐛 Issues Encountered & Solutions

### Issue 1: User Not Found (status=217/USER)
**Problem:** appuser didn't exist  
**Solution:** Created user with `sudo useradd -r -s /bin/false appuser`

### Issue 2: Binary Not Executable (status=203/EXEC)
**Problem:** Binary not in /opt/job-listing  
**Solution:** 
- Copied files to /opt/job-listing
- Set executable permission: `chmod +x`
- Fixed ownership: `chown appuser:appuser`

---

## 📊 Docker vs VM Comparison

### Deployment Time
- **Docker:** ~30 seconds (`docker compose up`)
- **VM Manual:** ~30-60mins (installation + debug)


### Steps Required
- **Docker:** 1 command
- **VM Manual:** 10+ steps + debugging
- **Insight:** Docker much simpler

### Memory Usage
- **Docker:** 4.91 MB
- **VM:** 1.4 MB (process only, +206 MB OS baseline)

### Reproducibility
- **Docker:** Perfect (Dockerfile)
- **VM:** Manual steps, error-prone

---

## 💡 Key Learnings

### What I Learned
1. systemd service file structure and troubleshooting
2. Linux user management and permissions
3. Go compilation and deployment
4. The value of containerization for reproducibility

### Pain Points Experienced
- Multiple manual steps prone to errors
- systemd error codes require debugging
- File permissions complexity
- No guarantee next deployment will work the same

---

## 📈 Current State

**Application Status:**
- Running: ✅
- Port: 8000
- Memory: 1.4 MB
- Service: enabled (auto-start on boot)

**Access:**
- Local: http://localhost:8000
- Public: http://15.237.74.221:8000

**Logs:**
```bash
sudo journalctl -u job-listing -f
```

---