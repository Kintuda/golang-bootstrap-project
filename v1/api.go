package v1

import (
	"github.com/Kintuda/golang-bootstrap-project/app"
	"github.com/gin-gonic/gin"
)

type API struct {
	Account    *gin.RouterGroup
	Operation  *gin.RouterGroup
	Currencies *gin.RouterGroup
}

func NewAPI(s *app.Server) {
	s.Router.POST("accounts", CreateAccount)
}
