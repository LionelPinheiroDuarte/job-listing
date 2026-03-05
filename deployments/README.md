# Docker Compose Configurations

## 🎯 Multi-Environment Setup

### Files Structure
```
deployments/
├── docker-compose.yml       # Base configuration (common)
├── docker-compose.dev.yml   # Development overrides
└── docker-compose.prod.yml  # Production overrides
```

---

## 🚀 Usage

### Development Mode
```bash
# Start
docker compose -f docker-compose.yml -f docker-compose.dev.yml up -d

# Stop
docker compose -f docker-compose.yml -f docker-compose.dev.yml down

# Logs
docker compose -f docker-compose.yml -f docker-compose.dev.yml logs -f
```

**Features:**
- Debug enabled
- Verbose logging (debug level)
- Generous resource limits (1GB)
- Hot-reload capable (volume mounts)

---

### Production Mode
```bash
# Start
docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d

# Stop
docker compose -f docker-compose.yml -f docker-compose.prod.yml down

# Logs
docker compose -f docker-compose.yml -f docker-compose.prod.yml logs -f
```

**Features:**
- Security hardened (no-new-privileges, read-only)
- Minimal logging (error level only)
- Strict resource limits (512MB)
- Production optimizations

---

## 📊 Configuration Comparison

| Feature | Development | Production |
|---------|-------------|------------|
| **LOG_LEVEL** | debug | error |
| **DEBUG** | true | false |
| **Memory** | 1GB | 512MB |
| **Security** | Standard | Hardened |
| **Log Size** | 100MB | 10MB |

---

**Variables:**
- `APP_PORT`: Application port (default: 8000)
- `APP_ENV`: Environment (development/production)
- `LOG_LEVEL`: Log verbosity (debug/info/error)
- `PROMETHEUS_PORT`: Prometheus port (default: 9090)

---

*Last updated: Day 6, Week 2, Phase 2*
