package db

import (
	"log"
	"os"
	"strings"

	"github.com/vharish836/api-server/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/serenize/snaker"
)

// Connect ...
func Connect() *gorm.DB {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil
	}

	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	db.LogMode(false)

	if gin.IsDebugging() {
		db.LogMode(true)
	}

	if os.Getenv("AUTOMIGRATE") == "1" {
		db.AutoMigrate(
			&models.User{},
		)
	}

	return db
}

// Instance ...
func Instance(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
}

// SetPreloads ...
func (self *Parameter) SetPreloads(db *gorm.DB) *gorm.DB {
	if self.Preloads == "" {
		return db
	}

	for _, preload := range strings.Split(self.Preloads, ",") {
		var a []string

		for _, s := range strings.Split(preload, ".") {
			a = append(a, snaker.SnakeToCamel(s))
		}

		db = db.Preload(strings.Join(a, "."))
	}

	return db
}
