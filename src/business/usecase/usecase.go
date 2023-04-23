package usecase

import (
	"jwt-restApi/src/business/repository"
	"jwt-restApi/src/sdk/auth"
)

type Usecase struct {
	User UserUsecase
}

func Inject(r *repository.Repository, auth auth.Auth) *Usecase {
	return &Usecase{User: NewUserUsecase(r.User, auth)}
}
