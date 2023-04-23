package sql

import "jwt-restApi/src/business/entity"

func Migrate() {
	DB.AutoMigrate(&entity.User{})
}
