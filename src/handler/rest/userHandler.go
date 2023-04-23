package rest

import (
	"jwt-restApi/src/business/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *rest) Registration(c *gin.Context) {
	var userRegistration entity.UserRegistration
	if err := c.ShouldBindJSON(&userRegistration); err != nil {
		r.ErrorResponse(c, http.StatusConflict, "Send the correct JSON Request", err)
		return
	}

	user, code, message, err := r.uc.User.Create(userRegistration)
	if err != nil {
		r.ErrorResponse(c, code, message, err)
		return
	}

	r.SuccessResponse(c, 202, "Successfully regist account", user)
}

func (r *rest) Login(c *gin.Context) {
	var UserLogin entity.UserLogin
	if err := c.ShouldBindJSON(&UserLogin); err != nil {
		r.ErrorResponse(c, http.StatusConflict, "Send the correct JSON Request", err)
		return
	}

	user, code, message, err := r.uc.User.Login(UserLogin, c)
	if err != nil {
		r.ErrorResponse(c, code, message, err)
		return
	}

	r.SuccessResponse(c, code, message, user)
}

func (r *rest) LoginCheck(c *gin.Context) {
	user, code, message, _ := r.uc.User.LoginCheck(c)
	r.SuccessResponse(c, code, message, user)
}
