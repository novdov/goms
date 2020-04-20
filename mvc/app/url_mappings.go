package app

import (
	"github.com/novdov/goms/mvc/controllers"
)

func MapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
