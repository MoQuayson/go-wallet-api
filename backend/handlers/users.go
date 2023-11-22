package handlers

import (
	"go-wallet-api/config"
	"go-wallet-api/models"
	"go-wallet-api/requests"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// Get All Users
func GetAllUsersHandler(ctx *fiber.Ctx) error {

	users, err := models.GetAllUsers(config.DbCtx)
	if err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: "Something went wrong when retrieving user data",
		})
	}

	return ctx.Status(200).JSON(&models.APIResponse{
		Code:    200,
		Message: "Users retrieved successfully",
		Data:    users,
	})
}

func GetUserByIdHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user, err := models.GetUserById(id, config.DbCtx)
	if err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: "Something went wrong when retrieving user data",
		})
	}

	if user.ID.IsNil() {

		return ctx.Status(404).JSON(&models.APIResponse{
			Code:    404,
			Message: "User does not exist!",
			Data:    nil,
		})
	}

	return ctx.Status(200).JSON(&models.APIResponse{
		Code:    200,
		Message: "User retrieved successfully",
		Data:    user,
	})
}

func CreateUserHandler(ctx *fiber.Ctx) error {
	payload := requests.UserRequest{}
	var user models.User
	var err error

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: "Something went wrong when creating user",
		})
	}

	if user, err = models.CreateNewUser(payload, config.DbCtx); err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: "Something went wrong when creating user",
		})
	}

	return ctx.Status(200).JSON(&models.APIResponse{
		Code:    200,
		Message: "User created successfully",
		Data:    user,
	})
}

func UpdateUserHandler(ctx *fiber.Ctx) error {
	payload := requests.UserRequest{}
	userId := ctx.Params("id")
	var user models.User
	var err error

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: "Something went wrong when updating user",
		})
	}

	user, err = models.GetUserById(userId, config.DbCtx)
	if err != nil {
		log.Errorf("%s in UpdateUserHandler function", err)
	}

	if user.ID.IsNil() {

		return ctx.Status(404).JSON(&models.APIResponse{
			Code:    404,
			Message: "User does not exist!",
			Data:    nil,
		})
	}

	if user, err = models.UpdateUser(userId, payload, config.DbCtx); err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: "Something went wrong when updating user",
		})
	}

	return ctx.Status(200).JSON(&models.APIResponse{
		Code:    200,
		Message: "User updated successfully",
		Data:    user,
	})
}
