# CRUD GoLang - Clean Architecture

Este projeto demonstra a implementação de uma API REST em Go seguindo os princípios de Clean Architecture e Clean Code, com foco em desacoplamento e testabilidade.

## 🏗️ Arquitetura

O projeto segue a Clean Architecture com as seguintes camadas:

```
┌─────────────────┐
│   Controller    │ ← Camada de apresentação (HTTP)
├─────────────────┤
│    Usecase      │ ← Camada de regras de negócio
├─────────────────┤
│   Repository    │ ← Camada de acesso a dados
├─────────────────┤
│   Database      │ ← Camada de infraestrutura
└─────────────────┘
```

### Estrutura de Diretórios

```
├── cmd/                    # Ponto de entrada da aplicação
├── controller/            # Controladores HTTP
├── usecase/              # Casos de uso (regras de negócio)
├── repository/            # Repositórios de dados
├── model/                 # Modelos de domínio
├── dto/                   # Objetos de transferência de dados
├── db/                    # Configuração e conexão com banco
├── test/                  # Mocks para testes
└── docker-compose.yml     # Configuração do ambiente
```

## 🚀 Como Executar

### Pré-requisitos

- Go 1.22+
- Docker e Docker Compose
- PostgreSQL

### 1. Clonar e configurar

```bash
git clone <repository-url>
cd CRUD-GOLANG
```

### 2. Configurar variáveis de ambiente

```bash
cp config.env.example .env
# Editar .env com suas configurações
```

### 3. Executar com Docker

```bash
docker-compose up -d
```

### 4. Executar a aplicação

```bash
go run cmd/main.go
```

A API estará disponível em `http://localhost:8000`

## 🧪 Testes

### Executar todos os testes

```bash
go test ./...
```

### Executar testes com coverage

```bash
go test -cover ./...
```

### Executar testes específicos

```bash
# Testes do controller
go test ./controller/...

# Testes do usecase
go test ./usecase/...

# Testes do repository
go test ./repository/...
```

## 📋 Endpoints da API

- `GET /ping` - Health check
- `GET /products` - Listar todos os produtos
- `POST /product` - Criar novo produto
- `GET /products/:id` - Buscar produto por ID
- `GET /swagger/*` - Documentação Swagger da API

## 📚 Documentação Swagger

A API possui documentação completa gerada automaticamente com Swagger:

- **URL**: `http://localhost:8000/swagger/index.html`
- **Comando para gerar**: `make swagger-init`
- **Comando para abrir**: `make swagger-serve`

### Como usar o Swagger:

1. **Acesse**: `http://localhost:8000/swagger/index.html`
2. **Explore** todos os endpoints disponíveis
3. **Teste** as APIs diretamente pela interface
4. **Veja** os modelos de dados e respostas
5. **Execute** requisições com diferentes parâmetros

## 🔧 Melhorias Implementadas

### 1. **Desacoplamento com Interfaces**
- `ProductRepositoryInterface` para o repository
- `ProductUsecase` para o usecase
- `DatabaseInterface` para operações de banco

### 2. **Injeção de Dependência**
- Todas as dependências são injetadas via construtor
- Facilita a criação de mocks para testes

### 3. **Configuração Flexível**
- Configuração via variáveis de ambiente
- Valores padrão para desenvolvimento

### 4. **Testes Unitários Robustos**
- Mocks centralizados no pacote `test/`
- Testes para todas as camadas (Controller, Usecase, Repository)
- Uso de `go-sqlmock` para testes de banco de dados

### 5. **Separação de Responsabilidades**
- Cada camada tem uma responsabilidade específica
- Controller: apenas validação de entrada e formatação de saída
- Usecase: regras de negócio
- Repository: acesso a dados
- Model: estrutura de dados

## 🎯 Benefícios da Arquitetura

1. **Testabilidade**: Cada camada pode ser testada independentemente
2. **Manutenibilidade**: Código organizado e fácil de entender
3. **Flexibilidade**: Fácil trocar implementações (ex: banco de dados)
4. **Escalabilidade**: Estrutura preparada para crescimento
5. **Desacoplamento**: Mudanças em uma camada não afetam outras

## 📚 Próximos Passos

- [x] Implementar validação de entrada com tags de validação
- [x] Adicionar logging estruturado
- [x] Implementar middleware de autenticação
- [x] Adicionar métricas e observabilidade
- [x] Implementar cache com Redis
- [x] Adicionar documentação da API com Swagger

## 🤝 Contribuindo

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.
