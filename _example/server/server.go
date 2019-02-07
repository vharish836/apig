package server

import (
	"github.com/vharish836/apig/_example/middleware"
	"github.com/vharish836/apig/_example/router"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Setup ...
func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.SetDBtoContext(db))
	router.Initialize(r)
	return r
}
