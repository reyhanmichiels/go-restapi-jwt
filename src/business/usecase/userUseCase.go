package usecase

import (
	"jwt-restApi/src/business/entity"
	"jwt-restApi/src/business/repository"
	"jwt-restApi/src/sdk/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	Create(userRegist entity.UserRegistration) (entity.UserResponse, int, string, error)
	Login(userLogin entity.UserLogin, c *gin.Context) (entity.UserResponse, int, string, error)
	LoginCheck(c *gin.Context) (entity.UserResponse, int, string, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
	auth           auth.Auth
}

func NewUserUsecase(userRepository repository.UserRepository, auth auth.Auth) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		auth:           auth,
	}
}

func (u *userUsecase) Create(userRegist entity.UserRegistration) (entity.UserResponse, int, string, error) {
	pass, err := u.auth.HashPassword(userRegist.Password)
	if err != nil {
		return entity.UserResponse{}, 500, "Failed hash password", err
	}

	user := entity.User{
		Name:     userRegist.Name,
		Email:    userRegist.Email,
		Password: pass,
	}

	if err := u.userRepository.Create(user); err != nil {
		return entity.UserResponse{}, 409, "Failed create new user", err
	}

	return userToUserResponse(user), 202, "Successfully registered user", nil
}

func (u *userUsecase) Login(userLogin entity.UserLogin, c *gin.Context) (entity.UserResponse, int, string, error) {
	user, err := u.userRepository.FindUserByEmail(userLogin.Email)
	if err != nil {
		return entity.UserResponse{}, 409, "User not found", err
	}

	if err := u.auth.VerifyPassword(user.Password, userLogin.Password); err != nil {
		return entity.UserResponse{}, 401, "Invalid Password", err
	}

	token, err := u.auth.SetJwtToken(user)
	if err != nil {
		return entity.UserResponse{}, 500, "Failed generate jwt token", err
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("jwt-token", token, (3600 * 24), "", "", false, true)

	return userToUserResponse(user), 200, "Successfully login", nil
}

func (u *userUsecase) LoginCheck(c *gin.Context) (entity.UserResponse, int, string, error) {
	user, _ := c.Get("user")
	return userToUserResponse(user.(entity.User)), 200, "Successfully find login user", nil
}

func userToUserResponse(user entity.User) entity.UserResponse {
	return entity.UserResponse{
		Name:  user.Name,
		Email: user.Email,
	}
}
