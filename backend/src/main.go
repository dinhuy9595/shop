package main

import (
	"fmt"
	"log"

	"github.com/dinhuy9595/shop/config"
	"github.com/dinhuy9595/shop/domain/model"
	"github.com/dinhuy9595/shop/infrastructure/datastore"
	"github.com/dinhuy9595/shop/infrastructure/router"
	"github.com/dinhuy9595/shop/registry"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Brand{})
	r := registry.NewRegistry(db)

	e := echo.New()
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
