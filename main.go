package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/MarcosJose/go-task-api/handlers"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error

	connStr := "host=postgres-go port=5432 user=postgres password=123456 dbname=tasks sslmode=disable"

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Erro ao abrir banco:", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Erro ao conectar ao PostgreSQL:", err)
		return
	}

	fmt.Println("PostgreSQL conectado com sucesso!")

	http.HandleFunc("/tasks", handlers.TasksHandler(db))
	http.HandleFunc("/tasks/", handlers.TasksHandler(db))

	fmt.Println("Servidor rodando em http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro:", err)
	}
}
