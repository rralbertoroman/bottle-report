package messaging

type Message struct {
	ID     string `json:"id"`
	Body   string `json:"body"`
	Sender string `json:"sender_id"`
	
}