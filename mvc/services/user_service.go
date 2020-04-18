package services

import (
	"github.com/novdov/goms/mvc/domain"
	"github.com/novdov/goms/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
