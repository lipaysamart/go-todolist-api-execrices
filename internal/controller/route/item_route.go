package route

import (
	"github.com/gin-gonic/gin"
	"github.com/lipaysamart/go-todolist-api-exercise/internal/controller/handle"
	"github.com/lipaysamart/go-todolist-api-exercise/internal/repository"
	"github.com/lipaysamart/go-todolist-api-exercise/internal/service"
	"github.com/lipaysamart/gocommon/dbs"
)

func TaskRoute(r *gin.RouterGroup, db dbs.IDatabase) {
	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandle := handle.NewTaskHandle(taskService)
	itemGroup := r.Group("/item")
	{
		itemGroup.POST("", taskHandle.AddItem)
		itemGroup.POST(":id", taskHandle.UpdateItem)
		itemGroup.GET("", taskHandle.GetItemList)
		itemGroup.GET(":id", taskHandle.GetItem)
		itemGroup.DELETE(":id", taskHandle.DeleteItem)
	}
}
