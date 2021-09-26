package datastore

import (
	"flag"
	"fmt"

	"github.com/dinhuy9595/shop/config"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	username := config.C.Database.User
	password := config.C.Database.Password
	dbName := config.C.Database.DBName
	dbHost := config.C.Database.Addr
	dbport := config.C.Database.Port
	var nFlag string
	flag.StringVar(&nFlag, "n", "hello", "help message for flag n")
	flag.Parse()
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbport, username, dbName, password) //Build connection string
	fmt.Println(dbUri)
	fmt.Println(nFlag)
	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}
	return db
}
