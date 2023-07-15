package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	"uji/helpers"
)

type User struct {
	ID          uint32        `json:"id" gorm:"primarykey"`
	Username    string        `json:"username" gorm:"NOT NULL;unique;type:varchar(255);" valid:"required"`
	Email       string        `json:"email" gorm:"NOT NULL;unique;type:varchar(255);" valid:"required,email"`
	Password    string        `json:"password,omitempty" gorm:"NOT NULL;type:text;" valid:"required,minstringlength(6)"`
	Age         uint          `json:"age" gorm:"NOT NULL;type:integer;" valid:"required,range(8|100)"`
	SocialMedia []SocialMedia `json:"social_media,omitempty" gorm:"foreignKey:user_id;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Photo       []Photo       `json:"photos,omitempty" gorm:"foreignKey:user_id;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comment     []Comment     `json:"comments,omitempty" gorm:"foreignKey:user_id;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type UserUpdateInput struct {
	Username  string    `json:"username" gorm:"NOT NULL;unique;type:varchar(255);"`
	Email     string    `json:"email" gorm:"NOT NULL;unique;type:varchar(255);" valid:"email"`
	Password  string    `json:"password,omitempty" gorm:"NOT NULL;type:text;" valid:"minstringlength(6)"`
	Age       uint      `json:"age" gorm:"NOT NULL;type:integer;" valid:"range(8|100)"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().ID()
	u.Password = helpers.HashPassword(u.Password)
	return nil
}

type UserUseCase interface {
	UserRegisterUc(user *User) error
	UserLoginUc(user *User) error
	GetUserByIdUc(id uint32) (*User, error)
	GetUsersUc(user *[]User) (*[]User, error)
	UpdateUserUc(id uint32, user *User) (*User, error)
	DeleteUserUc(id uint32) error
}

type UserRepository interface {
	UserRegisterRepository(user *User) error
	UserLoginRepository(user *User) error
	GetUserByIdRepository(id uint32) (*User, error)
	GetUsersRepository(user *[]User) (*[]User, error)
	UpdateUserRepository(id uint32, user *User) (*User, error)
	DeleteUserRepository(id uint32) error
}
