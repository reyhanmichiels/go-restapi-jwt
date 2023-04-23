package auth

import (
	"jwt-restApi/src/business/entity"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Auth interface {
	HashPassword(password string) (string, error)
	VerifyPassword(hashPassword string, password string) error
	SetJwtToken(user entity.User) (string, error)
}

type auth struct {
}

func Inject() Auth {
	return &auth{}
}

func (a *auth) HashPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(pass), err
}

func (a *auth) VerifyPassword(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

func (a *auth) SetJwtToken(user entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": user.ID,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_TOKEN")))

	return tokenString, err
}
