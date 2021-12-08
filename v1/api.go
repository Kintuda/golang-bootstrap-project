package v1

import (
	"github.com/Kintuda/golang-bootstrap-project/app"
)

func NewAPI(s *app.Server) {
	s.Router.GET("/status", HeathCheck)
}
