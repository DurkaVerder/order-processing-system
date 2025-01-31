package postgres

const (
	addOrderQuery     = "INSERT INTO orders (user_id, total_amount, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	getOrderQuery     = "SELECT * FROM orders WHERE id = $1"
	getAllOrdersQuery = "SELECT * FROM orders WHERE user_id = $1"
	deleteOrderQuery  = "DELETE FROM orders WHERE id = $1"
	GetUserEmailQuery = "SELECT email FROM users WHERE id = $1"
)
