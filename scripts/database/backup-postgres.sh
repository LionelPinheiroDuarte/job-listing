#!/bin/bash
# PostgreSQL Backup Script

set -e

# Configuration
CONTAINER_NAME="job-listing-postgres"
POSTGRES_USER="${POSTGRES_USER:-jobuser}"
POSTGRES_DB="${POSTGRES_DB:-jobdb}"
BACKUP_DIR="./backups/postgres"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
BACKUP_FILE="backup_${POSTGRES_DB}_${TIMESTAMP}.sql"

# Create backup folder
mkdir -p "$BACKUP_DIR"

echo "🗄️  Starting PostgreSQL backup..."
echo "📦 Database: $POSTGRES_DB"
echo "📁 Backup file: $BACKUP_FILE"

# Backup
docker exec "$CONTAINER_NAME" pg_dump -U "$POSTGRES_USER" "$POSTGRES_DB" > "$BACKUP_DIR/$BACKUP_FILE"

# Compress file
gzip "$BACKUP_DIR/$BACKUP_FILE"

echo "✅ Backup completed: $BACKUP_DIR/$BACKUP_FILE.gz"
echo "📊 Backup size: $(du -h "$BACKUP_DIR/$BACKUP_FILE.gz" | cut -f1)"

# List backups
echo ""
echo "📋 Available backups:"
ls -lh "$BACKUP_DIR/" | tail -5

# Remove backup older than 7 days
echo ""
echo "🧹 Cleaning old backups (keeping last 7 days)..."
find "$BACKUP_DIR" -name "backup_*.sql.gz" -mtime +7 -delete
echo "✅ Cleanup completed"
