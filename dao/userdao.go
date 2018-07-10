package dao

import (
	"github.com/0pen-source/Carpooling/models"
	"github.com/pkg/errors"
)

// NewDeveloperDAO creates a new mysql database access object
func GetUser(phone string) (models.User, error) {

	var user models.User
	query := "SELECT * FROM `user` where phone = ?"
	err := cache.Get(&user, query, phone)
	return user, errors.WithMessage(err, "fail to get ads")

}
