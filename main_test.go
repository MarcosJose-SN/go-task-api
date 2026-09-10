package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MarcosJose/go-task-api/handlers"
	_ "github.com/lib/pq"
)

func TestTasksHandler(t *testing.T) {
	var err error

	db, err = sql.Open(
		"postgres",
		"host=localhost port=5432 user=postgres password=123456 dbname=tasks sslmode=disable",
	)
	if err != nil {
		t.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		t.Fatal("Não foi possível conectar ao PostgreSQL:", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec := httptest.NewRecorder()

	handler := handlers.TasksHandler(db)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("esperado status 200, obtido %d", rec.Code)
	}
}
