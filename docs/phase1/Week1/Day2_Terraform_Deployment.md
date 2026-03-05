---
title: Day 2 - First EC2 Deployment with Terraform
parent: Week 1
grand_parent: Phase 1
nav_order: 2
---

# Day 2 - First EC2 Deployment with Terraform

---

## 🎯 Objectives Achieved

- [x] Created Terraform project structure
- [x] Configured AWS provider and variables
- [x] Created EC2 instance with Terraform
- [x] Configured security groups (SSH restricted, HTTP/App public)
- [x] Successfully connected via SSH
- [x] Verified Ubuntu 22.04 deployment

---

## 📊 Infrastructure Created

### Resources Deployed
- **EC2 Instance:** t3.micro (Ubuntu 22.04.5 LTS)
- **Security Group:** 3 ingress rules (SSH/HTTP/App)
- **Key Pair:** job-listing-deployer
- **Subnet:** Created in default VPC

### Instance Details
- **Instance ID:** i-02680dcdff26ac4de
- **Public IP:** 15.237.74.221
- **Region:** eu-west-3 (Paris)
- **AMI:** Ubuntu 22.04 LTS (auto-selected)

---

## 🔐 Security Configuration

**Ports Opened:**
- Port 22 (SSH): My IP only ✅
- Port 80 (HTTP): Public (0.0.0.0/0) ✅
- Port 8000 (App): Public (0.0.0.0/0) ✅
- Port 9090 (Prometheus): CLOSED (access via SSH tunnel) ✅

**Security Features:**
- Non-root SSH access (ubuntu user)
- SSH key-based authentication
- Restricted SSH access to single IP
- Prometheus not exposed publicly

---

## 📁 Terraform Files Structure
```
infrastructure/terraform/ec2-basic/
├── main.tf           # Infrastructure resources
├── variables.tf      # Variable definitions
├── outputs.tf        # Output values
└── terraform.tfvars  # Variable values (gitignored)
```

**Resources in main.tf:**
1. VPC data source (default VPC)
2. Subnet resource (created)
3. AMI data source (Ubuntu 22.04)
4. Security group
5. SSH key pair
6. EC2 instance

---

## 🐛 Issues Encountered & Solutions

### Issue 1: Terraform not reading tfvars
**Problem:** File named `terraform.tvars` instead of `terraform.tfvars`  
**Solution:** Renamed to `terraform.tfvars` (note the 'f')

### Issue 2: Tilde (~) not recognized
**Problem:** `file("~/.ssh/id_rsa.pub")` failed  
**Solution:** Used `file(pathexpand("~/.ssh/id_rsa.pub"))`

### Issue 3: No subnets in default VPC
**Problem:** VPC had no subnets  
**Solution:** Created subnet resource with CIDR 172.31.48.0/20

---

## 💡 Key Learnings

### Terraform Concepts Mastered
- Infrastructure as Code basics
- Provider configuration
- Data sources vs Resources
- Variable management
- Output values
- Resource dependencies

### AWS Concepts
- VPC and subnets
- Security groups (ingress/egress rules)
- EC2 instance creation
- Key pair management

### Docker vs VM Observations
**Memory Usage:**
- Docker local: ~55 MB (app + Prometheus)
- VM baseline: 206 MB (just OS)
- **Insight:** Docker much more lightweight

**Deployment:**
- Docker: `docker compose up` (instant)
- VM: Terraform apply + manual setup needed
- **Insight:** Docker much faster for deployment

---

## 🔧 Current VM State

**Installed:**
- ✅ Ubuntu 22.04.5 LTS
- ✅ systemd operational

**Not Installed (yet):**
- ❌ Go 1.21
- ❌ Application code
- ❌ PostgreSQL
- ❌ Nginx
- ❌ Prometheus

**Next Steps:** Manual installation and configuration

---

## 🎯 Preparation for Day 3

### What's Coming
1. Install Go 1.21 manually
2. Clone and compile application
3. Create systemd service
4. Deploy application
5. Test and validate

### Prerequisites Verified
- [x] VM accessible via SSH
- [x] Ubuntu 22.04 running
- [x] systemd operational
- [x] Service file template ready (from Day 1)

---

## 📊 Cost Tracking

**Current AWS Resources:**
- EC2 t3.micro: ~$0.0104/hour (~$7.50/month)
- EBS 8GB: ~$0.80/month
- Data transfer: Minimal (free tier)

**Estimated Monthly Cost:** ~$8-10

**Note:** Remember to stop/terminate when not in use!

aws ec2 stop-instances --instance-ids [instance_id]
terraform destroy
---
