# Production Deployment Guide

This guide explains how to set up automatic deployment to your production server using GitHub Actions.

## Prerequisites

1. **Production Server Requirements:**
   - Ubuntu/Debian server with Docker and Docker Compose installed
   - Git installed
   - SSH access configured
   - Repository cloned to the server

2. **GitHub Repository Secrets:**
   You need to configure the following secrets in your GitHub repository settings:

## GitHub Secrets Configuration

Go to your repository → Settings → Secrets and variables → Actions, and add these secrets:

### Required Secrets

| Secret Name | Description | Example |
|-------------|-------------|---------|
| `PROD_HOST` | Production server IP address or domain | `192.168.1.100` or `your-server.com` |
| `PROD_USER` | SSH username for production server | `deploy` or `ubuntu` |
| `PROD_SSH_KEY` | Private SSH key for authentication | Contents of your private key file |
| `PROD_PORT` | SSH port (optional, defaults to 22) | `22` |
| `PROD_REPO_PATH` | Full path to repository on production server | `/opt/gofred-website` |

### Setting up SSH Key

1. **Generate SSH key pair on your local machine:**
   ```bash
   ssh-keygen -t rsa -b 4096 -C "github-actions-deploy"
   ```

2. **Add public key to production server:**
   ```bash
   # Copy public key to server
   ssh-copy-id -i ~/.ssh/id_rsa.pub user@your-server.com
   
   # Or manually add to ~/.ssh/authorized_keys
   cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys
   ```

3. **Add private key to GitHub Secrets:**
   ```bash
   # Copy private key content
   cat ~/.ssh/id_rsa
   ```
   Paste the entire content (including `-----BEGIN` and `-----END` lines) into the `PROD_SSH_KEY` secret.

## Production Profiles

The deployment uses Docker Compose profiles to run different production configurations:

- **`production`** (default): SSL-enabled setup with ports 80 and 443
- **`prod-http`**: HTTP-only setup with port 80 only

The deployment script uses the `production` profile by default, which includes SSL support.

## Production Server Setup

### 1. Install Dependencies

```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker $USER

# Install Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# Install Git
sudo apt install git -y
```

### 2. Clone Repository

```bash
# Create directory for the application
sudo mkdir -p /opt/gofred-website
sudo chown $USER:$USER /opt/gofred-website

# Clone repository
cd /opt
git clone https://github.com/your-username/gofred-website.git
cd gofred-website
```

### 3. Configure Logging

```bash
# Create log directory
sudo mkdir -p /var/log
sudo touch /var/log/gofred-deploy.log
sudo chown $USER:$USER /var/log/gofred-deploy.log
```

### 4. Test Manual Deployment

```bash
# Make deployment script executable
chmod +x scripts/deploy.sh

# Test deployment
./scripts/deploy.sh
```

## GitHub Actions Workflow

The deployment workflow (`.github/workflows/deploy.yml`) will:

1. Trigger on every push to the `master` branch
2. Connect to your production server via SSH
3. Pull the latest changes from the repository
4. Run the deployment script which:
   - Stops existing containers (using production profile)
   - Builds new Docker images
   - Starts the updated containers (using production profile)
   - Performs health checks on port 80
   - Cleans up old images

## Manual Deployment

You can also trigger deployment manually:

1. Go to your repository on GitHub
2. Click on "Actions" tab
3. Select "Deploy to Production" workflow
4. Click "Run workflow" button

## Monitoring Deployment

### View Deployment Logs

```bash
# On production server
tail -f /var/log/gofred-deploy.log
```

### Check Container Status

```bash
# Check running containers
docker-compose ps

# View container logs
docker-compose logs -f
```

### Health Check

```bash
# Check if application is responding (production profile uses port 80)
curl http://localhost/health

# Or check with explicit port
curl http://localhost:80/health

# For SSL setup, also check HTTPS
curl https://localhost/health
```

## Troubleshooting

### Common Issues

1. **SSH Connection Failed:**
   - Verify SSH key is correctly added to GitHub secrets
   - Check if public key is in `~/.ssh/authorized_keys` on server
   - Ensure firewall allows SSH connections

2. **Git Pull Failed:**
   - Check if repository path is correct
   - Verify git remote is configured properly
   - Ensure user has read access to repository

3. **Docker Build Failed:**
   - Check if Docker is running on production server
   - Verify `docker-compose.yml` file exists
   - Check Docker logs for specific error messages

4. **Application Not Responding:**
   - Check container logs: `docker-compose --profile production logs`
   - Verify port 80 is not blocked by firewall
   - Check if application is binding to correct interface
   - For SSL setup, also check port 443

### Rollback

If deployment fails, you can rollback to previous version:

```bash
# On production server
cd /opt/gofred-website
git log --oneline  # Find previous commit
git reset --hard <previous-commit-hash>
./scripts/deploy.sh
```

## Security Considerations

1. **SSH Key Security:**
   - Use a dedicated SSH key for deployments
   - Regularly rotate SSH keys
   - Limit SSH key permissions on server

2. **Server Security:**
   - Keep server updated
   - Use firewall to restrict access
   - Monitor deployment logs for suspicious activity

3. **Repository Security:**
   - Use GitHub's secret scanning
   - Regularly audit repository access
   - Consider using environment-specific secrets

## Support

If you encounter issues with deployment:

1. Check the GitHub Actions logs
2. Review the deployment log on the server
3. Verify all secrets are correctly configured
4. Test SSH connection manually
5. Check server resources (disk space, memory)
