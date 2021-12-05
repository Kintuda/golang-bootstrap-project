package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"ok": "ok"})
}
