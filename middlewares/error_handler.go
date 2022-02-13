package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			// log, handle, etc.
			fmt.Println(err)

		}

		c.JSON(http.StatusInternalServerError, "")
	}
}
