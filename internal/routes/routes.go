package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-backend-task/internal/handler"
	"go-backend-task/internal/repository"
)

func SetupRoutes(app *fiber.App, repo *repository.Repository) {
	userHandler := handler.NewUserHandler(repo)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	users := v1.Group("/users")
	users.Post("/", userHandler.CreateUser)
	users.Get("/:id", userHandler.GetUser)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)
	users.Get("/", userHandler.ListUsers)
}
