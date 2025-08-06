# 🏗️ Desafio Clean Architecture GO

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![Docker](https://img.shields.io/badge/Docker-Required-blue.svg)](https://docker.com)
[![MySQL](https://img.shields.io/badge/MySQL-8.0-orange.svg)](https://mysql.com)

## 📋 Sobre o Projeto

Este projeto é resultado de um **desafio prático** de implementação de Clean Architecture em Go. Partindo de uma base existente com funcionalidade de **CreateOrder**, foi desenvolvida a funcionalidade completa de **ListOrder** seguindo os mesmos padrões arquiteturais.

### 🎯 Objetivo do Desafio
Implementar a listagem de pedidos (ListOrder) em um sistema já existente, garantindo consistência arquitetural através de **múltiplas interfaces de acesso**:

- 🌐 **REST API** - Endpoints HTTP tradicionais
- ⚡ **gRPC** - Comunicação de alta performance  
- 🎯 **GraphQL** - Query language flexível

### 🏆 Funcionalidades Implementadas

- ✅ **Criar Pedido** (CreateOrder) - *Já existente*
- 🆕 **Listar Pedidos** (ListOrder) - **IMPLEMENTADO NO DESAFIO**
- ✅ **Persistência com MySQL**
- ✅ **Injeção de Dependência com Wire**
- ✅ **Event-Driven Architecture**

---

## 🚀 Desafio Proposto

**Contexto:** A partir de um projeto base com funcionalidade de criação de pedidos, implementar a funcionalidade de listagem mantendo os padrões de Clean Architecture.

### 📝 Requisitos do Desafio

1. **UseCase ListOrders** - Implementar a lógica de negócio
2. **REST Endpoint** - `GET /order` para listagem via HTTP
3. **gRPC Service** - `ListOrder` para comunicação RPC
4. **GraphQL Query** - `listOrder` para consultas flexíveis
5. **Migrações** - Estrutura de banco compatível
6. **Testes** - Arquivos `.http` para validação

---

## 🚀 Como Executar

### Pré-requisitos

- [Go 1.21+](https://golang.org/dl/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### 1️⃣ Subir o Banco de Dados

```bash
docker compose up -d
```

> 🔍 **Verificar se subiu:** `docker ps`

### 2️⃣ Inicializar o Banco de Dados

```bash
cd cmd/initdb
go run main.go
```

### 3️⃣ Executar a Aplicação

```bash
cd cmd/ordersystem
go run main.go wire_gen.go
```

### ✅ Confirmação dos Serviços

Se tudo estiver funcionando, você verá:

```
Starting web server on port :8000
Starting gRPC server on port 50051  
Starting GraphQL server on port 8081
```

---

## 🧪 Como Testar

### 🌐 **REST API** - Porta 8000

#### Criar Pedido
```http
POST http://localhost:8000/order
Content-Type: application/json

{
    "id": "order-001",
    "price": 100.50,
    "tax": 10.05
}
```

#### Listar Pedidos
```http
GET http://localhost:8000/order
```

> 📁 **Arquivos de teste:** `api/create_order.http` e `api/list_order.http`

---

### ⚡ **gRPC** - Porta 50051

#### Usando Evans (Recomendado)

```bash
# Instalar Evans
choco install evans  # Windows
# ou baixar de: https://github.com/ktr0731/evans/releases

# Conectar
evans -r repl -p 50051

# Dentro do Evans:
package pb
service OrderService
call CreateOrder
call ListOrder
```

---

### 🎯 **GraphQL** - Porta 8081

#### GraphQL Playground
Acesse: **http://localhost:8081**

#### Mutations e Queries

**Criar Pedido:**
```graphql
mutation {
  createOrder(input: {
    id: "gql-001"
    Price: 200.0
    Tax: 20.0
  }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

**Listar Pedidos:**
```graphql
query {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

---

## 📊 Arquitetura

```
┌─────────────────────────────────────────┐
│                Interfaces               │
│  REST API  │  gRPC  │  GraphQL          │
├─────────────────────────────────────────┤
│              Use Cases                  │
│  CreateOrder  │  ListOrder              │
├─────────────────────────────────────────┤
│               Entities                  │
│             Order                       │
├─────────────────────────────────────────┤
│            Infrastructure               │
│  Database  │  Events  │  Web            │
└─────────────────────────────────────────┘
```

---

## 📝 Estrutura do Projeto

```
├── cmd/
│   ├── initdb/          # Inicialização do banco
│   └── ordersystem/     # Aplicação principal
├── internal/
│   ├── entity/          # Entidades de domínio
│   ├── usecase/         # Casos de uso
│   └── infra/           # Infraestrutura
│       ├── database/    # Repositórios
│       ├── grpc/        # Servidor gRPC
│       ├── graph/       # GraphQL
│       └── web/         # REST API
├── api/                 # Arquivos de teste HTTP
└── docker-compose.yml   # Configuração Docker
```

---

## 🛠️ Tecnologias Utilizadas

- **Go** - Linguagem principal
- **MySQL** - Banco de dados
- **Wire** - Injeção de dependência
- **Chi Router** - HTTP router
- **gRPC** - Comunicação RPC
- **GraphQL** - Query language
- **Docker** - Containerização

---

## 📸 Evidências de Teste

### REST API
![REST API Test](images/ListOrderHttpRequest.PNG)

### gRPC
![gRPC Test](images/ListOrderGRPC.PNG)

### GraphQL
![GraphQL Test](images/ListOrderGraphQL.PNG)


---

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.