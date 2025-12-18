package routes

import (
	"go-backend-task/internal/handler"
	"go-backend-task/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, repo *repository.Repository) { // *fiber.App is similar to &mut FiberApp in rust
	userHandler := handler.NewUserHandler(repo)

	users := app.Group("/users")
	users.Post("/", userHandler.CreateUser)
	users.Get("/:id", userHandler.GetUser)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)
	users.Get("/", userHandler.ListUsers)
}
