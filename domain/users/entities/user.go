package entities

import (
	"time"

	"github.com/NicklasWallgren/go-template/domain/common"
)

// User model.
type User struct {
	common.Entity
	// ID        common.PrimaryID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       uint8     `json:"age"`
	Birthday  time.Time `json:"time"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"` // auditable
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated"`   // auditable
}

func NewUser(name string, email string, age uint8, birthday time.Time) User {
	return User{Name: name, Email: email, Age: age, Birthday: birthday}
}

func NewUserWithID(id common.PrimaryID, name string, email string, age uint8, birthday time.Time) User {
	return User{Entity: common.Entity{ID: id}, Name: name, Email: email, Age: age, Birthday: birthday}
}

// TableName gives table name of model.
func (u User) TableName() string {
	return "users"
}
