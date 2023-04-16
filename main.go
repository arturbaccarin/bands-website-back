package main

import (
	"fmt"

	"github.com/arturbaccarin/bands-website-back/pkg/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("error in load env file")
	}

	db := database.DB_GORM

	fmt.Println(db.Config)
}
