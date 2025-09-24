# Dockerfile for gofred-website (Static files only)
# Stage 1: Build WebAssembly
FROM golang:1.25.0-alpine AS builder

# Environment variables for DigitalOcean Spaces upload
# These are set by the deploy.sh script
ENV DO_SPACES_REGION=${DO_SPACES_REGION}
ENV DO_SPACES_BUCKET=${DO_SPACES_BUCKET}
ENV DO_SPACES_ACCESS_KEY=${DO_SPACES_ACCESS_KEY}
ENV DO_SPACES_SECRET_KEY=${DO_SPACES_SECRET_KEY}
ENV DO_API_TOKEN=${DO_API_TOKEN}
ENV DO_CDN_ENDPOINT_ID=${DO_CDN_ENDPOINT_ID}
ENV DO_PURGE_CACHE=${DO_PURGE_CACHE}

# Install necessary packages
RUN apk add --no-cache git ca-certificates tzdata curl jq

# Set working directory
WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the WebAssembly binary
RUN GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o web/main.wasm .

# List the scripts
RUN ls -la scripts/

# Run the upload-wasm.sh script
RUN sh scripts/upload-wasm.sh

# Verify the wasm file was created
RUN ls -la web/main.wasm

# Stage 2: Final stage with static files only
FROM nginx:alpine AS runtime

# Remove default nginx config
RUN rm /etc/nginx/conf.d/default.conf

# Copy custom nginx configuration
COPY nginx.conf /etc/nginx/nginx.conf

# Create directory for static files
RUN mkdir -p /usr/share/nginx/html

# Copy static files and WebAssembly from builder
COPY --from=builder /app/web/ /usr/share/nginx/html/

# Create env.js file with empty window.env object for production
RUN echo "window.env = {WASM_URL: 'https://cdn.gofred.io/web/main.wasm'}" > /usr/share/nginx/html/env.js

# Create nginx user and set permissions
RUN chown -R nginx:nginx /usr/share/nginx/html && \
    chmod -R 755 /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost/health || exit 1

# Start nginx
CMD ["nginx", "-g", "daemon off;"]
