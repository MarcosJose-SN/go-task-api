# 🚀 Go Task API

API REST desenvolvida em **Go (Golang)** para gerenciamento de tarefas, com autenticação JWT, PostgreSQL e ambiente totalmente containerizado com Docker.

Projeto desenvolvido com foco em práticas de **Backend**, organização de código, autenticação, banco de dados, testes e integração com Docker.

---

## 🛠️ Tecnologias

* **Go 1.27**
* **PostgreSQL 16**
* **Docker**
* **Docker Compose**
* **JWT**
* **bcrypt**
* **REST API**
* **SQL**
* **Git / GitHub**

---

## 📌 Funcionalidades

### 🔐 Autenticação

* Cadastro de usuários
* Login com usuário e senha
* Senhas protegidas com bcrypt
* Geração de tokens JWT
* Validação de tokens
* Middleware de autenticação
* Proteção das rotas da API

### 📋 Gerenciamento de tarefas

* Criar tarefa
* Listar tarefas
* Buscar tarefa por ID
* Atualizar tarefa
* Excluir tarefa
* Marcar tarefa como concluída
* Validação dos dados enviados

### 🗄️ Banco de dados

Utiliza **PostgreSQL** para armazenamento dos usuários e tarefas.

Estrutura principal:

```text
users
├── id
├── username
├── email
└── password_hash

tasks
├── id
├── title
└── completed
```

---

## 🐳 Docker

O projeto possui ambiente configurado com **Docker Compose**, permitindo executar a API e o PostgreSQL através de containers.

Arquitetura:

```text
┌─────────────────────────┐
│      Cliente / API      │
│      localhost:8080     │
└────────────┬────────────┘
             │
             ▼
┌─────────────────────────┐
│      Go Task API        │
│        Go + JWT         │
└────────────┬────────────┘
             │
             ▼
┌─────────────────────────┐
│      PostgreSQL 16      │
│        port 5432        │
└─────────────────────────┘
```

Para iniciar o projeto:

```bash
docker compose up -d --build
```

Para verificar os containers:

```bash
docker ps
```

Para parar os containers:

```bash
docker compose down
```

---

## ⚙️ Configuração

Crie um arquivo `.env` na raiz do projeto:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=tasks
JWT_SECRET=sua_chave_secreta
```

> ⚠️ O arquivo `.env` não deve ser enviado para o GitHub. Ele está protegido pelo `.gitignore`.

---

## ▶️ Executando localmente

Instale as dependências:

```bash
go mod download
```

Execute os testes:

```bash
go test ./...
```

Execute a API:

```bash
go run .
```

A API ficará disponível em:

```text
http://localhost:8080
```

---

## 🔑 Endpoints

### Autenticação

| Método | Endpoint    | Descrição         |
| ------ | ----------- | ----------------- |
| POST   | `/register` | Cadastrar usuário |
| POST   | `/login`    | Realizar login    |

### Tarefas

| Método | Endpoint      | Descrição        |
| ------ | ------------- | ---------------- |
| GET    | `/tasks`      | Listar tarefas   |
| GET    | `/tasks/{id}` | Buscar tarefa    |
| POST   | `/tasks`      | Criar tarefa     |
| PUT    | `/tasks/{id}` | Atualizar tarefa |
| DELETE | `/tasks/{id}` | Excluir tarefa   |

As rotas de tarefas exigem autenticação através de:

```text
Authorization: Bearer SEU_TOKEN
```

---

## 🧪 Testes

O projeto possui testes automatizados utilizando o pacote de testes nativo do Go.

Execute:

```bash
go test ./...
```

Resultado esperado:

```text
ok github.com/MarcosJose/go-task-api
```

---

## 📂 Estrutura do projeto

```text
go-task-api/
│
├── auth/
│   ├── jwt.go
│   └── middleware.go
│
├── handlers/
│   ├── auth_handler.go
│   └── task_handler.go
│
├── models/
│   └── task.go
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── main_test.go
└── .gitignore
```

---

## 🎯 Objetivo do projeto

Este projeto foi desenvolvido como parte do meu **portfólio de Backend**, com foco no desenvolvimento de APIs utilizando Go.

O objetivo é demonstrar conhecimentos práticos em:

* Desenvolvimento de APIs REST
* Go
* PostgreSQL
* SQL
* Autenticação JWT
* Segurança de senhas
* Middleware
* Docker
* Docker Compose
* Testes automatizados
* Git e GitHub
* Organização de projetos Backend

---

## 👨‍💻 Autor

**Marcos José**

Desenvolvedor em formação com foco em **Backend, Go, APIs REST, bancos de dados e desenvolvimento de software**.

### Contato

* GitHub: [MarcosJose-SN](https://github.com/MarcosJose-SN)
* LinkedIn: [Marcos José](https://www.linkedin.com/)

---

⭐ Se este projeto foi útil ou interessante, considere deixar uma estrela no repositório.
