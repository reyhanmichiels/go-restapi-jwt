package main

import (
	"jwt-restApi/src/business/repository"
	"jwt-restApi/src/business/usecase"
	"jwt-restApi/src/handler/rest"
	"jwt-restApi/src/sdk/auth"
	"jwt-restApi/src/sdk/database/sql"
	"jwt-restApi/src/sdk/env"
)

func init() {
	env.LoadEnv()
	sql.ConnectDB()
	sql.Migrate()
}

func main() {
	auth := auth.Inject()
	repo := repository.Inject(sql.DB)
	usecase := usecase.Inject(repo, auth)
	handler := rest.Inject(usecase)

	handler.Run()
}
