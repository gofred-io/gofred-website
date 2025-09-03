# 🐳 Docker Deployment Guide for gofred Website (Nginx Static Files)

This guide explains how to deploy the gofred website using Docker and Docker Compose with Nginx serving static files directly.

## 🏗️ Architecture

The deployment uses a streamlined architecture:
- **Build Stage**: Compiles Go source code to WebAssembly
- **Runtime Stage**: Nginx serves static files (HTML, CSS, JS, WASM, images)
- **No Go Server**: All files served directly by Nginx for maximum performance
- **File Location**: Static files served from `/usr/share/nginx/html` in container

## 📋 Prerequisites

- Docker Engine 20.10.0 or higher
- Docker Compose 2.0.0 or higher
- At least 256MB of available RAM (reduced from previous version)
- Port 3000 available for development (or modify the port mapping)
- Ports 80/443 available for production deployment

## 🚀 Quick Start

### 1. Simple Docker Run

```bash
# Build the Docker image
docker build -t gofred-website .

# Run the container (nginx serves static files)
docker run -d \
  --name gofred-website \
  -p 3000:80 \
  gofred-website

# Access the website at http://localhost:3000
```

### 2. Using Docker Compose (Recommended)

```bash
# Start the application (development mode)
docker-compose up -d

# View logs
docker-compose logs -f

# Stop the application
docker-compose down
```

### 3. Production Deployment with SSL

```bash
# Place SSL certificates in ./ssl/ directory:
# - fullchain.pem
# - privkey.pem

# Start with SSL support
docker-compose --profile production up -d

# This will start nginx with HTTPS on ports 80/443
```

## 🏗️ Build Process

The Dockerfile uses a streamlined 2-stage build process:

1. **Builder Stage**: Compiles the Go application to WebAssembly using `GOOS=js GOARCH=wasm`
2. **Runtime Stage**: Creates an nginx-alpine image with static files served from `/usr/share/nginx/html`

**Key Benefits:**
- **Smaller Image Size**: ~50MB (nginx-alpine + static files)
- **Better Performance**: Direct file serving by nginx
- **Simplified Architecture**: No Go server runtime overhead
- **Production Ready**: Built-in caching, compression, and security headers

## 🔧 Configuration

### Port Mapping

- **Development**: `3000:80` (host:container) - nginx serves static files
- **Production with SSL**: `80:80` and `443:443` - nginx with HTTPS redirect

### File Structure in Container

```
/usr/share/nginx/html/
├── index.html          # Main HTML file
├── index.css          # Stylesheets
├── index.js           # JavaScript files
├── main.wasm          # Compiled WebAssembly from Go
├── wasm_exec.js       # Go WebAssembly runtime
├── env.js             # Environment configuration
├── gofred.css         # Framework styles
└── img/               # Images and assets
    ├── gofred.ico
    └── gofred.png
```

## 📊 Health Checks

The nginx server includes built-in health checks:

- **Endpoint**: `http://localhost/health` (or `https://localhost/health` for SSL)
- **Response**: `{"status":"healthy","service":"gofred-website","server":"nginx"}`
- **Docker Health Check**: Runs every 30 seconds using wget
- **No Backend Dependencies**: Health check is handled directly by nginx

## 🛠️ Customization

### Modify Port Mapping

Edit `docker-compose.yml`:

```yaml
ports:
  - "8080:8080"  # Change first port for host
```

### Resource Limits

Adjust in `docker-compose.yml` (reduced requirements for nginx-only):

```yaml
deploy:
  resources:
    limits:
      memory: 128M  # Sufficient for nginx + static files
      cpus: '0.25'  # Low CPU usage for static serving
```

### SSL/HTTPS Setup

1. Obtain SSL certificates (Let's Encrypt recommended)
2. Place certificates in `./ssl/` directory:
   - `fullchain.pem`
   - `privkey.pem`
3. Uncomment HTTPS server block in `nginx.conf`
4. Update `server_name` in nginx configuration
5. Start with production profile

## 🔍 Troubleshooting

### Build Issues

```bash
# Check if gofred dependency path is correct
ls -la ../gofred

# Clean Docker cache if build fails
docker system prune -a

# Build with verbose output
docker build --progress=plain -t gofred-website .
```

### Container Issues

```bash
# Check container logs
docker-compose logs gofred-website

# Access container shell
docker-compose exec gofred-website sh

# Check if files are properly copied
docker-compose exec gofred-website ls -la server/
```

### Network Issues

```bash
# Check if port is available
lsof -i :3000

# Test health endpoint
curl http://localhost:3000/health

# Check Docker networks
docker network ls
```

## 📈 Performance Optimization

### For Production

1. **Enable Nginx caching** (uncomment caching directives)
2. **Use CDN** for static assets
3. **Enable Gzip compression** (already configured)
4. **Set appropriate resource limits**
5. **Use SSL/TLS** for security

### Monitor Resource Usage

```bash
# Container resource usage
docker stats gofred-website

# Container processes
docker-compose exec gofred-website ps aux
```

## 🔒 Security Considerations

- Application runs as non-root user (`appuser`)
- Security headers are configured in nginx
- WebAssembly files have appropriate CORS headers
- Rate limiting is enabled (10 requests/second)
- Container has minimal attack surface (Alpine Linux)

## 🔄 Updates and Maintenance

### Update Application

```bash
# Rebuild and restart
docker-compose down
docker-compose build --no-cache
docker-compose up -d
```

### Backup

```bash
# Export container
docker save gofred-website > gofred-website-backup.tar

# Backup volumes (if any)
docker-compose exec gofred-website tar -czf /tmp/backup.tar.gz /app
```

## 📝 Logs

- **Application logs**: `docker-compose logs gofred-website`
- **Nginx logs**: `docker-compose logs nginx`
- **Container logs location**: `/var/log/` inside containers

## 🌐 Production Deployment Checklist

- [ ] Update `server_name` in nginx.conf
- [ ] Configure SSL certificates
- [ ] Set up monitoring (Prometheus, Grafana)
- [ ] Configure log rotation
- [ ] Set up backup strategy
- [ ] Configure firewall rules
- [ ] Test health checks
- [ ] Set up CI/CD pipeline
- [ ] Configure domain DNS
- [ ] Test SSL configuration

## 📞 Support

If you encounter issues:

1. Check the logs: `docker-compose logs`
2. Verify prerequisites are met
3. Check port availability
4. Review Docker and Docker Compose versions
5. Consult the [gofred documentation](https://github.com/gofred-io/gofred)

## 🎉 Success!

Once deployed, your gofred website will be available at:
- **Development**: http://localhost:3000
- **Production**: http://your-domain.com (or https://your-domain.com with SSL)

The website showcases the power of gofred for building responsive web applications using only Go!
