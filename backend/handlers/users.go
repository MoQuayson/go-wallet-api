package handlers

import (
	"go-wallet-api/models"
	"go-wallet-api/requests"
	"go-wallet-api/services"

	"github.com/gofiber/fiber/v2"
)

// Get All Users
func GetAllUsersHandler(ctx *fiber.Ctx) error {
	s := services.GetUserService()
	users, err := s.FindAll()
	if err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: models.GET_USER_ERR,
		})
	}

	return ctx.Status(200).JSON(&models.APIResponse{
		Code:    200,
		Message: models.GET_USER_SUCCESS,
		Data:    users,
	})
}

// Get user by Id
func GetUserByIdHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	s := services.GetUserService()
	user, err := s.FindById(id)
	if err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: models.GET_USER_ERR,
		})
	}

	if user.ID.IsNil() {

		return ctx.Status(404).JSON(&models.APIResponse{
			Code:    404,
			Message: models.USER_NOT_FOUND,
			Data:    nil,
		})
	}

	return ctx.Status(200).JSON(&models.APIResponse{
		Code:    200,
		Message: models.GET_USER_SUCCESS,
		Data:    user,
	})
}

func CreateUserHandler(ctx *fiber.Ctx) error {
	payload := requests.UserRequest{}
	var user models.User
	var err error

	s := services.GetUserService()

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: models.CREATE_USER_ERR,
		})
	}

	if user, err = s.CreateNewUser(payload); err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: models.CREATE_USER_ERR,
		})
	}

	return ctx.Status(201).JSON(&models.APIResponse{
		Code:    201,
		Message: models.CREATE_USER_SUCCESS,
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
			Message: models.UPDATE_USER_ERR,
		})
	}

	s := services.GetUserService() //get service

	user, _ = s.FindById(userId)

	//if user does not exist
	if user.ID.IsNil() {

		return ctx.Status(404).JSON(&models.APIResponse{
			Code:    404,
			Message: models.USER_NOT_FOUND,
			Data:    nil,
		})
	}

	//assign data
	user.Name = payload.Name
	user.Email = payload.Email
	user.PhoneNum = models.ConvertToNullString(payload.PhoneNum)

	if user, err = s.UpdateUser(userId, user); err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: models.UPDATE_USER_ERR,
		})
	}

	return ctx.Status(200).JSON(&models.APIResponse{
		Code:    200,
		Message: models.UPDATE_USER_SUCCESS,
		Data:    user,
	})
}

// Function for delete user route
func DeleteUserHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	s := services.GetUserService()

	if err := s.DeleteUser(id); err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: models.DELETE_USER_ERR,
		})
	}

	return ctx.Status(200).JSON(&models.APIResponse{
		Code:    200,
		Message: models.DELETE_USER_SUCCESS,
	})
}
