.PHONY: help test test-coverage test-verbose build run clean docker-up docker-down

# Variáveis
APP_NAME=crud-golang
MAIN_FILE=cmd/main.go
DOCKER_COMPOSE_FILE=docker-compose.yml

# Comando padrão
help: ## Mostra esta ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Testes
test: ## Executa todos os testes
	go test ./...

test-verbose: ## Executa todos os testes com output detalhado
	go test -v ./...

test-coverage: ## Executa testes com relatório de cobertura
	go test -cover ./...

test-coverage-html: ## Executa testes e gera relatório HTML de cobertura
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Relatório de cobertura gerado: coverage.html"

test-controller: ## Executa apenas testes do controller
	go test -v ./controller/...

test-usecase: ## Executa apenas testes do usecase
	go test -v ./usecase/...

test-repository: ## Executa apenas testes do repository
	go test -v ./repository/...

# Build e execução
build: ## Compila a aplicação
	go build -o bin/$(APP_NAME) $(MAIN_FILE)

run: ## Executa a aplicação
	go run $(MAIN_FILE)

run-build: build ## Compila e executa a aplicação
	./bin/$(APP_NAME)

# Docker
docker-up: ## Inicia os serviços Docker
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

docker-down: ## Para os serviços Docker
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

docker-logs: ## Mostra logs dos serviços Docker
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f

# Swagger
swagger-init: ## Gera a documentação Swagger
	swag init -g cmd/main.go

swagger-serve: ## Abre o Swagger UI no navegador
	@echo "Acesse: http://localhost:8000/swagger/index.html"

# Desenvolvimento
deps: ## Baixa todas as dependências
	go mod download
	go mod tidy

deps-update: ## Atualiza todas as dependências
	go get -u ./...
	go mod tidy

# Limpeza
clean: ## Remove arquivos gerados
	rm -rf bin/
	rm -f coverage.out coverage.html
	go clean

# Linting e formatação
fmt: ## Formata o código
	go fmt ./...

vet: ## Executa go vet para verificar o código
	go vet ./...

lint: fmt vet ## Executa formatação e verificação

# Verificação completa
check: lint test-coverage ## Executa todas as verificações

# Instalação de ferramentas de desenvolvimento
install-tools: ## Instala ferramentas de desenvolvimento
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/axw/gocov/gocov@latest
	go install github.com/AlekSi/gocov-xml@latest

# Relatório de cobertura detalhado
coverage-report: test-coverage-html ## Gera relatório detalhado de cobertura
	@echo "Cobertura por pacote:"
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out | grep total:
	@echo "Relatório HTML gerado: coverage.html"