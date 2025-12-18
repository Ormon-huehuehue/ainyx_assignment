package handler

import (
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	db "go-backend-task/db/sqlc"
	"go-backend-task/internal/logger"
	"go-backend-task/internal/models"
	"go-backend-task/internal/repository"
	"go-backend-task/internal/service"
)

type UserHandler struct {
	Repo      *repository.Repository
	Validator *validator.Validate
}

func NewUserHandler(repo *repository.Repository) *UserHandler {
	return &UserHandler{
		Repo:      repo,
		Validator: validator.New(),
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error { // h *Userhandler is the receiver -> similar to &self in Rust
	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Log.Error("Failed to parse request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.Validator.Struct(req); err != nil {
		logger.Log.Warn("Validation failed", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	dob, err := time.Parse("2006-01-02", req.Dob) // go's default date format is year-month-day ( 06 -> year, 01 -> month, 02 -> day)
	if err != nil {
		logger.Log.Error("Failed to parse DOB", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format, use YYYY-MM-DD"})
	}

	arg := db.CreateUserParams{
		Name: req.Name,
		Dob:  dob,
	}

	user, err := h.Repo.Queries.CreateUser(c.Context(), arg)
	if err != nil {
		logger.Log.Error("Failed to create user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	age := service.CalculateAge(user.Dob)

	return c.Status(fiber.StatusCreated).JSON(models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  age,
	})
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id")) // Ascii To Integer
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user, err := h.Repo.Queries.GetUser(c.Context(), int32(id))
	if err != nil {
		logger.Log.Error("Failed to get user", zap.Int("id", id), zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	age := service.CalculateAge(user.Dob)

	return c.JSON(models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  age,
	})
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.Validator.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format"})
	}

	arg := db.UpdateUserParams{
		ID:   int32(id),
		Name: req.Name,
		Dob:  dob,
	}

	user, err := h.Repo.Queries.UpdateUser(c.Context(), arg)
	if err != nil {
		logger.Log.Error("Failed to update user", zap.Int("id", id), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	age := service.CalculateAge(user.Dob)

	return c.JSON(models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  age,
	})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	err = h.Repo.Queries.DeleteUser(c.Context(), int32(id))
	if err != nil {
		logger.Log.Error("Failed to delete user", zap.Int("id", id), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	arg := db.ListUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	users, err := h.Repo.Queries.ListUsers(c.Context(), arg)
	if err != nil {
		logger.Log.Error("Failed to list users", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to list users"})
	}

	total, err := h.Repo.Queries.CountUsers(c.Context())
	if err != nil {
		logger.Log.Error("Failed to count users", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to count users"})
	}

	var userResponses []models.UserResponse
	for _, u := range users {
		userResponses = append(userResponses, models.UserResponse{
			ID:   u.ID,
			Name: u.Name,
			Dob:  u.Dob.Format("2006-01-02"),
			Age:  service.CalculateAge(u.Dob),
		})
	}

	return c.JSON(models.ListUsersResponse{
		Users: userResponses,
		Total: total,
		Page:  page,
		Limit: limit,
	})
}
