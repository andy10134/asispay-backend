package routes

import (
	"github.com/andy10134/asisPay-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func UseAuthRoute(app *fiber.App) {

	auth := app.Group("/auth")
	auth.Post("/login", controllers.LoginUser)
	auth.Post("/register", controllers.RegisterUser)
	auth.Get("/logout", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Ok"})
	})

}
