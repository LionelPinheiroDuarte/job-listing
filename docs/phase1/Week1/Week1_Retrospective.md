# Week 1 Retrospective - VM Deployment Journey

## 🎯 Week Objectives Review

### Planned Objectives
- [x] Understand VM deployment process
- [x] Deploy application manually
- [x] Create systemd service
- [x] Automate deployment
- [x] Document everything

### Bonus Achievements
- [x] Terraform infrastructure as code
- [x] Complete automation script
- [x] Security best practices (non-root user)
- [x] Professional documentation

---

## 📊 Metrics Summary

### Deployment Times
| Method | Time | Complexity | Reproducibility |
|--------|------|------------|-----------------|
| **Manual (Day 1)** | 15-20 min | High | Low |
| **Script (Day 4)** | 1 min 37 sec | Medium | High |
| **Docker (Ref)** | 30 sec | Low | Perfect |

### Infrastructure Metrics
- **Terraform apply:** 37 seconds
- **Application build:** ~30 seconds
- **Service startup:** ~5 seconds
- **Total automation:** 1 min 37 sec

### Cost Analysis
- **Daily cost (running):** ~$0.25/day
- **Total Week 1 cost:** ~$1.50
- **If kept running:** ~$7.50/month

---

## 💡 Key Learnings

### Technical Skills Acquired
1. **Infrastructure as Code**
   - Terraform providers and resources
   - AWS EC2, VPC, Security Groups
   - State management

2. **Linux System Administration**
   - systemd service management
   - User and permission management
   - Process management (ports, PIDs)

3. **Bash Scripting**
   - Automation scripts
   - Error handling
   - Idempotency

4. **Go Deployment**
   - Manual compilation
   - Binary deployment
   - Production setup

### Pain Points Experienced

1. **Manual Steps are Error-Prone**
   - Forgot files → service crash
   - Wrong permissions → 217/USER error
   - Port conflicts → 203/EXEC error
   
2. **Dependency Management**
   - Go installation manual
   - Version consistency issues
   - System package updates

3. **Service Configuration**
   - systemd complexity
   - User/group setup
   - Path management

4. **Reproducibility Challenges**
   - Fresh VM ≠ same result initially
   - Script iterations needed
   - Documentation critical

---

## 🐳 Docker vs VM Comparison

### What Docker Solves

| Problem | VM Experience | Docker Solution |
|---------|---------------|-----------------|
| **Dependencies** | Manual Go install | Built into image |
| **Reproducibility** | Script + docs needed | Dockerfile guarantees |
| **Speed** | 1min37sec | 30 seconds |
| **Isolation** | systemd + users | Container isolation |
| **Cleanup** | Manual uninstall | `docker compose down` |
| **Portability** | Platform-specific | Works anywhere |

### Why VM Experience Was Valuable

- ✅ **Understand what Docker automates**
- ✅ **Appreciate containerization benefits**
- ✅ **Linux system admin skills**
- ✅ **Debugging capability**
- ✅ **Professional deployment knowledge**

---

## 🎓 Personal Reflections

### What was easier than expected
The process was straightforward:
    - Build the app locally
    - Deploy a server with Terraform (this was a particularly nice experience)
    - Download requirements
    - Create a systemd service
    - Start the service
    - Destroy and repeat

The step-by-step nature of VM deployment made it easy to understand what was happening at each stage, unlike the "magic" of Docker.

### What was harder than expected
Even though the process was easy to understand, dealing with unexpected errors was tiresome. Each deployment revealed new edge cases: port conflicts, missing service files, permission issues. Luckily, automation exists to solve these problems, but creating that automation required experiencing the pain points first.
---

## 📈 Progress Tracking

### Days Completed
- ✅ Day 1: Analysis & AWS Setup
- ✅ Day 2: Terraform & EC2
- ✅ Day 3: Manual Deployment
- ✅ Day 4: Automation Script
- ✅ Day 5: Week Review

### Decision

**Move to Phase 2 (Docker Production)**

**Reasoning:**
- Week 1 accomplished its goal: I now understand what Docker automates
- The pain points are clear (manual steps, error-prone, slower deployment)
- I'm motivated to see the Docker solution with this fresh context
- The 1min37sec automation is good, but Docker's 30sec is better
- I can always return to advanced VM topics (Ansible, load balancing) later if needed

**What I'll carry forward:**
- systemd knowledge for debugging
- Linux administration skills
- Terraform expertise
- Appreciation for Docker's value proposition
