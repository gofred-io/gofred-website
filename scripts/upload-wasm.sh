#!/bin/bash

# Upload web/main.wasm to DigitalOcean Spaces (S3-compatible)
# This script uploads the WASM file to a DigitalOcean Spaces bucket
# Optimized for Ubuntu - automatically installs required dependencies

set -e  # Exit on any error

# Configuration - Set these environment variables or modify the values below
SPACES_REGION="${DO_SPACES_REGION:-fra1}"
SPACES_BUCKET="${DO_SPACES_BUCKET:-gofred}"
SPACES_ENDPOINT="${DO_SPACES_ENDPOINT:-https://${SPACES_REGION}.digitaloceanspaces.com}"
ACCESS_KEY="${DO_SPACES_ACCESS_KEY}"
SECRET_KEY="${DO_SPACES_SECRET_KEY}"
WASM_FILE="${WASM_FILE:-web/main.wasm}"
WASM_KEY="web/main.wasm"  # Key in the bucket (can be different from filename)
CONTENT_TYPE="application/wasm"

# CDN Cache Purging Configuration
PURGE_CACHE="${DO_PURGE_CACHE:-true}"  # Set to false to skip cache purging
DO_API_TOKEN="${DO_API_TOKEN}"  # DigitalOcean API token for cache purging
CDN_ENDPOINT_ID="${DO_CDN_ENDPOINT_ID}"  # CDN endpoint ID (optional, will be auto-detected)

# Logging function
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1"
}

# Check if required environment variables are set
check_env() {
    if [ -z "$ACCESS_KEY" ]; then
        log "ERROR: DO_SPACES_ACCESS_KEY environment variable is not set"
        log "Please set it with: export DO_SPACES_ACCESS_KEY='your-access-key'"
        exit 1
    fi
    
    if [ -z "$SECRET_KEY" ]; then
        log "ERROR: DO_SPACES_SECRET_KEY environment variable is not set"
        log "Please set it with: export DO_SPACES_SECRET_KEY='your-secret-key'"
        exit 1
    fi
}

# Check if AWS CLI is installed and install if not (Ubuntu only)
check_aws_cli() {
    if ! command -v aws &> /dev/null; then
        log "AWS CLI is not installed, installing on Ubuntu..."
        log "Note: This requires sudo privileges"
        
        # Install AWS CLI on Ubuntu
        log "Updating package list and installing AWS CLI..."
        if sudo apt-get update && sudo apt-get install -y awscli; then
            log "SUCCESS: AWS CLI installed successfully"
        else
            log "ERROR: Failed to install AWS CLI"
            log "Please install manually with: sudo apt-get install awscli"
            exit 1
        fi
    else
        log "AWS CLI is already installed"
    fi
}

# Check if jq is installed and install if not (Ubuntu only)
check_jq() {
    if ! command -v jq &> /dev/null; then
        log "jq is not installed, installing on Ubuntu..."
        log "Note: This requires sudo privileges"
        
        # Install jq on Ubuntu
        log "Installing jq..."
        if sudo apt-get install -y jq; then
            log "SUCCESS: jq installed successfully"
        else
            log "WARNING: Failed to install jq"
            log "Please install manually with: sudo apt-get install jq"
            log "Or set DO_CDN_ENDPOINT_ID environment variable to skip auto-detection"
            return 1
        fi
    else
        log "jq is already installed"
    fi
}

# Check if WASM file exists
check_wasm_file() {
    if [ ! -f "$WASM_FILE" ]; then
        log "ERROR: WASM file not found: $WASM_FILE"
        log "Please make sure the file exists and you're running this script from the project root"
        exit 1
    fi
    
    # Get file size for logging
    FILE_SIZE=$(du -h "$WASM_FILE" | cut -f1)
    log "Found WASM file: $WASM_FILE (Size: $FILE_SIZE)"
}

# Configure AWS CLI for DigitalOcean Spaces
configure_aws() {
    log "Configuring AWS CLI for DigitalOcean Spaces..."
    
    # Set AWS credentials
    export AWS_ACCESS_KEY_ID="$ACCESS_KEY"
    export AWS_SECRET_ACCESS_KEY="$SECRET_KEY"
    export AWS_DEFAULT_REGION="$SPACES_REGION"
    
    # Test connection
    log "Testing connection to DigitalOcean Spaces..."
    if aws s3 ls "s3://$SPACES_BUCKET" --endpoint-url="$SPACES_ENDPOINT" > /dev/null 2>&1; then
        log "SUCCESS: Connected to DigitalOcean Spaces bucket: $SPACES_BUCKET"
    else
        log "ERROR: Failed to connect to DigitalOcean Spaces bucket: $SPACES_BUCKET"
        log "Please check your credentials and bucket name"
        exit 1
    fi
}

# Upload the WASM file
upload_wasm() {
    log "Uploading WASM file to DigitalOcean Spaces..."
    
    # Upload with proper content type and cache headers
    if aws s3 cp "$WASM_FILE" "s3://$SPACES_BUCKET/$WASM_KEY" \
        --endpoint-url="$SPACES_ENDPOINT" \
        --content-type="$CONTENT_TYPE" \
        --cache-control="public, max-age=31536000" \
        --metadata="Content-Encoding=gzip" 2>/dev/null; then
        log "SUCCESS: WASM file uploaded successfully"
    else
        log "ERROR: Failed to upload WASM file"
        exit 1
    fi
    
    # Get the public URL
    PUBLIC_URL="$SPACES_ENDPOINT/$SPACES_BUCKET/$WASM_KEY"
    log "Public URL: $PUBLIC_URL"
}

# Verify the upload
verify_upload() {
    log "Verifying upload..."
    
    if aws s3 ls "s3://$SPACES_BUCKET/$WASM_KEY" --endpoint-url="$SPACES_ENDPOINT" > /dev/null 2>&1; then
        log "SUCCESS: File verified in bucket"
        
        # Get file info
        FILE_INFO=$(aws s3 ls "s3://$SPACES_BUCKET/$WASM_KEY" --endpoint-url="$SPACES_ENDPOINT" --human-readable)
        log "File info: $FILE_INFO"
    else
        log "ERROR: File not found in bucket after upload"
        exit 1
    fi
}

# Get CDN endpoint ID for the bucket
get_cdn_endpoint_id() {
    if [ -n "$CDN_ENDPOINT_ID" ]; then
        log "Using provided CDN endpoint ID: $CDN_ENDPOINT_ID"
        return 0
    fi
    
    if [ -z "$DO_API_TOKEN" ]; then
        log "WARNING: DO_API_TOKEN not set, cannot auto-detect CDN endpoint ID"
        log "Set DO_CDN_ENDPOINT_ID environment variable or DO_API_TOKEN for auto-detection"
        return 1
    fi
    
    log "Auto-detecting CDN endpoint ID..."
    
    # Get all CDN endpoints
    CDN_ENDPOINTS=$(curl -s -X GET \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer $DO_API_TOKEN" \
        "https://api.digitalocean.com/v2/cdn/endpoints" 2>/dev/null)
    
    if [ $? -ne 0 ]; then
        log "ERROR: Failed to fetch CDN endpoints from DigitalOcean API"
        return 1
    fi
    
    # Extract endpoint ID for our bucket
    CDN_ENDPOINT_ID=$(echo "$CDN_ENDPOINTS" | jq -r --arg bucket "$SPACES_BUCKET" \
        '.endpoints[] | select(.origin == $bucket) | .id' 2>/dev/null)
    
    if [ -z "$CDN_ENDPOINT_ID" ] || [ "$CDN_ENDPOINT_ID" = "null" ]; then
        log "WARNING: No CDN endpoint found for bucket: $SPACES_BUCKET"
        log "Make sure CDN is enabled for your DigitalOcean Space"
        return 1
    fi
    
    log "Found CDN endpoint ID: $CDN_ENDPOINT_ID"
    return 0
}

# Purge CDN cache for the uploaded file
purge_cdn_cache() {
    if [ "$PURGE_CACHE" != "true" ]; then
        log "Cache purging disabled (DO_PURGE_CACHE=false)"
        return 0
    fi
    
    if [ -z "$DO_API_TOKEN" ]; then
        log "WARNING: DO_API_TOKEN not set, skipping cache purging"
        log "Set DO_API_TOKEN to enable cache purging"
        return 0
    fi
    
    log "Purging CDN cache..."
    
    # Get CDN endpoint ID
    if ! get_cdn_endpoint_id; then
        log "WARNING: Could not get CDN endpoint ID, skipping cache purging"
        return 0
    fi
    
    # Purge cache for the specific file
    PURGE_RESPONSE=$(curl -s -X DELETE \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer $DO_API_TOKEN" \
        "https://api.digitalocean.com/v2/cdn/endpoints/$CDN_ENDPOINT_ID/cache" \
        -d "{\"files\":[\"/$WASM_KEY\"]}" 2>/dev/null)
    
    if [ $? -eq 0 ]; then
        log "SUCCESS: CDN cache purged for file: /$WASM_KEY"
    else
        log "WARNING: Failed to purge CDN cache (this is not critical)"
        log "You may need to manually purge the cache or wait for TTL expiration"
    fi
}

# Main execution
main() {
    log "Starting WASM upload to DigitalOcean Spaces (Ubuntu optimized)..."
    log "Bucket: $SPACES_BUCKET"
    log "Region: $SPACES_REGION"
    log "Endpoint: $SPACES_ENDPOINT"
    log "Cache purging: $PURGE_CACHE"
    
    check_env
    check_aws_cli
    check_jq || log "WARNING: jq installation failed, cache purging may not work"
    check_wasm_file
    configure_aws
    upload_wasm
    verify_upload
    purge_cdn_cache
    
    log "WASM upload completed successfully!"
    log ""
    log "To use this WASM file in your application, update your code to load it from:"
    log "$SPACES_ENDPOINT/$SPACES_BUCKET/$WASM_KEY"
}

# Run main function
main "$@"
