package main

import (
	"fmt"

	"github.com/arturbaccarin/bands-website-back/pkg/database"
	"github.com/arturbaccarin/bands-website-back/pkg/routes"
)

// init start db configs
func init() {
	database.Config{}.Start()
	routes.Config{}.Start()
}

func main() {
	db := database.DB_GORM

	fmt.Println(db.Config)
}
