package repository

import (
	"customer_module/model"
	"database/sql"

	_ "github.com/lib/pq"
)

// Creating table
func TableCreation(db *sql.DB) error {
	customer := `
	CREATE TABLE IF NOT EXISTS customers (
		customer_id INT PRIMARY KEY,
		firstName VARCHAR(255), 
		lastName VARCHAR(255), 
		mobileNo VARCHAR(16), 
		email VARCHAR(255) UNIQUE,
		password VARCHAR(255)
	);
	CREATE TABLE IF NOT EXISTS token_authentications (
		user_id INT REFERENCES customers(customer_id),
		user_email VARCHAR(255) UNIQUE,
		token VARCHAR(255)
	);
`
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = db.Exec(customer)
	if err != nil {
		panic(err)
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// Adding the new user
func CreateUser(db *sql.DB, data model.Customer) error {
	_, err := db.Exec("INSERT INTO customers (customer_id, firstName, lastName, mobileNo, email, password) VALUES ($1, $2, $3, $4, $5, $6)", data.CustomerId, data.FirstName, data.LastName, data.MobileNo, data.Email, data.Password)
	if err != nil {
		return err
	}
	return nil
}

// Retrieving the User details
func FindUser(db *sql.DB, email string) (customerData model.Customer, err error) {
	rows, err := db.Query("SELECT * FROM customers WHERE email = $1", email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&customerData.CustomerId, &customerData.FirstName, &customerData.LastName, &customerData.MobileNo, &customerData.Email, &customerData.Password); err != nil {
			panic(err)
		}
	}
	return customerData, err
}

// Adding the token
func AddToken(db *sql.DB, auth model.TokenAuthentication) error {
	_, err := db.Exec("INSERT INTO token_authentications (user_id, user_email, token) VALUES ($1, $2, $3)", auth.UserId, auth.Email, auth.Token)
	if err != nil {
		return err
	}
	return nil
}

// Retrieving the login users by token
func FindUsersByToken(db *sql.DB, token string) (tokenData model.TokenAuthentication, err error) {
	rows, err := db.Query("SELECT * FROM token_authentications WHERE token = $1", token)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&tokenData.UserId, &tokenData.Email, &tokenData.Token); err != nil {
			panic(err)
		}
	}
	return tokenData, err
}
