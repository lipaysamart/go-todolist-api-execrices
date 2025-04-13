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
	itemGroup := r.Group("/item")
	{
		itemGroup.POST("", taskHandle.AddItem)
		itemGroup.POST(":id", taskHandle.UpdateItem)
		itemGroup.GET("", taskHandle.GetItemList)
		itemGroup.GET(":id", taskHandle.GetItem)
		itemGroup.DELETE("", taskHandle.DeleteItem)
	}
}
