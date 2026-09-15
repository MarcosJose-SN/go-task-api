package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/MarcosJose/go-task-api/models"
)

// TasksHandler direciona as requisições para cada operação do CRUD.
func TasksHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			TasksGetHandler(db)(w, r)

		case http.MethodPost:
			TasksPostHandler(db)(w, r)

		case http.MethodPut:
			TasksPutHandler(db)(w, r)

		case http.MethodDelete:
			TasksDeleteHandler(db)(w, r)

		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	}
}

// TasksGetHandler godoc
// @Summary Listar tarefas
// @Description Retorna todas as tarefas cadastradas.
// @Tags Tarefas
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Task
// @Failure 401 {string} string
// @Failure 500 {string} string
// @Router /tasks [get]
func TasksGetHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		rows, err := db.Query(
			"SELECT id, title, completed FROM tasks ORDER BY id",
		)

		if err != nil {
			http.Error(
				w,
				"Erro ao buscar tarefas",
				http.StatusInternalServerError,
			)
			return
		}

		defer rows.Close()

		var tasks []models.Task

		for rows.Next() {
			var task models.Task

			if err := rows.Scan(
				&task.ID,
				&task.Title,
				&task.Completed,
			); err != nil {
				http.Error(
					w,
					"Erro ao ler tarefa",
					http.StatusInternalServerError,
				)
				return
			}

			tasks = append(tasks, task)
		}

		json.NewEncoder(w).Encode(tasks)
	}
}

// TasksPostHandler godoc
// @Summary Criar tarefa
// @Description Cria uma nova tarefa.
// @Tags Tarefas
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param task body models.Task true "Dados da tarefa"
// @Success 201 {object} models.Task
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Failure 500 {string} string
// @Router /tasks [post]
func TasksPostHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var task models.Task

		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(
				w,
				"JSON inválido",
				http.StatusBadRequest,
			)
			return
		}

		if strings.TrimSpace(task.Title) == "" {
			http.Error(
				w,
				"O título é obrigatório",
				http.StatusBadRequest,
			)
			return
		}

		err := db.QueryRow(
			"INSERT INTO tasks (title, completed) VALUES ($1, $2) RETURNING id",
			task.Title,
			task.Completed,
		).Scan(&task.ID)

		if err != nil {
			http.Error(
				w,
				"Erro ao criar tarefa",
				http.StatusInternalServerError,
			)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
	}
}

// TasksPutHandler godoc
// @Summary Atualizar tarefa
// @Description Atualiza uma tarefa existente.
// @Tags Tarefas
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID da tarefa"
// @Param task body models.Task true "Dados da tarefa"
// @Success 200 {object} models.Task
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /tasks/{id} [put]
func TasksPutHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		idText := strings.TrimPrefix(r.URL.Path, "/tasks/")
		id, err := strconv.Atoi(idText)

		if err != nil {
			http.Error(
				w,
				"ID inválido",
				http.StatusBadRequest,
			)
			return
		}

		var task models.Task

		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(
				w,
				"JSON inválido",
				http.StatusBadRequest,
			)
			return
		}

		if strings.TrimSpace(task.Title) == "" {
			http.Error(
				w,
				"O título é obrigatório",
				http.StatusBadRequest,
			)
			return
		}

		result, err := db.Exec(
			"UPDATE tasks SET title = $1, completed = $2 WHERE id = $3",
			task.Title,
			task.Completed,
			id,
		)

		if err != nil {
			http.Error(
				w,
				"Erro ao atualizar tarefa",
				http.StatusInternalServerError,
			)
			return
		}

		rowsAffected, _ := result.RowsAffected()

		if rowsAffected == 0 {
			http.Error(
				w,
				"Tarefa não encontrada",
				http.StatusNotFound,
			)
			return
		}

		task.ID = id

		json.NewEncoder(w).Encode(task)
	}
}

// TasksDeleteHandler godoc
// @Summary Excluir tarefa
// @Description Exclui uma tarefa pelo ID.
// @Tags Tarefas
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID da tarefa"
// @Success 204
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /tasks/{id} [delete]
func TasksDeleteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idText := strings.TrimPrefix(r.URL.Path, "/tasks/")
		id, err := strconv.Atoi(idText)

		if err != nil {
			http.Error(
				w,
				"ID inválido",
				http.StatusBadRequest,
			)
			return
		}

		result, err := db.Exec(
			"DELETE FROM tasks WHERE id = $1",
			id,
		)

		if err != nil {
			http.Error(
				w,
				"Erro ao excluir tarefa",
				http.StatusInternalServerError,
			)
			return
		}

		rowsAffected, _ := result.RowsAffected()

		if rowsAffected == 0 {
			http.Error(
				w,
				"Tarefa não encontrada",
				http.StatusNotFound,
			)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
