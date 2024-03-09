package handlers

import (
	shared "go-wallet-api/features/shared/models"
	"go-wallet-api/features/shared/utils/enums"
	"go-wallet-api/features/users/business_logic/app/models"
	userDI "go-wallet-api/features/users/di"
	"log"

	"github.com/gofiber/fiber/v2"
)

// GetAllUsersHandler Gets All Users
func GetAllUsersHandler(ctx *fiber.Ctx) error {
	srv := userDI.WithUserInjector.Service
	users, err := srv.FindAllUsers()
	if err != nil {
		return ctx.Status(500).JSON(&shared.APIResponse{
			Code:    500,
			Message: enums.ResponseMsg_GET_USER_ERR,
		})
	}

	return ctx.Status(200).JSON(&shared.APIResponse{
		Code:    200,
		Message: enums.ResponseMsg_GET_USER_SUCCESS,
		Data:    users,
	})
}

// GetUserByIdHandler Get user by Id
func GetUserByIdHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	srv := userDI.WithUserInjector.Service
	data, err := srv.FindUserById(id)
	if err != nil {
		log.Println(err)
		return ctx.Status(500).JSON(&shared.APIResponse{
			Code:    500,
			Message: enums.ResponseMsg_GET_USER_ERR,
		})
	}

	return ctx.Status(200).JSON(&shared.APIResponse{
		Code:    200,
		Message: enums.ResponseMsg_GET_USER_SUCCESS,
		Data:    data,
	})
}

// CreateUserHandler create new user
func CreateUserHandler(ctx *fiber.Ctx) error {
	req := &models.UserRequest{}

	srv := userDI.WithUserInjector.Service

	if err := ctx.BodyParser(req); err != nil {
		log.Println(err)
		return ctx.Status(500).JSON(&shared.APIResponse{
			Code:    500,
			Message: enums.ResponseMsg_CREATE_USER_ERR,
		})
	}

	user, err := srv.CreateNewUser(req)

	if err != nil {
		return ctx.Status(500).JSON(&shared.APIResponse{
			Code:    500,
			Message: enums.ResponseMsg_CREATE_USER_ERR,
		})
	}

	return ctx.Status(201).JSON(&shared.APIResponse{
		Code:    201,
		Message: enums.ResponseMsg_CREATE_USER_SUCCESS,
		Data:    user,
	})
}

func UpdateUserHandler(ctx *fiber.Ctx) error {
	req := &models.UserRequest{}
	userId := ctx.Params("id")

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(500).JSON(&shared.APIResponse{
			Code:    500,
			Message: enums.ResponseMsg_UPDATE_USER_ERR,
		})
	}

	srv := userDI.WithUserInjector.Service
	user, err := srv.UpdateUser(userId, req)

	if err != nil {
		log.Println(err)
		return returnErrorResponse(ctx, enums.ResponseMsg_GET_USER_ERR)
	}

	return ctx.Status(200).JSON(&shared.APIResponse{
		Code:    200,
		Message: enums.ResponseMsg_UPDATE_USER_SUCCESS,
		Data:    user,
	})
}

// DeleteUserHandler delete user route
func DeleteUserHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	srv := userDI.WithUserInjector.Service

	if err := srv.DeleteUser(id); err != nil {
		log.Println(err)
		return returnErrorResponse(ctx, enums.ResponseMsg_DELETE_USER_ERR)
	}

	return ctx.Status(200).JSON(&shared.APIResponse{
		Code:    200,
		Message: enums.ResponseMsg_DELETE_USER_SUCCESS,
	})
}
