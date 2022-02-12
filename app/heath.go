package app

import (
	"net/http"
	"time"

	"github.com/Kintuda/golang-bootstrap-project/models"
	"github.com/gin-gonic/gin"
)

func HeathCheck(c *gin.Context) {
	heathCheck := &models.HeathCheckStatus{Status: "ok", Time: time.Now()}
	c.JSON(http.StatusOK, heathCheck)
}
