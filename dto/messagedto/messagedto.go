package messagedto

type SendMessageRequest struct {
	TickerStatus string `json:"tickerStatus"`
}

type MessageRecipientsResponse struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
	Content     string `json:"content"`
	Status      string `json:"status"`
}
