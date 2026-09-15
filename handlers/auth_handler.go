package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/MarcosJose/go-task-api/auth"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterHandler godoc
// @Summary Cadastrar usuário
// @Description Cria um novo usuário com senha criptografada.
// @Tags Autenticação
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Dados do usuário"
// @Success 201 {object} map[string]string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /register [post]
func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var request RegisterRequest

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		if request.Username == "" || request.Email == "" || request.Password == "" {
			http.Error(w, "Usuário, email e senha são obrigatórios", http.StatusBadRequest)
			return
		}

		passwordHash, err := bcrypt.GenerateFromPassword(
			[]byte(request.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			http.Error(w, "Erro ao gerar senha", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec(
			"INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)",
			request.Username,
			request.Email,
			string(passwordHash),
		)

		if err != nil {
			http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Usuário criado com sucesso",
		})
	}
}

// LoginHandler godoc
// @Summary Login do usuário
// @Description Autentica o usuário e retorna um token JWT.
// @Tags Autenticação
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Credenciais de login"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Failure 500 {string} string
// @Router /login [post]
func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var request LoginRequest

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		if request.Username == "" || request.Password == "" {
			http.Error(w, "Usuário e senha são obrigatórios", http.StatusBadRequest)
			return
		}

		var passwordHash string

		err = db.QueryRow(
			"SELECT password_hash FROM users WHERE username = $1",
			request.Username,
		).Scan(&passwordHash)

		if err != nil {
			http.Error(w, "Usuário ou senha inválidos", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword(
			[]byte(passwordHash),
			[]byte(request.Password),
		)

		if err != nil {
			http.Error(w, "Usuário ou senha inválidos", http.StatusUnauthorized)
			return
		}

		token, err := auth.GenerateToken(request.Username)
		if err != nil {
			http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})
	}
}
