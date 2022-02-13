package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUser struct {
	Name string `json:"name"`
}

func ValidateUser(c *gin.Context) {
	var payload CreateUser

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusOK, payload)
}
