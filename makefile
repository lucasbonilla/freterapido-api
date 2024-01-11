.PHONY: default run build test docs clean
APP_NAME=freterapido-api

default: run-with-docs

env:
	docker-compose up -d

env-stop:
	docker-compose stop

env-remove:
	docker-compose rm

build:
	docker build --rm -t $(APP_NAME) .

run:
	docker run -p 8080:8080 $(APP_NAME)
	
build-run:
	docker build --rm -t $(APP_NAME) .
	docker run -p 8080:8080 $(APP_NAME)

test:
	docker build -f Dockerfile.multistage -t docker-freterapido-api-test --progress plain --no-cache --target run-test-stage .