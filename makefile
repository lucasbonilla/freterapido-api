.PHONY: default run build test docs clean
APP_NAME=freterapido-api
DOCKER_RM=

default: run-with-docs

env-serve:
	docker-compose up -d

env-stop:
	docker-compose stop

env-remove:
	docker-compose rm
	
build:
	docker build --rm -t $(APP_NAME) .

test:
	docker build -f Dockerfile.multistage -t docker-freterapido-api-test --progress plain --no-cache --target run-test-stage .