package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/MarcosJose/go-task-api/auth"
	"github.com/MarcosJose/go-task-api/models"
)

func TasksHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		username, ok := r.Context().Value(auth.UsernameKey).(string)

		if !ok || username == "" {
			http.Error(w, "Usuário não identificado", http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		switch r.Method {

		case http.MethodGet:
			rows, err := db.Query("SELECT id, title, completed FROM tasks ORDER BY id")
			if err != nil {
				http.Error(w, "Erro ao buscar tarefas", http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			var tasks []models.Task

			for rows.Next() {
				var task models.Task

				if err := rows.Scan(&task.ID, &task.Title, &task.Completed); err != nil {
					http.Error(w, "Erro ao ler tarefa", http.StatusInternalServerError)
					return
				}

				tasks = append(tasks, task)
			}

			json.NewEncoder(w).Encode(tasks)

		case http.MethodPost:
			var task models.Task

			if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
				http.Error(w, "JSON inválido", http.StatusBadRequest)
				return
			}

			if strings.TrimSpace(task.Title) == "" {
				http.Error(w, "O título é obrigatório", http.StatusBadRequest)
				return
			}

			err := db.QueryRow(
				"INSERT INTO tasks (title, completed) VALUES ($1, $2) RETURNING id",
				task.Title,
				task.Completed,
			).Scan(&task.ID)

			if err != nil {
				http.Error(w, "Erro ao criar tarefa", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(task)

		case http.MethodPut:
			idText := strings.TrimPrefix(r.URL.Path, "/tasks/")
			id, err := strconv.Atoi(idText)

			if err != nil {
				http.Error(w, "ID inválido", http.StatusBadRequest)
				return
			}

			var task models.Task

			if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
				http.Error(w, "JSON inválido", http.StatusBadRequest)
				return
			}

			if strings.TrimSpace(task.Title) == "" {
				http.Error(w, "O título é obrigatório", http.StatusBadRequest)
				return
			}

			result, err := db.Exec(
				"UPDATE tasks SET title = $1, completed = $2 WHERE id = $3",
				task.Title,
				task.Completed,
				id,
			)

			if err != nil {
				http.Error(w, "Erro ao atualizar tarefa", http.StatusInternalServerError)
				return
			}

			rowsAffected, _ := result.RowsAffected()

			if rowsAffected == 0 {
				http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
				return
			}

			task.ID = id
			json.NewEncoder(w).Encode(task)

		case http.MethodDelete:
			idText := strings.TrimPrefix(r.URL.Path, "/tasks/")
			id, err := strconv.Atoi(idText)

			if err != nil {
				http.Error(w, "ID inválido", http.StatusBadRequest)
				return
			}

			result, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)

			if err != nil {
				http.Error(w, "Erro ao excluir tarefa", http.StatusInternalServerError)
				return
			}

			rowsAffected, _ := result.RowsAffected()

			if rowsAffected == 0 {
				http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusNoContent)

		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	}
}
