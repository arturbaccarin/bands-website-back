package main

import (
	"fmt"

	"github.com/arturbaccarin/bands-website-back/pkg/database"
)

// init start db configs
func init() {
	database.Config{}.Start()
}

func main() {
	db := database.DB_GORM

	fmt.Println(db.Config)
}
