package messaging

import (
	"encoding/json"
	"io"
	"log"
)

func SaveMessage(body io.ReadCloser){
	var message Message

	json.NewDecoder(body).Decode(&message)
	
	log.Printf("\n{ID: %s, Body: %s, Sender: %s}\n", message.ID, message.Body, message.Sender)
}