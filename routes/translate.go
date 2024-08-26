package routes

import (
	"SinalizeB/constants"
	"SinalizeB/controllers"
	"github.com/labstack/echo/v4"
)

func Translate(app *echo.Group) {
	group := app.Group(string(constants.TranslateRoute))
	controller := controllers.NewTranslateController()

	group.POST("", controller.Create)
	group.PUT("/:id", controller.Update)
	group.DELETE("/:id", controller.Delete)
	group.GET("", controller.GetAll)
	group.GET("/:id", controller.Get)
}
