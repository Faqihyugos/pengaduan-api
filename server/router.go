package server

import (
	"github.com/faqihyugos/pengaduan-api/delivery/http"
	"github.com/faqihyugos/pengaduan-api/middleware"
	repositoryuser "github.com/faqihyugos/pengaduan-api/repositories/repositoryUser"
	"github.com/faqihyugos/pengaduan-api/services/serviceuser"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {
	repoUser := repositoryuser.New(db)
	srvUser := serviceuser.New(repoUser)
	deliUser := http.NewUser(srvUser)
	routeUser := r.Group("/users")

	// route user
	routeUser.POST("/register", deliUser.Create)
	routeUser.POST("/login", deliUser.Login)
	routeUser.PUT("", middleware.Authorization, deliUser.Update)
	routeUser.DELETE("", middleware.Authorization, deliUser.DeleteByID)
}
