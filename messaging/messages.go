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


func AllMessages() []Message{
	return []Message{
		{
			ID: "123456",
			Body: "De Barbosa a Ceguera 150 pesos",
			Sender: "chuchi_lamora",
		},{
			ID: "122556",
			Body: "Novia del Mediodía a Frank País 150",
			Sender: "miguelTanquepoporeso",
		},{
			ID: "127856",
			Body: "Frank País a Ceguera 150 CUP",
			Sender: "marcos",
		},
	}
}