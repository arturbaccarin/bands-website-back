package database

import (
	"log"
	"os"

	"github.com/arturbaccarin/bands-website-back/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB_GORM *gorm.DB

// Config has all db configurations
type Config struct{}

// Start the connection
func (c Config) Start() {
	DB_GORM = c.connect()
}

// setup configs gorm
func (Config) setup() (dsn string) {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	dsn = utils.ConcatenateStrings(dbUser, ":", dbPass, "@tcp(localhost:3306)/bands?charset=utf8&parseTime=True&loc=Local")

	return
}

// connect with db
func (c Config) connect() (db *gorm.DB) {
	dsn := c.setup()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("failed to connect database")
	}

	return
}
