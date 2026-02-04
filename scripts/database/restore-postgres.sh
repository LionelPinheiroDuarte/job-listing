#!/bin/bash
# PostgreSQL Restore Script

set -e

# Configuration
CONTAINER_NAME="job-listing-postgres"
POSTGRES_USER="${POSTGRES_USER:-jobuser}"
POSTGRES_DB="${POSTGRES_DB:-jobdb}"
BACKUP_DIR="./backups/postgres"

# Pass the backup file as argument
if [ -z "$1" ]; then
    echo "❌ Error: Please specify backup file"
    echo "Usage: $0 <backup_file.sql.gz>"
    echo ""
    echo "📋 Available backups:"
    ls -lh "$BACKUP_DIR/"
    exit 1
fi

BACKUP_FILE="$1"

# Does the file exist ?
if [ ! -f "$BACKUP_DIR/$BACKUP_FILE" ]; then
    echo "❌ Error: Backup file not found: $BACKUP_DIR/$BACKUP_FILE"
    exit 1
fi

echo "🔄 Starting PostgreSQL restore..."
echo "📦 Database: $POSTGRES_DB"
echo "📁 Backup file: $BACKUP_FILE"
echo ""
read -p "⚠️  This will OVERWRITE the current database. Continue? (yes/no): " confirm

if [ "$confirm" != "yes" ]; then
    echo "❌ Restore cancelled"
    exit 0
fi

# Uncompress and restore
echo "📥 Decompressing and restoring..."
gunzip -c "$BACKUP_DIR/$BACKUP_FILE" | docker exec -i "$CONTAINER_NAME" psql -U "$POSTGRES_USER" "$POSTGRES_DB"

echo "✅ Restore completed successfully!"

# Verification
echo ""
echo "🔍 Verifying data..."
docker exec "$CONTAINER_NAME" psql -U "$POSTGRES_USER" "$POSTGRES_DB" -c "SELECT COUNT(*) as total_jobs FROM jobs;"
