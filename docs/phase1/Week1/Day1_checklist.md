---
title: Day 1 - Phase 1 Setup
parent: Week 1
grand_parent: Phase 1
nav_order: 1
---

# Day 1 Checklist - Phase 1 Setup

**Status:** ✅ Completed

---

## ✅ Tasks Completed

### Documentation Created
- [x] Project structure created (`docs/phase1/`)
- [x] Current infrastructure documented
- [x] Docker automation analysis
- [x] AWS configuration documented (security-conscious)
- [x] Systemd learning notes
- [x] job-listing.service file created
- [x] Phase 1 README index
- [x] Day 1 checklist (this file)

### Technical Setup
- [x] AWS CLI verified (v2.33.7)
- [x] AWS account tested (eu-west-3 region)
- [x] Public IP identified and stored locally
- [x] Git commits organized (3 commits)

### Knowledge Acquired
- [x] Understanding what Docker automates
- [x] VM deployment challenges identified
- [x] systemd basics learned
- [x] Service unit file structure understood

---

## 📊 Baseline Metrics Captured

**Current Docker Stack:**
- App image: 36 MB
- Prometheus image: 313 MB
- App memory: 4.91 MB
- Prometheus memory: 50.45 MB
- Both services: ✅ Running

---

## 🎓 Key Learnings

### What I Learned Today

1. **Docker Automation:**
   - Docker automates: process management, networking, logging, health checks
   - VM equivalent requires: systemd, iptables, journalctl, custom scripts

2. **systemd Concepts:**
   - System and service manager for Linux
   - Service units defined in .service files
   - Commands: start, stop, enable, status, journalctl

3. **Security Awareness:**
   - Never commit credentials or sensitive data
   - Use .gitignore for secrets
   - Redact public IPs in documentation

---

## 💭 Reflections

### What Would Be Hardest on VM?
Managing all startups with systemd and dealing with configuration/networking.

### Why Learn VM if Docker is Easier?
Not all systems rely on containers. Understanding traditional deployment gives context for why Docker exists.

### Anticipated Challenges (Week 1)
- Manual dependency management
- Configuration complexity
- Networking setup
- Debugging without container isolation

---

## 🎯 Preparation for Day 2

### Prerequisites Verified
- [x] AWS account configured
- [x] Understanding of systemd
- [x] Service file template ready
- [x] Terraform installed
- [x] Ansible installed

### What's Coming Tomorrow
1. Install Terraform and Ansible
2. Create first EC2 instance with Terraform
3. Manual Go installation on VM
4. First systemd service deployment

---

*Created: Day 1, Week 1, Phase 1*  
*Next: Day 2 - First EC2 Deployment*