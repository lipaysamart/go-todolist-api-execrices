package route

import (
	"github.com/gin-gonic/gin"
	"github.com/lipaysamart/go-todolist-api-execrices/internal/controller/handle"
	"github.com/lipaysamart/go-todolist-api-execrices/internal/repository"
	"github.com/lipaysamart/go-todolist-api-execrices/internal/service"
	"github.com/lipaysamart/go-todolist-api-execrices/pkg/db"
)

func TaskRoute(r *gin.RouterGroup, db db.IDatabase) {
	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandle := handle.NewTaskHandle(taskService)
	r.Group("/item")
	{
		r.POST("/", taskHandle.AddItem)
		r.POST("/:id", taskHandle.UpdateItem)
		r.GET("/", taskHandle.GetItemList)
		r.GET("/:id", taskHandle.GetItem)
		r.DELETE("/", taskHandle.DeleteItem)
	}
}
