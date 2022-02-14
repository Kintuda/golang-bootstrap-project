package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Kintuda/golang-bootstrap-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func ErrorHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {

			if _, ok := err.Err.(validator.ValidationErrors); ok {
				var validationErrors []models.ValidationDetails
				errs := err.Err.(validator.ValidationErrors)

				for _, err := range errs {
					var vl models.ValidationDetails
					vl.Field = err.Field()
					vl.Constraint = err.ActualTag()
					vl.Value = fmt.Sprintf("%v", err.Value())
					validationErrors = append(validationErrors, vl)
				}

				c.JSON(http.StatusUnprocessableEntity, gin.H{"details": validationErrors})
				return
			}
			c.JSON(-1, err.Error())
		}
	}
}
