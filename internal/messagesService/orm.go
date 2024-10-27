package messagesService

import "gorm.io/gorm"

// Message - структура, представляющая сообщение в БД.
type Message struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}
