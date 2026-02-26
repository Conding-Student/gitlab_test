package routes

import (
	"intern_template_v1/controller"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {

	// CREATE YOUR ENDPOINTS HERE
	app.Get("getnames", controller.Get_names)
	// --------------------------

	app.Post("insertnames", controller.Insert_name)
}
