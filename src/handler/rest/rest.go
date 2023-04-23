package rest

import (
	"jwt-restApi/src/business/usecase"
	"jwt-restApi/src/sdk/middleware"

	"github.com/gin-gonic/gin"
)

type Rest interface{
	Run()
}

type rest struct {
	uc *usecase.Usecase
	gin *gin.Engine
}

func Inject(usecase *usecase.Usecase) Rest {
	r := &rest{
		uc: usecase,
		gin: gin.Default(),
	}

	r.Route()
	return r
}

func (r *rest) Run() {
	r.gin.Run()
}

func (r *rest) Route() {
	v1 := r.gin.Group("/api/v1")
	
	v1.GET("/", middleware.Auth ,func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	v1.POST("/regist", r.Registration)
	v1.POST("/login", r.Login)
	v1.GET("/loginCheck", middleware.Auth, r.LoginCheck)
}