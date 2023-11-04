package driver

import (
	h "customer_module/helper"
	repo "customer_module/repository"
	"log"

	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func DbConnection() *gorm.DB {
	//Loading the .env file
	if err := h.ConfigEnv(".env"); err != nil {
		fmt.Println(h.Err, err)
	}
	host := os.Getenv(h.HostName)
	port := os.Getenv(h.PortNum)
	user := os.Getenv(h.DbUserName)
	password := os.Getenv(h.DbPassword)
	dbname := os.Getenv(h.DbName)
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname= %s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(h.Database, connectionUrl)
	if err != nil {
		panic(err)
	}
	log.Println(h.DbSuccessMsg)
	repo.TableCreation(db)
	return db
}
