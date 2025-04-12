package handle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lipaysamart/go-todolist-api-execrices/internal/model"
	"github.com/lipaysamart/go-todolist-api-execrices/internal/service"
)

type UserHandle struct {
	userHandle service.IUserService
}

func NewUserHandle(handle service.IUserService) *UserHandle {
	return &UserHandle{
		userHandle: handle,
	}
}

func (h *UserHandle) Login(ctx *gin.Context) {
	var httpReq model.UserLoginReq

	if err := ctx.ShouldBind(&httpReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed parse request body...",
			"error":   err.Error(),
		})
		return
	}

	user, err := h.userHandle.Login(ctx, &httpReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to login user...",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login success...",
		"data":    user,
	})
}
func (h *UserHandle) Register(ctx *gin.Context) {
	var httpReq model.UserRegisterReq

	if err := ctx.ShouldBind(&httpReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed parse request body...",
			"error":   err.Error(),
		})
		return
	}
	if err := h.userHandle.Register(ctx, &httpReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to register user...",
			"error":   err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Register success...",
	})
}
func (h *UserHandle) UpdateProfile(ctx *gin.Context) {
	var httpReq model.User

	id := ctx.Param("id")
	if err := ctx.ShouldBind(&httpReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed parse request body...",
			"error":   err.Error(),
		})
		return
	}
	user, err := h.userHandle.UpdateProfile(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update user...",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update success...",
		"data":    user,
	})
}
