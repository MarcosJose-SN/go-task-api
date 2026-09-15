# 🚀 Go Task API

API REST desenvolvida em **Go** para gerenciamento de tarefas, com autenticação de usuários, JWT, PostgreSQL, Docker e documentação interativa com Swagger/OpenAPI.

Projeto de portfólio focado em demonstrar conhecimentos práticos de **Backend com Go**, APIs REST, banco de dados, autenticação, segurança, containerização, testes e documentação.

---

## 🛠️ Tecnologias

- **Go 1.27**
- **net/http**
- **PostgreSQL 16**
- **database/sql**
- **lib/pq**
- **JWT** — `github.com/golang-jwt/jwt/v5`
- **bcrypt** — `golang.org/x/crypto/bcrypt`
- **godotenv**
- **Docker**
- **Docker Compose**
- **Swagger / OpenAPI**
- **SQL**
- **Git / GitHub**

---

## 📌 Funcionalidades

### 🔐 Autenticação

- Cadastro de usuários
- Login com usuário e senha
- Senhas protegidas com bcrypt
- Geração de tokens JWT
- Validação de tokens JWT através de middleware
- Proteção das rotas de tarefas
- Autenticação utilizando `Authorization: Bearer`

### 📋 Gerenciamento de tarefas

- Listar tarefas
- Criar tarefas
- Atualizar tarefas
- Excluir tarefas
- Validação de dados
- Tratamento de erros
- Respostas HTTP apropriadas

### 📚 Documentação

- Documentação da API utilizando Swagger/OpenAPI
- Visualização dos endpoints
- Teste dos endpoints diretamente pelo navegador
- Documentação da autenticação JWT
- Documentação dos parâmetros e respostas

---

## 🗄️ Banco de dados

O projeto utiliza **PostgreSQL 16** para armazenar usuários e tarefas.

### Tabela `users`

```text
users
├── id
├── username
├── email
├── password_hash
└── created_at
Tabela tasks
tasks
├── id
├── title
└── completed


🐳 Docker

O projeto possui um ambiente configurado com Docker Compose, permitindo executar a API e o PostgreSQL em containers.

Arquitetura
┌─────────────────────────┐
│        Cliente          │
│      localhost:8080     │
└────────────┬────────────┘
             │
             ▼
┌─────────────────────────┐
│      Go Task API        │
│      Go + JWT + REST    │
└────────────┬────────────┘
             │
             ▼
┌─────────────────────────┐
│      PostgreSQL 16      │
│        port 5432        │
└─────────────────────────┘
Iniciar o projeto

Com o Docker Desktop iniciado:

docker compose up -d --build
Verificar os containers
docker compose ps
Parar os containers
docker compose down
Visualizar os logs da API
docker compose logs -f api


⚙️ Configuração

Crie um arquivo .env na raiz do projeto:

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=tasks
JWT_SECRET=sua_chave_secreta

⚠️ O arquivo .env contém informações sensíveis e não deve ser enviado para o GitHub. Ele está protegido pelo .gitignore.


▶️ Executando localmente

Para baixar as dependências:

go mod download

Para executar a aplicação:

go run .

A API ficará disponível em:

http://localhost:8080


📚 Swagger / OpenAPI

A API possui documentação interativa utilizando Swagger/OpenAPI.

Depois de iniciar a aplicação, acesse:

http://localhost:8080/swagger/index.html

O Swagger permite:

Visualizar os endpoints
Consultar parâmetros
Visualizar modelos JSON
Autenticar utilizando JWT
Executar requisições
Visualizar respostas HTTP


🔑 Autenticação pelo Swagger
Execute POST /login.
Copie o token JWT retornado.
Clique em Authorize.
Informe:
Bearer SEU_TOKEN_JWT
Execute os endpoints protegidos de /tasks.


🔑 Endpoints
Autenticação
Método	Endpoint	Descrição
POST	/register	Cadastrar usuário
POST	/login	Realizar login
Tarefas
Método	Endpoint	Autenticação	Descrição
GET	/tasks	JWT	Listar tarefas
POST	/tasks	JWT	Criar tarefa
PUT	/tasks/{id}	JWT	Atualizar tarefa
DELETE	/tasks/{id}	JWT	Excluir tarefa

As rotas de tarefas exigem o seguinte header:

Authorization: Bearer SEU_TOKEN_JWT


👤 Cadastro de usuário
POST /register

Exemplo de requisição:

{
  "username": "marcos",
  "email": "marcos@email.com",
  "password": "123456"
}

A senha é convertida em hash utilizando bcrypt antes de ser armazenada no banco de dados.


🔐 Login
POST /login

Exemplo de requisição:

{
  "username": "marcos",
  "password": "123456"
}

A API retorna um token JWT após a autenticação:

{
  "token": "SEU_TOKEN_JWT"
}


➕ Criar tarefa
POST /tasks

Header:

Authorization: Bearer SEU_TOKEN_JWT

Body:

{
  "title": "Estudar Go",
  "completed": false
}


📋 Listar tarefas
GET /tasks

Header:

Authorization: Bearer SEU_TOKEN_JWT

Exemplo de resposta:

[
  {
    "id": 1,
    "title": "Estudar Go",
    "completed": false
  }
]


✏️ Atualizar tarefa
PUT /tasks/{id}

Exemplo:

PUT /tasks/1

Body:

{
  "title": "Estudar Go e PostgreSQL",
  "completed": true
}


🗑️ Excluir tarefa
DELETE /tasks/{id}

Exemplo:

DELETE /tasks/1

Resposta:

204 No Content


🧪 Testes

O projeto possui testes automatizados utilizando o pacote nativo de testes do Go.

Execute:

go test ./...

Resultado esperado:

ok      github.com/MarcosJose/go-task-api


📖 Gerar documentação Swagger

Caso as anotações da API sejam alteradas, a documentação pode ser regenerada utilizando:

go run github.com/swaggo/swag/cmd/swag init

Os arquivos gerados ficam na pasta:

docs/
├── docs.go
├── swagger.json
└── swagger.yaml
📂 Estrutura do projeto
go-task-api/
│
├── auth/
│   ├── jwt.go
│   └── middleware.go
│
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
│
├── handlers/
│   ├── auth_handler.go
│   └── task_handler.go
│
├── models/
│   └── task.go
│
├── .gitignore
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── main_test.go
└── README.md

O arquivo .env também existe localmente, mas não é versionado no GitHub por estar protegido pelo .gitignore.

🔄 Fluxo da aplicação
                    ┌──────────────┐
                    │    Cliente   │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │   HTTP API   │
                    └──────┬───────┘
                           │
                           ▼
                  ┌─────────────────┐
                  │ JWT Middleware  │
                  └────────┬────────┘
                           │
                    Token válido?
                       │       │
                      NÃO     SIM
                       │       │
                       ▼       ▼
                      401   Handlers
                               │
                               ▼
                         PostgreSQL
                               │
                               ▼
                         JSON Response


📊 Status HTTP utilizados
Código	Significado
200	Requisição processada com sucesso
201	Recurso criado
204	Requisição processada sem conteúdo
400	Requisição inválida
401	Não autorizado
404	Recurso não encontrado
405	Método não permitido
500	Erro interno do servidor


🎯 Objetivo do projeto

Este projeto foi desenvolvido como parte do meu portfólio de Backend, com foco no desenvolvimento de APIs utilizando Go.

O objetivo é demonstrar conhecimentos práticos em:

Desenvolvimento de APIs REST
Go
PostgreSQL
SQL
CRUD
Autenticação JWT
Segurança de senhas com bcrypt
Middleware
Docker
Docker Compose
Swagger/OpenAPI
Testes automatizados
Variáveis de ambiente
Git e GitHub
Organização de projetos Backend


🚧 Possíveis evoluções

Como próximos passos, o projeto pode receber:

Testes unitários mais abrangentes
Testes de integração
Migrations
Paginação
Filtros e ordenação
Relacionamento entre usuários e tarefas
Refresh Token
WebSockets
Logs estruturados
CI/CD
Deploy em cloud


## 👨‍💻 Autor

**Marcos José**

Desenvolvedor em formação com foco em **Backend, Go, APIs REST, bancos de dados e desenvolvimento de software**.

Também possuo experiência com programação, desenvolvimento de jogos digitais e tecnologias educacionais.

### 🌐 Contato

- GitHub: https://github.com/MarcosJose-SN
- LinkedIn: https://www.linkedin.com/in/marcos-jose-380915389/
- [Google](https://google.com).
---

⭐ Se este projeto foi útil ou interessante, considere deixar uma estrela no repositório.