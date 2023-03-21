package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type JSON json.RawMessage

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j *JSON) Value() (driver.Value, error) {
	if len(*j) == 0 {
		return nil, nil
	}
	return json.RawMessage(*j).MarshalJSON()
}

// Model Definitions

type User struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	IsActive  bool   `gorm:"default:true"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
}

type Conversation struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	UserID    uint
	User      User
}

type Query struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Query          string `gorm:"not null"`
	Context        int32  `gorm:"default:0"`
	ConversationID uint
	Response       JSON `gorm:"type:jsonb;default:'{}'"`
	Conversation   Conversation
}

type Prompt struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Query     string `gorm:"not null"`
	UserID    uint
	User      User
}
