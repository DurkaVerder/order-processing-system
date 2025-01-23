package postgres

const (
	findUserQuery = "SELECT id, password FROM users WHERE login = $1"
	addUserQuery  = "INSERT INTO users (email, login, password, username, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
)
