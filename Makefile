# Define image name and container name
DOCKER_COMPOSE = docker compose
PROJECT_NAME = notification-service

# The default target when running `make` without arguments
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make run          Start services with Docker Compose"
	@echo "  make stop         Stop services"
	@echo "  make build        Build Docker images"
	@echo "  make restart      Restart services"
	@echo "  make logs         Tail the logs of all services"
	@echo "  make logs-service <service>  Tail logs for a specific service"

# Start the services (using Docker Compose)
.PHONY: run
run:
	@$(DOCKER_COMPOSE) up --build -d

# Stop the services
.PHONY: stop
stop:
	@$(DOCKER_COMPOSE) down

# Build Docker images without starting the services
.PHONY: build
build:
	@$(DOCKER_COMPOSE) build

# Restart the services (stop and start them again)
.PHONY: restart
restart:
	@$(DOCKER_COMPOSE) down
	@$(DOCKER_COMPOSE) up --build -d

# Tail the logs of all services
.PHONY: logs
logs:
	@$(DOCKER_COMPOSE) logs -f

# Tail the logs for a specific service (e.g., make logs-service notification-service)
.PHONY: logs-service
logs-service:
	@$(DOCKER_COMPOSE) logs -f $(service)

# Remove all containers and volumes (to clean up)
.PHONY: clean
clean:
	@$(DOCKER_COMPOSE) down -v
