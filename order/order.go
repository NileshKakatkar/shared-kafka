package order

const (
	PlacedTopic = "shopcore.order-events"
	PlacedType  = "order.placed"
)

type Placed struct {
	EventID     string `json:"event_id"`
	OrderID     string `json:"order_id"`
	UserID      string `json:"user_id"`
	Email       string `json:"email"`
	TotalAmount int64  `json:"total_amount"`
	Currency    string `json:"currency"`
}
