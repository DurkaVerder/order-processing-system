package postgres

const (
	updateStatusQuery       = "UPDATE orders SET status = $1 WHERE id = $2"
	createRecordStatusQuery = "INSERT INTO order_status_history (order_id, status) VALUES ($1, $2)"
)
