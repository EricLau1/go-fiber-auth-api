package routes

import (
	"go-fiber-auth-api/controllers"

	"github.com/gofiber/fiber/v2"
)

type authRoutes struct {
	authController controllers.AuthController
}

func NewAuthRoutes(authController controllers.AuthController) Routes {
	return &authRoutes{authController}
}

func (r *authRoutes) Install(app *fiber.App) {
	app.Post("/signup", r.authController.SignUp)
	app.Post("/signin", r.authController.SignIn)
	app.Get("/users", AuthRequired, r.authController.GetUsers)
	app.Get("/users/:id", AuthRequired, r.authController.GetUser)
	app.Put("/users/:id", AuthRequired, r.authController.PutUser)
	app.Delete("/users/:id", AuthRequired, r.authController.DeleteUser)
}
