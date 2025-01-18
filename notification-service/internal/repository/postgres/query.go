package postgres

const (
	SelectUserEmailByOrderId = `SELECT email FROM users WHERE id = (SELECT user_id FROM orders WHERE id = $1)`
)
