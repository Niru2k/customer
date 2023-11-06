package model

type Customer struct {
	CustomerId int    `json:"customer_id" validate:"required"`
	FirstName  string `json:"firstName" validate:"required"`
	LastName   string `json:"lastName" validate:"required"`
	MobileNo   string `json:"mobileNo" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
}

type TokenAuthentication struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

func (*Customer) TableName() string {
	return "customers"
}

func (*TokenAuthentication) TableName() string {
	return "token_authentications"
}
