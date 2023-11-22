package models

import (
	"go-wallet-api/requests"
	"time"

	"github.com/gofiber/fiber/v2/log"
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
	CreatedAt time.Time    `gorm:"column:created_at;type:timestamp" json:"created_at" form:"created_at"`
	UpdatedAt NullDateTime `gorm:"column:updated_at;type:timestamp" json:"updated_at" form:"updated_at"`
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

func HashPassword(pwd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash)
}

// Gets all users from DB
func GetAllUsers(db *gorm.DB) ([]User, error) {
	users := []User{}

	err := db.Find(&users).Error
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

// Gets user by id from DB
func GetUserById(id string, db *gorm.DB) (User, error) {
	user := User{}

	err := db.Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return User{}, err
	}

	return user, nil
}

func CreateNewUser(req requests.UserRequest, db *gorm.DB) (User, error) {
	u := User{
		Name:      req.Name,
		Email:     req.Email,
		PhoneNum:  nullString.FromString(req.PhoneNum),
		Role:      "User",
		Password:  "password",
		CreatedAt: time.Now(),
	}
	if err := db.Create(&u).Error; err != nil {
		log.Error(err)
		return u, err
	}

	log.Info("User created")
	return u, nil
}

// Updates user
func UpdateUser(id string, req requests.UserRequest, db *gorm.DB) (User, error) {
	user := User{}
	var err error

	if user, err = GetUserById(id, db); err != nil {
		return user, err
	}

	user.Name = req.Name
	user.Email = req.Email
	user.PhoneNum = nullString.FromString(req.PhoneNum)

	err = db.Save(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// Check whether user has exists via phone no.
func CheckIfUserPhoneNumExists(email string, db *gorm.DB) bool {
	var count int64
	db.Where("email = ?", email).Count(&count)

	return count >= 1
}

// Check whether user exists via email
func CheckIfUserEmailExists(phoneNum string, db *gorm.DB) bool {
	var count int64
	db.Where("phone_num = ?", phoneNum).Count(&count)

	return count >= 1
}

// Function to check if user exists
func (u *User) UserExists(db *gorm.DB) bool {
	return false
}
