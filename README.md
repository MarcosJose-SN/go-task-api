# Go Task API

API REST para gerenciamento de tarefas, desenvolvida em Go com PostgreSQL e Docker.

## Tecnologias

- Go
- PostgreSQL
- Docker
- REST API
- Git/GitHub

## Funcionalidades

- Criar tarefas
- Listar tarefas
- Atualizar tarefas
- Excluir tarefas
- Validação de dados
- Persistência em PostgreSQL

## Endpoints

| Método | Endpoint | Descrição |
|---|---|---|
| GET | `/tasks` | Lista todas as tarefas |
| POST | `/tasks` | Cria uma tarefa |
| PUT | `/tasks/{id}` | Atualiza uma tarefa |
| DELETE | `/tasks/{id}` | Exclui uma tarefa |

## Como executar

```bash
docker build -t go-task-api .
docker run --name go-task-api --network go-network -p 8080:8080 go-task-api