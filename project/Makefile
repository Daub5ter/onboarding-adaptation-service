FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp
AUTH_BINARY=authApp
ONBOARDING_BINARY=onboardingApp
ADAPTATION_BINARY=adaptationApp
MAIL_BINARY=mailerApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_auth build_onboarding build_adaptation build_mail
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## build_onboarding: builds the onboarding binary as a linux executable
build_onboarding:
	@echo "Building logger binary..."
	cd ../onboarding-service && env GOOS=linux CGO_ENABLED=0 go build -o ${ONBOARDING_BINARY} ./cmd/api
	@echo "Done!"

## build_adaptation: builds the adaptation binary as a linux executable
build_adaptation:
	@echo "Building logger binary..."
	cd ../adaptation-service && env GOOS=linux CGO_ENABLED=0 go build -o ${ADAPTATION_BINARY} ./cmd/api
	@echo "Done!"

## build_auth: builds the auth binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

build_mail:
	@echo "Building auth binary..."
	cd ../mail-service && env GOOS=linux CGO_ENABLED=0 go build -o ${MAIL_BINARY} ./cmd/api
	@echo "Done!"

## build_front: builds the front end binary
build_front:
	@echo "Building front end..."
	cd ../front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## build_swarm: deploy swarm
build_swarm:
	@echo "deploy swarm myapp..."
	docker stack deploy -c swarm.yml myapp
	@echo "Done!"

## down_swarm: remove swarm
down_swarm:
	@echo "remove swarm myapp..."
	docker stack rm myapp
	@echo "Done!"

## leave_swarm: delete swarm
leave_swarm:
	@echo "full delete (leave) swarm..."
	docker swarm leave --force
	@echo "Done!"

