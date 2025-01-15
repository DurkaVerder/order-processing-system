package postgres

const (
	findUserQuery = "SELECT id FROM users WHERE username = $1 AND password = $2"
	addUserQuery  = "INSERT INTO users (email, login, password, username, created_at, update_at) VALUES ($1, $2, $3, $4, $5, $6)"
)
