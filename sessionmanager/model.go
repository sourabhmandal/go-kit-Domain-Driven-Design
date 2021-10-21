package sessionmanager

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
 * This is the model that defines the
 * structure of Session data in the Database
 * i.e
 * This is the Schema of Sessions Table
 */
type Session struct {
	gorm.Model
	UserID uuid.UUID `gorm:"unique" json:"userId"`
}