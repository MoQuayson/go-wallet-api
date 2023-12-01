package models

import (
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	EMAIL string = "EMAIL"
	PHONE string = "PHONE"
)

type User struct {
	ID        uuid.UUID    `gorm:"column:id;type:uuid;primaryKey" json:"id" form:"id"`
	Name      string       `gorm:"column:name;size:255" json:"name" form:"name"`
	Email     string       `gorm:"column:email;size:255;unique" json:"email" form:"email"`
	PhoneNum  NullString   `gorm:"column:phone_num;size:255;unique" json:"phone_num" form:"phone_num"`
	Role      string       `gorm:"column:role;size:255" json:"role" form:"role"`
	Password  string       `gorm:"column:password;size:255" json:"-" form:"-"`
	Token     string       `gorm:"-" json:"token,omitempty" form:"token,omitempty"`
	CreatedAt time.Time    `gorm:"column:created_at;type:timestamp" json:"created_at,omitempty" form:"created_at"`
	UpdatedAt NullDateTime `gorm:"column:updated_at;type:timestamp" json:"updated_at,omitempty" form:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.ID, _ = uuid.NewV4()
	u.CreatedAt = time.Now()
	u.Password = string(hash)

	return nil
}
func (u *User) BeforeSave(tx *gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}

// func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
// 	if u.Role == "Admin" {
// 	  return errors.New("admin user not allowed to delete")
// 	}
// 	return
//   }

func HashPassword(pwd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash)
}
