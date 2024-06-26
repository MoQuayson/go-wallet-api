package entities

import (
	"go-wallet-api/features/shared/utils"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserEntity struct {
	ID        uuid.UUID  `gorm:"column:id;type:uuid;primaryKey" json:"id" form:"id"`
	Name      string     `gorm:"column:name;size:255" json:"name" form:"name"`
	Email     string     `gorm:"column:email;size:255;unique" json:"email" form:"email"`
	PhoneNum  *string    `gorm:"column:phone_num;size:255;unique" json:"phone_num" form:"phone_num"`
	Role      string     `gorm:"column:role;size:255" json:"role" form:"role"`
	Password  string     `gorm:"column:password;size:255" json:"-" form:"-"`
	Token     string     `gorm:"-" json:"token,omitempty" form:"token,omitempty"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp" json:"created_at,omitempty" form:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at,omitempty" form:"updated_at"`
}

func (*UserEntity) TableName() string {
	return "users"
}

func (u *UserEntity) BeforeCreate(tx *gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if u.ID.IsNil() {
		u.ID = utils.NewUUID()
	}
	u.CreatedAt = time.Now()
	u.Password = string(hash)

	return nil
}
func (u *UserEntity) BeforeSave(tx *gorm.DB) error {
	//hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	return err
	//}
	//
	//u.Password = string(hash)

	return nil
}

// func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
// 	if u.Role == "Admin" {
// 	  return errors.New("admin user not allowed to delete")
// 	}
// 	return
//   }
