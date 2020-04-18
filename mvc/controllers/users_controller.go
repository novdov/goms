package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/novdov/goms/mvc/services"
	"github.com/novdov/goms/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		aplErr := &utils.ApplicationError{
			Message:    fmt.Sprintf("user %v not found", userId),
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		jsonValue, _ := json.Marshal(aplErr)
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write(jsonValue)
	}

	user, aplErr := services.GetUser(userId)
	if aplErr != nil {
		jsonValue, _ := json.Marshal(aplErr)
		resp.WriteHeader(aplErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
