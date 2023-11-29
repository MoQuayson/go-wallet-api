package handlers

import (
	"fmt"
	"go-wallet-api/models"
	"go-wallet-api/repositories"
	"go-wallet-api/requests"
	"go-wallet-api/services"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// Get All Wallets
func GetAllWalletsHandler(ctx *fiber.Ctx) error {

	s := services.GetWalletService()
	Wallets, err := s.FindAll()
	if err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: models.GET_WALLET_SUCCESS,
		})
	}

	return ctx.Status(200).JSON(&models.APIResponse{
		Code:    200,
		Message: models.GET_WALLET_SUCCESS,
		Data:    Wallets,
	})
}

func GetWalletByIdHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	s := services.GetWalletService()
	wallet, err := s.FindById(id)
	if err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: models.GET_WALLET_ERR,
		})
	}

	if wallet.ID.IsNil() {

		return ctx.Status(404).JSON(&models.APIResponse{
			Code:    404,
			Message: models.WALLET_NOT_FOUND,
			Data:    nil,
		})
	}

	return ctx.Status(200).JSON(&models.APIResponse{
		Code:    200,
		Message: models.GET_WALLET_SUCCESS,
		Data:    wallet,
	})
}

func CreateWalletHandler(ctx *fiber.Ctx) error {
	payload := requests.WalletRequest{}
	var wallet models.Wallet
	var err error

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: models.CREATE_WALLET_ERR,
		})
	}

	s := services.GetWalletService()

	walletCount := s.WalletsCount(payload.Owner)

	if walletCount >= models.MAX_WALLET_COUNT {
		return ctx.Status(http.StatusBadRequest).JSON(&models.APIResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Cannot have more than %v wallets", models.MAX_WALLET_COUNT),
		})
	}

	//return nil

	if wallet, err = s.CreateNewWallet(payload); err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: models.CREATE_WALLET_ERR,
		})
	}

	return ctx.Status(201).JSON(&models.APIResponse{
		Code:    201,
		Message: models.CREATE_WALLET_SUCCESS,
		Data:    wallet,
	})
}

func UpdateWalletHandler(ctx *fiber.Ctx) error {
	payload := requests.WalletRequest{}
	walletId := ctx.Params("id")

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: models.UPDATE_WALLET_ERR,
		})
	}

	s := services.GetWalletService()
	wallet, err := s.FindById(walletId)
	if err != nil {
		log.Errorf("%s in UpdateWalletHandler function", err)
	}

	if wallet.ID.IsNil() {

		return ctx.Status(404).JSON(&models.APIResponse{
			Code:    404,
			Message: models.WALLET_NOT_FOUND,
			Data:    nil,
		})
	}

	wallet.Name = fmt.Sprintf("%s %s Wallet", payload.AccountScheme, strings.ToUpper(payload.Type))
	wallet.Type = payload.Type
	wallet.AccountNumber = repositories.TrimAccountNumber(payload.Type, payload.AccountNumber)
	wallet.AccountScheme = payload.AccountScheme

	if wallet, err = s.UpdateWallet(wallet); err != nil {
		log.Errorf("UpdateWallet Error: %s", err.Error())
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: models.UPDATE_WALLET_ERR,
		})
	}

	return ctx.Status(200).JSON(&models.APIResponse{
		Code:    200,
		Message: models.UPDATE_WALLET_SUCCESS,
		Data:    wallet,
	})
}

// Function for delete wallet route
func DeleteWalletHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	s := services.GetWalletService()

	if err := s.DeleteWallet(id); err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: models.DELETE_WALLET_ERR,
		})
	}

	return ctx.Status(200).JSON(&models.APIResponse{
		Code:    200,
		Message: models.DELETE_WALLET_SUCCESS,
	})
}
