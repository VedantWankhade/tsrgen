GATEWAY_BINARY=gateway
CONFLUENCE_BINARY=confluence
JIRA_BINARY=jira
STATS_BINARY=stats

up:
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Docker images started!"

up_build: build_gateway build_confluence build_jira build_stats
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"

down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"

build_gateway:
	@echo "Building gateway binary..."
	cd ./gateway-service && env GOOS=linux CGO_ENABLED=0 go build -o ./bin/${GATEWAY_BINARY} ./cmd
	@echo "Done!"

build_stats:
	@echo "Building stats binary..."
	cd ./stats-service && env GOOS=linux CGO_ENABLED=0 go build -o ./bin/${STATS_BINARY} ./cmd
	@echo "Done!"

build_jira:
	@echo "Building jira binary..."
	cd ./jira-service && env GOOS=linux CGO_ENABLED=0 go build -o ./bin/${JIRA_BINARY} ./cmd
	@echo "Done!"

build_confluence:
	@echo "Building confluence binary..."
	cd ./confluence-service && env GOOS=linux CGO_ENABLED=0 go build -o ./bin/${CONFLUENCE_BINARY} ./cmd
	@echo "Done!"
