package handlers

import (
	"fmt"
	"go-wallet-api/config"
	"go-wallet-api/models"
	"go-wallet-api/requests"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// Get All Wallets
func GetAllWalletsHandler(ctx *fiber.Ctx) error {

	Wallets, err := models.GetAllWallets(config.DbCtx)
	if err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: "Something went wrong when retrieving wallet data",
		})
	}

	return ctx.Status(200).JSON(&models.APIResponse{
		Code:    200,
		Message: "Wallets retrieved successfully",
		Data:    Wallets,
	})
}

func GetWalletByIdHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	wallet, err := models.GetWalletById(id, config.DbCtx)
	if err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: "Something went wrong when retrieving wallet data",
		})
	}

	if wallet.ID.IsNil() {

		return ctx.Status(404).JSON(&models.APIResponse{
			Code:    404,
			Message: "Wallet does not exist!",
			Data:    nil,
		})
	}

	return ctx.Status(200).JSON(&models.APIResponse{
		Code:    200,
		Message: "Wallet retrieved successfully",
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
			Message: "Something went wrong when creating user",
		})
	}

	walletCount := models.GetUserWalletCount(payload.Owner, config.DbCtx)

	if walletCount >= models.MAX_WALLET_COUNT {
		return ctx.Status(http.StatusBadRequest).JSON(&models.APIResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Cannot have more than %v wallets", models.MAX_WALLET_COUNT),
		})
	}

	//return nil

	if wallet, err = models.CreateNewWallet(payload, config.DbCtx); err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: "Something went wrong when creating wallet",
		})
	}

	return ctx.Status(200).JSON(&models.APIResponse{
		Code:    200,
		Message: "Wallet added successfully",
		Data:    wallet,
	})
}

func UpdateWalletHandler(ctx *fiber.Ctx) error {
	payload := requests.WalletRequest{}
	walletId := ctx.Params("id")
	var wallet models.Wallet
	var err error

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: "Something went wrong when updating wallet",
		})
	}

	wallet, err = models.GetWalletById(walletId, config.DbCtx)
	if err != nil {
		log.Errorf("%s in UpdateWalletHandler function", err)
	}

	if wallet.ID.IsNil() {

		return ctx.Status(404).JSON(&models.APIResponse{
			Code:    404,
			Message: "Wallet does not exist!",
			Data:    nil,
		})
	}

	if wallet, err = models.UpdateWallet(walletId, payload, config.DbCtx); err != nil {
		return ctx.Status(500).JSON(&models.APIResponse{
			Code:    500,
			Message: "Something went wrong when updating wallet",
		})
	}

	return ctx.Status(200).JSON(&models.APIResponse{
		Code:    200,
		Message: "Wallet updated successfully",
		Data:    wallet,
	})
}
