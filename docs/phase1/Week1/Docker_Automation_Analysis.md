# Docker Automation Analysis

## What Docker Does Automatically

### 1. Process Management
**Docker:** `restart: always` in docker-compose.yml
**VM Equivalent:** systemd service configuration
**My thoughts:** Docker gives out of the box tools to manage services while
I need to setup in a VM.

### 2. Dependency Management
**Docker:** Everything in the image (Go binary, libraries)
**VM Equivalent:** Manual installation (apt install golang, dependencies)
**My thoughts:** I can forgot some dependencies and I've to deal with more scripts

### 3. Network Isolation
**Docker:** Automatic container networking
**VM Equivalent:** Manual iptables configuration
**My thoughts:** Docker relieves me from some work.

### 4. Port Mapping
**Docker:** `ports: 8000:8000` in compose file
**VM Equivalent:** Firewall rules + routing
**My thoughts:** I never did it with a VM.

### 5. Logging
**Docker:** `docker compose logs -f`
**VM Equivalent:** journalctl + rsyslog configuration
**My thoughts:** With the VM equivalent I've to deal with multiple commands.

### 6. Resource Limits
**Docker:** `deploy.resources.limits` in compose
**VM Equivalent:** Manual cgroups configuration
**My thoughts:** All configuration in a single place.

### 7. Health Checks
**Docker:** `HEALTHCHECK` in Dockerfile
**VM Equivalent:** Custom monitoring scripts
**My thoughts:** I can miss something while creating a custom script and I have
to manage it.

### 8. Updates/Rollback
**Docker:** Build new image, switch tags
**VM Equivalent:** Manual deployment scripts, backup/restore
**My thoughts:** Docker gives all set of tools through the cli to handle things
that I need to create with a VM.

## Key Questions

1. What would be the hardest to reproduce on a VM?
   - Answer: Managing all the start ups while working with systemD. 

2. What VM challenges do you anticipate in Phase 1?
   - Answer: Dealing with configuration and networking.

3. Why learn VM deployment if Docker is easier?
   - Answer: Not all systems rely on containers.