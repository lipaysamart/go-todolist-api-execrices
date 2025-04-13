package route

import (
	"github.com/gin-gonic/gin"
	"github.com/lipaysamart/go-todolist-api-execrices/internal/controller/handle"
	"github.com/lipaysamart/go-todolist-api-execrices/internal/repository"
	"github.com/lipaysamart/go-todolist-api-execrices/internal/service"
	"github.com/lipaysamart/go-todolist-api-execrices/pkg/db"
)

func UserRoute(r *gin.RouterGroup, db db.IDatabase) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandle := handle.NewUserHandle(userService)
	userGroup := r.Group("/auth")
	{
		userGroup.POST("/register", userHandle.Register)
		userGroup.POST("/login", userHandle.Login)
		userGroup.POST("/profile/:id", userHandle.UpdateProfile)
	}
}
