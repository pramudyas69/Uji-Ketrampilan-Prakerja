package domain

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	u.ID = uuid.New().ID()
	u.Password = helpers.HashPassword(u.Password)
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}

type UserUseCase interface {
	UserRegisterUc(user *User) error
	UserLoginUc(user *User) error
	GetUserByIdUc(ctx echo.Context) (*User, error)
	GetUsersUc(ctx echo.Context) ([]*User, error)
	UpdateUserUc(ctx echo.Context) (*User, error)
	DeleteUserUc(ctx echo.Context) (*User, error)
}

type UserRepository interface {
	UserRegisterRepository(user *User) error
	UserLoginRepository(user *User) error
	GetUserByIdRepository(ctx echo.Context) (*User, error)
	GetUsersRepository(ctx echo.Context) ([]*User, error)
	UpdateUserRepository(ctx echo.Context) (*User, error)
	DeleteUserRepository(ctx echo.Context) (*User, error)
}
