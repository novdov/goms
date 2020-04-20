package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/novdov/goms/mvc/services"
	"github.com/novdov/goms/mvc/utils"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    fmt.Sprintf("user %v not found", userId),
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
	}

	utils.Respond(c, http.StatusOK, user)
}
