# Day 4 - Deployment Automation Script

---

## 🎯 Objectives Achieved

- [x] Created automated deployment script
- [x] Handled edge cases (port conflicts, service restarts)
- [x] Tested script on fresh VM instance
- [x] Documented deployment automation

---

## 📜 Script Features

### Automation Includes
- Automatic Go installation detection
- Repository clone/update
- Application compilation
- systemd service management
- User and permission setup
- Port conflict resolution

### Key Improvements
- **Variables:** project_name, project_port, repo_path
- **Error Handling:** Force stop + port cleanup
- **Idempotency:** Can run multiple times safely
- **Status Reporting:** Clear feedback at each step

---

## 🐛 Issues Resolved

### Issue 1: Port Already in Use
**Problem:** `bind: address already in use`  
**Solution:** 
```bash
sudo systemctl stop ${project_name}
sleep 2
sudo fuser -k ${project_port}/tcp
```

### Issue 2: Script Path Assumptions
**Problem:** Script failed when run from different directories  
**Solution:** Use absolute paths with `$HOME/${project_name}`

---

## ⏱️ Deployment Time Comparison

### Manual Deployment
- Steps: 10+ manual commands
- Time: ~5-10 minutes
- Error-prone: Yes
- Reproducible: No

### Automated Script
- Steps: 1 command (`./scripts/deploy.sh`)
- Time: ~2 minutes
- Error-prone: No
- Reproducible: Yes ✅

### Docker (Reference)
- Steps: 1 command (`docker compose up`)
- Time: ~30 seconds
- Error-prone: No
- Reproducible: Yes ✅

**Insight:** Automation helps, but Docker still 4x faster

---

## 🚀 Script Usage
```bash
curl -sSL https://raw.githubusercontent.com/LionelPinheiroDuarte/job-listing/dev/scripts/deploy | bash
```

**Output:**
- Updates system packages
- Installs/verifies Go
- Updates repository
- Builds application
- Deploys to /opt/job-listing
- Restarts service
- Shows status

---

## 💡 Key Learnings

1. **Automation reduces errors** but adds complexity
2. **Edge cases matter** (port conflicts, permissions)
3. **Idempotency is critical** for reliable scripts
4. **Docker solves these problems** by design