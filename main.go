package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/MarcosJose/go-task-api/auth"
	"github.com/MarcosJose/go-task-api/handlers"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/MarcosJose/go-task-api/docs"
	_ "github.com/lib/pq"
)

var db *sql.DB

// @title Go Task API
// @version 1.0
// @description API REST para gerenciamento de tarefas com autenticação JWT.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		fmt.Println("Aviso: não foi possível carregar o .env:", err)
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

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

	tasksHandler := handlers.TasksHandler(db)

	http.Handle("/tasks", auth.AuthMiddleware(tasksHandler))
	http.Handle("/tasks/", auth.AuthMiddleware(tasksHandler))
	http.Handle("/login", handlers.LoginHandler(db))
	http.Handle("/register", handlers.RegisterHandler(db))
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro:", err)
	}
}
