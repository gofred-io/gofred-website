#!/bin/bash

# Production deployment script for gofred-website
# This script is executed on the production server

set -e  # Exit on any error

# Configuration
REPO_PATH="${PROD_REPO_PATH:-/opt/gofred-website}"
DOCKER_COMPOSE_FILE="docker-compose.yml"
LOG_FILE="/var/log/gofred-deploy.log"

# DigitalOcean Spaces Configuration
export DO_SPACES_REGION="${DO_SPACES_REGION:-fra1}"
export DO_SPACES_BUCKET="${DO_SPACES_BUCKET:-gofred}"
export DO_SPACES_ACCESS_KEY="${DO_SPACES_ACCESS_KEY}"
export DO_SPACES_SECRET_KEY="${DO_SPACES_SECRET_KEY}"
export DO_API_TOKEN="${DO_API_TOKEN}"
export DO_CDN_ENDPOINT_ID="${DO_CDN_ENDPOINT_ID}"
export DO_PURGE_CACHE="${DO_PURGE_CACHE:-true}"

# Logging function
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1" | tee -a "$LOG_FILE"
}

log "Starting deployment process..."

# Change to repository directory
cd "$REPO_PATH" || {
    log "ERROR: Failed to change to repository directory: $REPO_PATH"
    exit 1
}

log "Changed to repository directory: $(pwd)"

# Pull latest changes from master branch
log "Pulling latest changes from master branch..."
git fetch origin master || {
    log "ERROR: Failed to fetch from origin"
    exit 1
}

git reset --hard origin/master || {
    log "ERROR: Failed to reset to origin/master"
    exit 1
}

log "Successfully pulled latest changes"

# Validate DigitalOcean Spaces credentials
if [ -z "$DO_SPACES_ACCESS_KEY" ] || [ -z "$DO_SPACES_SECRET_KEY" ]; then
    log "WARNING: DigitalOcean Spaces credentials not set"
    log "Set DO_SPACES_ACCESS_KEY and DO_SPACES_SECRET_KEY to enable WASM upload"
    log "WASM file will be built but not uploaded to DigitalOcean Spaces"
fi

# Check if docker-compose.yml exists
if [ ! -f "$DOCKER_COMPOSE_FILE" ]; then
    log "ERROR: docker-compose.yml not found in $REPO_PATH"
    exit 1
fi

# Build and start new containers with production profile
log "Building and starting new containers with production profile..."
docker compose --profile production up --build -d || {
    log "ERROR: Failed to build and start containers with production profile"
    exit 1
}

# Wait a moment for containers to start
sleep 10

# Check if containers are running
log "Checking container status..."
docker compose --profile production ps

# Verify the application is responding
log "Verifying application health..."
if curl -f http://localhost/health > /dev/null 2>&1; then
    log "SUCCESS: Application is responding on port 80"
elif curl -f http://localhost:80/health > /dev/null 2>&1; then
    log "SUCCESS: Application is responding on port 80"
else
    log "WARNING: Health check failed, but deployment completed"
    log "Trying to check container logs..."
    docker compose --profile production logs --tail=20
fi

# Clean up old Docker images
log "Cleaning up old Docker images..."
docker image prune -f || {
    log "WARNING: Failed to clean up old images"
}

log "Deployment completed successfully!"
