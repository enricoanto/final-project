package helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Error(c *gin.Context, code int, err error) {
	message := err.Error()
	if caseOject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range caseOject {
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required", err.Field())
			case "email":
				message = fmt.Sprintf("%s must filled email format", err.Field())
			case "min":
				message = fmt.Sprintf("%s must greater than or equal %v", err.Field(), err.Param())
			case "max":
				message = fmt.Sprintf("%s must lower than or equal %v", err.Field(), err.Param())
			}
		}
	}
	c.JSON(code, gin.H{"message": message})
	return
}

func Success(c *gin.Context, code int, response interface{}) {
	c.JSON(code, response)
	return
}
