package foosball

import (
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	ID       int64 `gorm:"primary_key"`
	Username string
	Email    string
	Password string
}

type CreateUserRequest struct {
	Username string
	Email    string
	Password string
}

type CreateUserResponse struct{}

func CreateUser(request CreateUserRequest, db *gorm.DB) (CreateUserResponse, error) {
	user := User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	db.Create(&user)

	errors := db.GetErrors()
	if len(errors) != 0 {
		log.Errorf("unable to create user %s: %v", request.Username, errors)
		return CreateUserResponse{}, db.Error
	}

	return CreateUserResponse{}, nil
}
