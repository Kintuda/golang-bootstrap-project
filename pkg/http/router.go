package http

import (
	"github.com/Kintuda/golang-bootstrap-project/pkg/config"
	"github.com/Kintuda/golang-bootstrap-project/pkg/postgres"
	"github.com/gin-gonic/gin"
)

//TODO: add agnostic implementation for database
type Router struct {
	Engine *gin.Engine
	Cfg    *config.AppConfig
	Db     *postgres.Pool
}

func NewRouter(cfg *config.AppConfig, db *postgres.Pool) (*Router, error) {
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	r := Router{
		Engine: router,
		Cfg:    cfg,
		Db:     db,
	}

	return &r, nil
}

func RegisterRoutes(r *Router) {
	
}
