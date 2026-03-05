# 🚀 Job-listing - DevOps Learning Journey

[![Terraform](https://img.shields.io/badge/Terraform-7B42BC?logo=terraform&logoColor=white)](https://www.terraform.io/)
[![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white)](https://www.docker.com/)
[![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)](https://golang.org/)
[![AWS](https://img.shields.io/badge/AWS-FF9900?logo=amazonaws&logoColor=white)](https://aws.amazon.com/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

> **A comprehensive DevOps learning project** documenting the journey from traditional VM deployment to modern container orchestration. Built to understand *why* each technology exists by experiencing the problems they solve.

**🎯 Project Goal:** Master production DevOps practices by implementing the same application across three deployment paradigms: VMs, Docker, and Kubernetes.

---

## 📸 Quick Overview
```
Traditional VM (Week 1)  →  Docker Compose (Week 2-5)  →  Kubernetes (Week 6-12)
   15-20 min deploy           1 min 37 sec deploy          30 sec deploy
   Manual steps               Automated                     Self-healing
   Fragile                    Reproducible                  Scalable
```

**Current Status:** ✅ VM Foundation Complete | ⏳ Docker Production In Progress

---

## 🛠️ Tech Stack

### Infrastructure
- **Cloud:** AWS (EC2, VPC, Security Groups)
- **IaC:** Terraform 1.14+
- **Config Mgmt:** Ansible (planned Week 3)
- **CI/CD:** GitHub Actions

### Application
- **Backend:** Go 1.21 (Gin framework)
- **Database:** PostgreSQL 15
- **Monitoring:** Prometheus + Grafana
- **Reverse Proxy:** Nginx

### Container Orchestration
- **Local Dev:** Docker Compose
- **Production:** Kubernetes (EKS) - planned

---

## 🏗️ Architecture Evolution

### Phase 1: Traditional VM Deployment (Week 1) ✅
```
┌─────────────────────────────────────────┐
│           Internet                      │
│              ↓                          │
│      AWS Security Group                 │
│   (SSH: My IP | HTTP: Public)           │
│              ↓                          │
│      ┌──────────────────┐               │
│      │   EC2 Instance   │               │ 
│      │   Ubuntu 22.04   │               │
│      │                  │               │
│      │  ┌────────────┐  │               │
│      │  │   Nginx    │  │ :80           │
│      │  │  (planned) │  │               │
│      │  └─────┬──────┘  │               │
│      │        ↓         │               │
│      │  ┌────────────┐  │               │
│      │  │  Go App    │  │ :8000         │ 
│      │  │  (systemd) │  │               │
│      │  └─────┬──────┘  │               │
│      │        ↓         │               │
│      │  ┌────────────┐  │               │
│      │  │ PostgreSQL │  │ :5432         │
│      │  │ (planned)  │  │               │
│      │  └────────────┘  │               │
│      └──────────────────┘               │
│                                         │
│  Provisioned: Terraform                 │ 
│  Deployed: Bash Script                  │
│  Managed: systemd                       │
└─────────────────────────────────────────┘
```

**Deployment Time:** 1 min 37 sec (automated) | 15-20 min (manual)

---

## 🎯 Key Accomplishments

### Week 1: VM Foundation ✅
- ✅ **Infrastructure as Code** - Terraform modules for AWS EC2
- ✅ **Automated Deployment** - Bash script reducing deploy time by 90%
- ✅ **Service Management** - systemd configuration and troubleshooting
- ✅ **Security Best Practices** - Non-root user, restricted SSH, firewall rules
- ✅ **Comprehensive Documentation** - 9 detailed docs covering every step

### Metrics Achieved
| Metric | Initial | Optimized | Improvement |
|--------|---------|-----------|-------------|
| **Deployment Time** | 15-20 min | 1 min 37 sec | 90% faster |
| **Container Size** | N/A | 28.8MB | Optimized |
| **Response Time P95** | N/A | <50ms | Excellent |
| **Infrastructure Provisioning** | Manual | 37 sec | Terraform |

---

## 🚀 Quick Start

### Prerequisites
- Docker & Docker Compose
- AWS Account (for cloud deployment)
- Terraform 1.14+

### Local Development
```bash
# Clone repository
git clone https://github.com/LionelPinheiroDuarte/job-listing.git
cd job-listing

# Start with Docker Compose
cd deployments
docker compose up --build

# Access application
open http://localhost:8000
```

### VM Deployment (Automated)
```bash
# Provision infrastructure
cd infrastructure/terraform/ec2-vm
terraform init
terraform apply

# Deploy application
ssh ubuntu@<instance-ip>
curl -sSL https://raw.githubusercontent.com/LionelPinheiroDuarte/job-listing/dev/scripts/deploy | bash

# Verify
curl http://<instance-ip>:8000/health
```

**Result:** Production-ready app in under 2 minutes!

---

## 📚 Learning Journey Documentation

### Phase 1: VM Traditional Deployment
- **[Week 1 Retrospective](docs/phase1/week1/Week1_Retrospective.md)** - Complete journey from manual to automated
- **[Terraform Deployment](docs/phase1/week1/Day2_Terraform.md)** - Infrastructure as Code implementation
- **[Automation Script](docs/phase1/week1/Day4_Automation.md)** - Building robust deployment automation

### Key Learnings
1. **Why Containerization Matters**
   - Manual VM deployment: error-prone, slow, inconsistent
   - Experienced firsthand: port conflicts, permission issues, dependency management
   - Docker solves these by design

2. **Infrastructure as Code**
   - Terraform provides reproducibility
   - Version-controlled infrastructure
   - Declarative vs imperative approaches

3. **Linux System Administration**
   - systemd service management
   - Process and port management
   - User permissions and security

---

## 🎓 Skills Demonstrated

### DevOps Engineering
- [x] Infrastructure as Code (Terraform)
- [x] Configuration Management (Bash, Ansible planned)
- [x] CI/CD Pipeline Design (GitHub Actions)
- [x] Container Orchestration (Docker, K8s planned)
- [x] Cloud Infrastructure (AWS EC2, VPC, Security Groups)

### System Administration
- [x] Linux server management (Ubuntu)
- [x] systemd service configuration
- [x] Network security (firewall, SSH hardening)
- [x] Performance monitoring (Prometheus)

### Software Development
- [x] Go application development
- [x] RESTful API design
- [x] Version control (Git workflow, semantic versioning)
- [x] Documentation (technical writing)

---

## 📊 Project Roadmap

### Completed ✅
- [x] **Phase 1, Week 1:** VM deployment foundation
  - Terraform infrastructure
  - Manual deployment process
  - Automation scripting
  - Comprehensive documentation

### In Progress ⏳
- [ ] **Phase 2, Week 2:** Docker Compose production
  - Multi-environment configurations
  - Volume and network management
  - Secrets management
  - AWS deployment

### Planned 📋
- [ ] **Phase 2, Weeks 3-5:** Advanced Docker & Ansible
- [ ] **Phase 3, Weeks 6-12:** Kubernetes orchestration
  - EKS cluster deployment
  - Helm charts
  - Service mesh (Istio)
  - GitOps (ArgoCD)

---

## 💡 Why This Project 

To understand the *why* each tool exists, not just *how* to use it.

---

## 📈 Performance Metrics

### Current Benchmarks (VM Deployment)
- **Infrastructure Provisioning:** 37 seconds (Terraform)
- **Application Deployment:** 60 seconds (automated script)
- **Total Time to Production:** 1 minute 37 seconds
- **Monthly AWS Cost:** ~$8-10 (t3.micro + storage)

### Compared to Docker (local)
- **Deployment Time:** 30 seconds (3x faster)
- **Reproducibility:** 100% (vs ~80% for scripted VM)
- **Developer Experience:** Significantly better

---

## 📞 Connect With Me

- **LinkedIn:** [Lionel Pinheiro Duarte](https://linkedin.com/in/lionelpinheiro)
- **Portfolio:** [website](https://lionelpinheiroduarte.com)
---

<div align="center">

*Last updated: January 30, 2026*

</div>