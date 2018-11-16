package foosball

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
}

func CreateUser(db *gorm.DB) error {
	return nil
}
