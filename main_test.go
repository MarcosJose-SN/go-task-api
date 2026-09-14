package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/MarcosJose/go-task-api/auth"
	"github.com/MarcosJose/go-task-api/handlers"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func TestTasksHandler(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal("Erro ao carregar .env:", err)
	}

	db, err = sql.Open(
		"postgres",
		"host="+os.Getenv("DB_HOST")+
			" port="+os.Getenv("DB_PORT")+
			" user="+os.Getenv("DB_USER")+
			" password="+os.Getenv("DB_PASSWORD")+
			" dbname="+os.Getenv("DB_NAME")+
			" sslmode=disable",
	)

	if err != nil {
		t.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		t.Fatal("Não foi possível conectar ao PostgreSQL:", err)
	}

	token, err := auth.GenerateToken("marcos")
	if err != nil {
		t.Fatal("Erro ao gerar token:", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()

	handler := auth.AuthMiddleware(handlers.TasksHandler(db))
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("esperado status 200, obtido %d", rec.Code)
	}
}
