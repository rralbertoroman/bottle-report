package handlers

import (
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/rralbertoroman/bottle-report/internal/app"
	"gorm.io/gorm"
)

type Message struct {
	ID        string         `json:"id" gorm:"column:id;primaryKey"`
	Body      string         `json:"body" gorm:"column:body"`
	Sender    string         `json:"sender_id" gorm:"column:sender"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func SaveMessage(ctx context.Context,body io.ReadCloser, a *app.App) (err error){
	var message Message

	json.NewDecoder(body).Decode(&message)

	err = gorm.G[Message](a.DB).Create(ctx, &message)
	return
}

func AllMessages(ctx context.Context, a *app.App) (messages []Message, err error){
	result := a.DB.Find(&messages)
	err = result.Error
	return
}
