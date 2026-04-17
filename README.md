# Go Auth API

![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Web%20Framework-008ECF?logo=gin&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-Database-47A248?logo=mongodb&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-Authentication-black?logo=jsonwebtokens&logoColor=white)
![Swagger](https://img.shields.io/badge/Swagger-API%20Docs-85EA2D?logo=swagger&logoColor=black)
![Docker](https://img.shields.io/badge/Docker-Local%20Setup-2496ED?logo=docker&logoColor=white)
![Docker Compose](https://img.shields.io/badge/Docker%20Compose-Orquestracao-1D63ED?logo=docker&logoColor=white)

API REST de autenticacao com Gin, MongoDB e JWT. Este projeto faz parte do meu processo de aprendizado em Go no backend, com foco em construir uma base limpa, executavel localmente e com organizacao que faca sentido para ambientes reais.


## Objetivo do projeto

Esta e uma API de estudo e evolucao tecnica. 

- aprendizado pratico de `Go` no contexto de backend
- entendimento de estrutura de API REST
- preocupacao com legibilidade, organizacao e manutencao
- capacidade de documentar e empacotar uma aplicacao para outras pessoas rodarem com facilidade

## Stack

- `Go` como linguagem principal
- `Gin` para rotas HTTP e middlewares
- `MongoDB` para persistencia de dados
- `JWT` para autenticacao
- `Swagger` para documentacao da API
- `Docker` e `Docker Compose` para ambiente local
- `Makefile` para comandos de desenvolvimento

## Destaques

- `MongoDB` rodando localmente com Docker, sem depender do MongoDB Atlas
- `Swagger` disponivel em `/swagger/index.html`
- `Request ID`, headers de seguranca e CORS configuravel
- endpoint de saude em `/health`
- rotas versionadas em `/api/v1`
- `Makefile` com comandos simples para desenvolvimento local
- setup pronto para quem quiser clonar e executar rapidamente

## Principais endpoints

- `GET /` retorna metadados da API
- `GET /health` retorna o status do servico
- `POST /api/v1/users` cria um usuario
- `POST /api/v1/users/login` autentica e retorna um token JWT

## Variaveis de ambiente

Exemplo base:

```env
APP_NAME=Go Auth API
APP_VERSION=1.1.0
HTTP_PORT=8080
ALLOWED_ORIGINS=*
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=go_api
MONGO_USER_COLLECTION=users
MONGO_TASK_COLLECTION=tasks
JWT_SECRET=supersecretkey
JWT_ISSUER=go_api_auth
```

## Como executar

### Opcao 1: API + MongoDB com Docker Compose

Para subir toda a stack com containers:

```bash
make up
```

Esse comando sobe:

- API em `http://localhost:8080`
- MongoDB em `mongodb://localhost:27017`

Para parar os containers:

```bash
make down
```

### Opcao 2: MongoDB no Docker e API local com Go

Se voce quiser desenvolver a API localmente e manter apenas o banco em container:

```bash
make dev
```

Esse fluxo sobe o MongoDB com Docker e executa a API localmente com `go run .`.

### Comandos uteis

```bash
make help
make mongo
make api
make logs
make build
make swagger
```

## Docker

O projeto inclui uma configuracao local pensada para facilitar testes e demonstracoes sem exigir conta externa em banco gerenciado.

Arquivos adicionados para isso:

- `Dockerfile` para gerar a imagem da API
- `docker-compose.yml` para subir a stack local
- volume persistente para o MongoDB
- `healthcheck` no container do banco
- integracao com `Makefile` para simplificar a execucao

## Swagger

Abra no navegador:

```text
http://localhost:8080/swagger/index.html
```

Para atualizar a documentacao Swagger:

```bash
go run github.com/swaggo/swag/cmd/swag init
```

## Observacao

Este repositorio representa meu processo de evolucao com Go.
