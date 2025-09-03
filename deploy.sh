#!/bin/bash

# gofred Website Docker Deployment Script
# Usage: ./deploy.sh [dev|prod|stop|rebuild]

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Functions
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

check_dependencies() {
    log_info "Checking dependencies..."
    
    if ! command -v docker &> /dev/null; then
        log_error "Docker is not installed. Please install Docker first."
        exit 1
    fi
    
    if ! command -v docker-compose &> /dev/null; then
        log_error "Docker Compose is not installed. Please install Docker Compose first."
        exit 1
    fi
    
    log_success "Dependencies check passed"
}

check_ports() {
    local port=$1
    if lsof -Pi :$port -sTCP:LISTEN -t >/dev/null ; then
        log_warning "Port $port is already in use"
        read -p "Continue anyway? (y/N): " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            exit 1
        fi
    fi
}

deploy_dev() {
    log_info "Deploying gofred website in development mode (nginx serving static files)..."
    
    check_ports 3000
    
    # Build and start
    docker-compose up --build -d
    
    log_success "Development deployment completed!"
    log_info "Website available at: http://localhost:3000"
    log_info "Health check: http://localhost:3000/health"
    log_info "View logs: docker-compose logs -f"
    log_info "Note: Nginx is serving static files directly from /usr/share/nginx/html"
}

deploy_prod() {
    log_info "Deploying gofred website in production mode with SSL..."
    
    check_ports 80
    check_ports 443
    
    # Check if SSL certificates exist
    if [ ! -f "./ssl/fullchain.pem" ] || [ ! -f "./ssl/privkey.pem" ]; then
        log_warning "SSL certificates not found in ./ssl/ directory"
        log_warning "Place fullchain.pem and privkey.pem in ./ssl/ for HTTPS support"
    fi
    
    # Build and start with production profile
    docker-compose --profile production up --build -d
    
    log_success "Production deployment completed!"
    log_info "Website available at: http://localhost (redirects to HTTPS)"
    log_info "HTTPS available at: https://localhost"
    log_info "Health check: https://localhost/health"
    log_info "View logs: docker-compose logs -f"
    log_info "Note: Using nginx with SSL configuration"
}

stop_deployment() {
    log_info "Stopping gofred website deployment..."
    
    # Stop all profiles
    docker-compose --profile production down
    docker-compose down
    
    log_success "Deployment stopped"
}

rebuild_deployment() {
    log_info "Rebuilding gofred website deployment..."
    
    # Stop, clean, and rebuild
    docker-compose down
    docker-compose build --no-cache
    docker-compose up -d
    
    log_success "Rebuild completed!"
    log_info "Website available at: http://localhost:3000"
    log_info "Note: Static files served by nginx from container"
}

show_status() {
    log_info "Deployment status:"
    echo
    docker-compose ps
    echo
    log_info "Docker images:"
    docker images | grep gofred-website
}

show_logs() {
    log_info "Showing recent logs..."
    docker-compose logs --tail=50 -f
}

cleanup() {
    log_info "Cleaning up Docker resources..."
    
    # Remove containers
    docker-compose down --remove-orphans
    
    # Remove images
    docker rmi gofred-website 2>/dev/null || true
    
    # Clean up unused resources
    docker system prune -f
    
    log_success "Cleanup completed"
}

show_help() {
    cat << EOF
🚀 gofred Website Docker Deployment Script (Nginx Static Files)

Usage: $0 [COMMAND]

Commands:
  dev       Deploy in development mode (port 3000) - nginx serving static files
  prod      Deploy in production mode with SSL (ports 80/443) - nginx with HTTPS
  stop      Stop all deployments
  rebuild   Rebuild and restart development deployment
  status    Show deployment status
  logs      Show application logs
  cleanup   Clean up Docker resources
  help      Show this help message

Examples:
  $0 dev              # Start development server (nginx on port 3000)
  $0 prod             # Start production server with SSL (nginx on ports 80/443)
  $0 stop             # Stop all services
  $0 rebuild          # Rebuild and restart
  $0 logs             # View logs

Architecture:
  - WebAssembly built from Go source code
  - Nginx serves static files (HTML, CSS, JS, WASM, images)
  - No Go server running in production
  - All files served from /usr/share/nginx/html in container

For more information, see DOCKER_DEPLOYMENT.md
EOF
}

# Main script logic
case ${1:-help} in
    dev)
        check_dependencies
        deploy_dev
        ;;
    prod)
        check_dependencies
        deploy_prod
        ;;
    stop)
        stop_deployment
        ;;
    rebuild)
        check_dependencies
        rebuild_deployment
        ;;
    status)
        show_status
        ;;
    logs)
        show_logs
        ;;
    cleanup)
        cleanup
        ;;
    help|--help|-h)
        show_help
        ;;
    *)
        log_error "Unknown command: $1"
        echo
        show_help
        exit 1
        ;;
esac
