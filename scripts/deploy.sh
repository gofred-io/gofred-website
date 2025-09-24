#!/bin/bash

# Production deployment script for gofred-website
# This script is executed on the production server

set -e  # Exit on any error

# Configuration
REPO_PATH="${PROD_REPO_PATH:-/opt/gofred-website}"
DOCKER_COMPOSE_FILE="docker-compose.yml"
LOG_FILE="/var/log/gofred-deploy.log"

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

# Check if docker-compose.yml exists
if [ ! -f "$DOCKER_COMPOSE_FILE" ]; then
    log "ERROR: docker-compose.yml not found in $REPO_PATH"
    exit 1
fi

# Build and start new containers with production profile
log "Building and starting new containers with production profile..."

# Build with DigitalOcean Spaces build arguments
docker compose --profile production build \
  --build-arg DO_SPACES_REGION="${DO_SPACES_REGION:-fra1}" \
  --build-arg DO_SPACES_BUCKET="${DO_SPACES_BUCKET:-gofred}" \
  --build-arg DO_SPACES_ACCESS_KEY="${DO_SPACES_ACCESS_KEY}" \
  --build-arg DO_SPACES_SECRET_KEY="${DO_SPACES_SECRET_KEY}" \
  --build-arg DO_API_TOKEN="${DO_API_TOKEN}" \
  --build-arg DO_CDN_ENDPOINT_ID="${DO_CDN_ENDPOINT_ID}" \
  --build-arg DO_PURGE_CACHE="${DO_PURGE_CACHE:-true}" \
  --build-arg WASM_FILE="${WASM_FILE:-web/main.wasm}" || {
    log "ERROR: Failed to build containers with DigitalOcean Spaces configuration"
    exit 1
}

# Start the containers
docker compose --profile production up -d || {
    log "ERROR: Failed to start containers with production profile"
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
