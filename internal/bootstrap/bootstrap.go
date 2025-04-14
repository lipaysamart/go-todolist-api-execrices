package bootstrap

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lipaysamart/go-todolist-api-execrices/internal/controller/route"
	"github.com/lipaysamart/go-todolist-api-execrices/pkg/config"
	"github.com/lipaysamart/go-todolist-api-execrices/pkg/db"
)

type BootStrap struct {
	engin    *gin.Engine
	database db.IDatabase
	cfg      *config.Schema
}

func NewBootStrap(db db.IDatabase) *BootStrap {
	return &BootStrap{
		engin:    gin.Default(),
		database: db,
		cfg:      config.GetConfig(),
	}
}

func (b *BootStrap) Run() error {
	_ = b.engin.SetTrustedProxies(nil)

	if b.cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := b.MapRoutes(); err != nil {
		return err
	}

	if err := b.engin.Run(fmt.Sprintf(":%v", b.cfg.HttpPort)); err != nil {
		return err
	}

	return nil
}

func (b *BootStrap) MapRoutes() error {
	v1 := b.engin.Group("/api/v1")
	route.UserRoute(v1, b.database)
	route.TaskRoute(v1, b.database)
	return nil
}
