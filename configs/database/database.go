package database

import (
	//"os"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// DB is reusable gorm sql connection.
	DB *gorm.DB
)

// ConnectDB connects this application to database instance.
func ConnectDB() error {
	h := "mariadb"
	u := "reihan"
	pwd := "reihan"
	p := "3306"
	d := "ktp"

	// h := "localhost"
	// u := "reihan"
	// pwd := "reihan"
	// p := "3311"
	// d := "ktp"

	dsn := u + ":" + pwd + "@tcp(" + h + ":" + p + ")/" + d + "?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = dbConnection
	fmt.Print("Succes Connect Database !! \n")
	return nil

}
