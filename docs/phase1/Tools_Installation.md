# Tools Installation - Phase 1

---

## Terraform

**Version:** 1.14.3  
**Installation Method:** snap  
**Command Used:** `sudo snap install terraform --classic`

**Verification:**
```bash
terraform --version
# Output: Terraform v1.14.3 on linux_amd64
```

**Status:** ✅ Installed and tested

---

## Ansible

**Version:** 2.16.3 (core)  
**Installation Method:** apt (system package)  
**Python Version:** 3.12.3

**Verification:**
```bash
ansible --version
# Output: ansible [core 2.16.3]

ansible localhost -m ping
# Output: localhost | SUCCESS => "ping": "pong"
```

**Status:** ✅ Installed and tested

---

## AWS CLI

**Version:** 2.33.7  
**Installation Method:** Previously configured  
**Region:** eu-west-3 (Paris)

**Status:** ✅ Configured and tested

---

## Summary

All required tools for Phase 1 are installed and operational:
- Terraform: Infrastructure provisioning ✅
- Ansible: Configuration management ✅
- AWS CLI: Cloud interaction ✅
