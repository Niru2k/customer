package model

type Customer struct {
	Id        int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	MobileNo  string `json:"mobileNo" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

type TokenAuthentication struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func (*Customer) TableName() string {
	return "customers"
}

func (*TokenAuthentication) TableName() string {
	return "token_authentications"
}
