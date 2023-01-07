# path to docker compose file
DCOMPOSE:=docker-compose.yml

# improve build time
DOCKER_BUILD_KIT:=COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1

all: down build up

debug: down build-debug up-debug

down:
	docker-compose -f ${DCOMPOSE} down --remove-orphans

build:
	${DOCKER_BUILD_KIT} docker-compose build

build-debug:
	${DOCKER_BUILD_KIT} docker-compose build

up:
	docker-compose --compatibility -f ${DCOMPOSE} up -d --remove-orphans

up-debug:
	docker-compose --compatibility -f ${DCOMPOSE} up --remove-orphans

# Vendoring is useful for local debugging since you don't have to
# reinstall all packages again and again in docker
mod:
	go mod tidy -compat=1.19 && go mod vendor && go install ./...

mock:
	mockgen -source=internal/users/repository.go -destination=internal/users/mocks/repository.go \
	&& mockgen -source=internal/orders/repository.go -destination=internal/orders/mocks/repository.go \
	&& mockgen -source=internal/currency/usecase.go -destination=internal/currency/mocks/usecase.go \
	&& mockgen -source=internal/OrderProducts/usecase.go -destination=internal/OrderProducts/mocks/usecase.go \
	&& mockgen -source=internal/products/usecase.go -destination=internal/users/products/usecase.go \

tests:
	go test ./internal/... -coverpkg=./internal/...
