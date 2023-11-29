package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Wallet struct {
	ID            uuid.UUID    `gorm:"column:id;type:uuid;primaryKey" json:"id" form:"id"`
	Name          string       `gorm:"column:name;size:255" json:"name" form:"name"`
	Type          string       `gorm:"column:type;size:255" json:"type" form:"type"`
	AccountNumber string       `gorm:"column:account_number;size:255;" json:"account_number" form:"account_number"`
	AccountScheme string       `gorm:"column:account_scheme;size:255" json:"account_scheme" form:"account_scheme"`
	Owner         uuid.UUID    `gorm:"column:owner;type:uuid" json:"owner" form:"owner"`
	User          User         `gorm:"foreignKey:Owner;references:ID" json:"user,omitempty"`
	CreatedAt     time.Time    `gorm:"column:created_at;type:timestamp" json:"created_at" form:"created_at"`
	UpdatedAt     NullDateTime `gorm:"column:updated_at;type:timestamp" json:"updated_at" form:"updated_at"`
}

func (Wallet) TableName() string {
	return "wallets"
}

func (w *Wallet) BeforeCreate(tx *gorm.DB) error {
	w.ID, _ = uuid.NewV4()
	w.CreatedAt = time.Now()

	return nil
}
