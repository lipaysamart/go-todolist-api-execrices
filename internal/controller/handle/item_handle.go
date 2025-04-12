package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lipaysamart/go-todolist-api-execrices/internal/model"
	"github.com/lipaysamart/go-todolist-api-execrices/internal/service"
)

type TaskHandle struct {
	taskHandle service.ITaskService
}

func NewTaskHandle(handle service.ITaskService) *TaskHandle {
	return &TaskHandle{
		taskHandle: handle,
	}
}

func (h *TaskHandle) AddItem(ctx *gin.Context) {
	var itemReq model.ItemReq
	if err := ctx.ShouldBind(&itemReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind request body",
			"error":   err.Error(),
		})
		return
	}

	if err := h.taskHandle.AddItem(ctx, &itemReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to Create item",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created item",
	})
}

func (h *TaskHandle) GetItemList(ctx *gin.Context) {

}

func (h *TaskHandle) GetItem(ctx *gin.Context) {

}

func (h *TaskHandle) UpdateItem(ctx *gin.Context) {

}
func (h *TaskHandle) DeleteItem(ctx *gin.Context) {

}
