package di

import (
	"database/sql"
	"goapi/app/domain/service"
	"goapi/app/handler"
	"goapi/app/infrastructure/mysql"
)

func InitUser(db *sql.DB) handler.IUserHandler {
	r := mysql.NewUserRepository(db)
	s := service.NewUserService(r)
	return handler.NewUserHandler(s)

}