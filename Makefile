all: build

build:
	GOARCH=wasm GOOS=js go build -o server/main.wasm main.go

serve:
	go run server/server.go

# Docker deployment targets
docker-build:
	docker build -t hasanhg/gofred-website:latest .

docker-build-tag:
	docker build -t hasanhg/gofred-website:latest -t hasanhg/gofred-website:$(VERSION) .

docker-push:
	docker push hasanhg/gofred-website:latest

docker-push-tag:
	docker push hasanhg/gofred-website:$(VERSION)
	docker push hasanhg/gofred-website:latest

# Deploy to Docker Hub (build and push)
deploy:
	@echo "Building and pushing gofred-website to Docker Hub..."
	docker build -t hasanhg/gofred-website:latest .
	docker push hasanhg/gofred-website:latest
	@echo "Deployment completed!"

# Deploy with version tag
deploy-version:
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION not specified. Use: make deploy-version VERSION=v1.0.0"; \
		exit 1; \
	fi
	@echo "Building and pushing hasanhg/gofred-website:$(VERSION) to Docker Hub..."
	docker build -t hasanhg/gofred-website:latest -t hasanhg/gofred-website:$(VERSION) .
	docker push hasanhg/gofred-website:$(VERSION)
	docker push hasanhg/gofred-website:latest
	@echo "Deployment completed for version $(VERSION)!"

# Login to Docker Hub (interactive)
docker-login:
	docker login

# Full deployment workflow with version
full-deploy:
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION not specified. Use: make full-deploy VERSION=v1.0.0"; \
		exit 1; \
	fi
	@echo "Starting full deployment workflow for version $(VERSION)..."
	@echo "1. Building WebAssembly..."
	$(MAKE) build
	@echo "2. Building Docker image..."
	docker build -t hasanhg/gofred-website:latest -t hasanhg/gofred-website:$(VERSION) .
	@echo "3. Pushing to Docker Hub..."
	docker push hasanhg/gofred-website:$(VERSION)
	docker push hasanhg/gofred-website:latest
	@echo "Full deployment completed for version $(VERSION)!"

# Development Docker commands
docker-dev:
	docker compose up --build -d

docker-dev-logs:
	docker compose logs -f

docker-dev-stop:
	docker compose down

# Production Docker commands
docker-prod:
	docker compose --profile production up --build -d

docker-prod-stop:
	docker compose --profile production down

# Clean up Docker images
docker-clean:
	docker rmi hasanhg/gofred-website:latest || true
	docker system prune -f

.PHONY: all build serve docker-build docker-build-tag docker-push docker-push-tag deploy deploy-version docker-login full-deploy docker-dev docker-dev-logs docker-dev-stop docker-prod docker-prod-stop docker-clean