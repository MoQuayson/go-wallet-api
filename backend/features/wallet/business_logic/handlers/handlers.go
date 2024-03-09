package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	shared "go-wallet-api/features/shared/models"
	"go-wallet-api/features/shared/utils/enums"
	"go-wallet-api/features/wallet/business_logic/app/models"
	"log"
	"net/http"
)

// GetAllWalletsHandler Get All Wallets
func GetAllWalletsHandler(ctx *fiber.Ctx) error {

	s := getWalletService()
	wallets, err := s.FindAllWallets()
	if err != nil {
		return ctx.Status(500).JSON(&shared.APIResponse{
			Code:    500,
			Message: enums.ResponseMsg_GetWalletErr,
		})
	}

	return ctx.Status(200).JSON(&shared.APIResponse{
		Code:    200,
		Message: enums.ResponseMsg_GetWalletSuccess,
		Data:    wallets,
	})
}

func GetWalletByIdHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	s := getWalletService()
	wallet, err := s.FindWalletById(id)
	if err != nil {
		return ctx.Status(500).JSON(&shared.APIResponse{
			Code:    500,
			Message: enums.ResponseMsg_GetWalletErr,
		})
	}

	if wallet == nil {

		return ctx.Status(404).JSON(&shared.APIResponse{
			Code:    404,
			Message: enums.ResponseMsg_WalletNotFound,
			Data:    nil,
		})
	}

	return ctx.Status(200).JSON(&shared.APIResponse{
		Code:    200,
		Message: enums.ResponseMsg_GetWalletSuccess,
		Data:    wallet,
	})
}

func CreateWalletHandler(ctx *fiber.Ctx) error {
	req := &models.WalletRequest{}
	var wallet *models.Wallet
	var err error

	if err = ctx.BodyParser(req); err != nil {
		return ctx.Status(500).JSON(&shared.APIResponse{
			Code:    500,
			Message: enums.ResponseMsg_CreateWalletErr,
		})
	}

	s := getWalletService()

	walletCount, err := s.GetWalletsCount(req.Owner)

	if err != nil {
		return ctx.Status(500).JSON(&shared.APIResponse{
			Code:    500,
			Message: enums.ResponseMsg_CreateWalletErr,
		})
	}

	if *walletCount >= int64(MaxWalletCount) {
		return ctx.Status(http.StatusBadRequest).JSON(&shared.APIResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Cannot have more than %v wallets", MaxWalletCount),
		})
	}

	if wallet, err = s.CreateNewWallet(req); err != nil {
		return ctx.Status(500).JSON(&shared.APIResponse{
			Code:    500,
			Message: enums.ResponseMsg_CreateWalletErr,
		})
	}

	return ctx.Status(201).JSON(&shared.APIResponse{
		Code:    201,
		Message: enums.ResponseMsg_CreateWalletSuccess,
		Data:    wallet,
	})
}

func UpdateWalletHandler(ctx *fiber.Ctx) error {
	req := &models.WalletRequest{}
	walletId := ctx.Params("id")

	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(500).JSON(&shared.APIResponse{
			Code:    500,
			Message: enums.ResponseMsg_UpdateWalletErr,
		})
	}

	s := getWalletService()
	wallet, err := s.FindWalletById(walletId)
	if err != nil {
		log.Printf("%s in UpdateWalletHandler function", err)
	}

	if wallet == nil {

		return ctx.Status(404).JSON(&shared.APIResponse{
			Code:    404,
			Message: enums.ResponseMsg_WalletNotFound,
			Data:    nil,
		})
	}

	if wallet, err = s.UpdateWallet(walletId, req); err != nil {
		log.Printf("UpdateWallet Error: %s", err.Error())
		return ctx.Status(500).JSON(&shared.APIResponse{
			Code:    500,
			Message: enums.ResponseMsg_UpdateWalletErr,
		})
	}

	return ctx.Status(200).JSON(&shared.APIResponse{
		Code:    200,
		Message: enums.ResponseMsg_UpdateWalletSuccess,
		Data:    wallet,
	})
}

// DeleteWalletHandler for delete wallet route
func DeleteWalletHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	s := getWalletService()

	if err := s.DeleteWallet(id); err != nil {
		return ctx.Status(500).JSON(&shared.APIResponse{
			Code:    500,
			Message: enums.ResponseMsg_DeleteWalletErr,
		})
	}

	return ctx.Status(200).JSON(&shared.APIResponse{
		Code:    200,
		Message: enums.ResponseMsg_DeleteWalletSuccess,
	})
}
