# Infrastructure - Docker Local

## Operational services
- Go App: port 8000, statut: [RUNNING]
- Prometheus: port 9090, statut: [RUNNING]

## Metrics
- Image app size: 36 MB
- Image prometheus size: 313 MB
- App memory: 4.91 MB
- Prometheus memory: 50.45 MB

## Tests Effectués
- App health: [✅]
- Prometheus UI: [✅]

## Key commands
- Start: `cd deployments && docker compose up -d`
- Stop: `docker compose down`
- Logs: `docker compose logs -f [service]`