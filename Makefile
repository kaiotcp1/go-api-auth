APP_PORT ?= 8080

.PHONY: help up down restart logs mongo api dev build swagger

help:
	@echo "Comandos disponiveis:"
	@echo "  make up       - sobe API + MongoDB com Docker Compose"
	@echo "  make down     - derruba os containers"
	@echo "  make restart  - reinicia os containers"
	@echo "  make logs     - exibe logs da stack"
	@echo "  make mongo    - sobe apenas o MongoDB no Docker"
	@echo "  make api      - inicia a API localmente com go run ."
	@echo "  make dev      - sobe Mongo no Docker e roda a API localmente"
	@echo "  make build    - compila a aplicacao"
	@echo "  make swagger  - atualiza a documentacao Swagger"

up:
	docker compose up --build -d

down:
	docker compose down

restart: down up

logs:
	docker compose logs -f

mongo:
	docker compose up -d mongo

api:
	go run .

dev: mongo
	go run .

build:
	go build ./...

swagger:
	go run github.com/swaggo/swag/cmd/swag init
