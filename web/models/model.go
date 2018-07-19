package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Model is the main model interface
type Model interface {
	createNew() (interface{}, error)
}

// Base is the forum-api base model
type Base struct {
	ID        string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// DBError is the database error type
type DBError struct {
	Errors map[string]interface{} `json:"errors"`
}

// Create adds a new record to the database
func Create(m Model) (interface{}, error) {
	i, err := m.createNew()
	return i, err
}

// Error wraps database error info in DBError object
func Error(err error) DBError {
	dbError := DBError{}
	dbError.Errors = make(map[string]interface{})
	dbError.Errors["database"] = err.Error()
	return dbError
}

// BeforeCreate sets Model's primary key value to uuid
func (b *Base) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()

	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return err
	}

	scope.SetColumn("ID", id)
	return nil
}
