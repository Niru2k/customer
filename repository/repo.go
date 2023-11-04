package repository

import (
	"customer_module/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Creating table
func TableCreation(Db *gorm.DB) {
	Db.AutoMigrate(&model.Customer{}, &model.TokenAuthentication{})
}

// Adding the new user
func CreateUser(Db *gorm.DB, data model.Customer) error {
	return Db.Create(&data).Error
}

// Retrieve a token by email
func ReadTokenByEmail(Db *gorm.DB, user model.Customer) (auth model.TokenAuthentication, err error) {
	err = Db.Where("email=?", user.Email).Find(&auth).Error
	return auth, err
}

// Retrieving the User details
func FindUser(Db *gorm.DB, email string) (data model.Customer, err error) {
	err = Db.Where("email=?", email).Find(&data).Error
	return data, err
}

// Adding the token
func AddToken(Db *gorm.DB, auth model.TokenAuthentication) error {
	err := Db.Create(&auth).Error
	return err
}
