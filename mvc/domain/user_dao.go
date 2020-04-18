package domain

import (
	"fmt"
	"net/http"

	"github.com/novdov/goms/mvc/utils"
)

var (
	users = map[int64]*User{
		1: {1, "nov", "dov", "master@exampe.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
