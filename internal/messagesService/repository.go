package messagesService

import "gorm.io/gorm"

// MessageRepository интерфейс, описывающий методы для работы с БД
type MessageRepository interface {
	CreateMessage(message Message) (Message, error)
	GetAllMessages() ([]Message, error)
	UpdateMessageByID(id int, message Message) (Message, error)
	DeleteMessageByID(id int) error
}

// Структура messageRepository, которая будет содержать подключение к БД
type messageRepository struct {
	db *gorm.DB
}

// NewMessageRepository создает новый репозиторий для работы с сообщениями
func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

// Реализация методов интерфейса MessageRepository

// CreateMessage создает новое сообщение в БД
func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

// GetAllMessages возвращает все сообщения из БД
func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}

// UpdateMessageByID обновляет сообщение с заданным id
func (r *messageRepository) UpdateMessageByID(id int, message Message) (Message, error) {
	var existingMessage Message
	if err := r.db.First(&existingMessage, id).Error; err != nil {
		return Message{}, err // возвращаем ошибку, если сообщение не найдено
	}

	// Обновляем поля существующего сообщения
	existingMessage.Content = message.Content
	existingMessage = message

	if err := r.db.Save(&existingMessage).Error; err != nil {
		return Message{}, err
	}

	return existingMessage, nil
}

// DeleteMessageByID удаляет сообщение с заданным id
func (r *messageRepository) DeleteMessageByID(id int) error {
	if err := r.db.Delete(&Message{}, id).Error; err != nil {
		return err
	}
	return nil
}
