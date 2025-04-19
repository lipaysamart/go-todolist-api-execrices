package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lipaysamart/go-todolist-api-exerices/internal/model"
	"github.com/lipaysamart/go-todolist-api-exerices/internal/service"
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
	resp, err := h.taskHandle.GetItemList(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get item list",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully get item list",
		"data":    resp,
	})

}

func (h *TaskHandle) GetItem(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.taskHandle.GetItem(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get item",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully get item",
		"data":    resp,
	})

}

func (h *TaskHandle) UpdateItem(ctx *gin.Context) {
	var itemReq model.ItemReq

	id := ctx.Param("id")
	if err := ctx.ShouldBind(&itemReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to bind request body",
			"error":   err.Error(),
		})
		return
	}

	resp, err := h.taskHandle.UpdateItem(ctx, id, &itemReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update item",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully update item",
		"data":    resp,
	})

}
func (h *TaskHandle) DeleteItem(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.taskHandle.DelItem(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete item",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully delete item",
	})
}
