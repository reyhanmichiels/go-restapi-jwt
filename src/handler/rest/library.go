package rest

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (r *rest) ErrorResponse(c *gin.Context, code int, message string, err error) {
	c.JSON(code, gin.H{
		"status": "error",
		"message": message,
		"error": err.Error(),
	})
}

func (r *rest) SuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{
		"status": "succesed",
		"message": message,
		"data": data,
	})
}

func (r *rest) HashPassword(password string) string {
	bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(password)
}