package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-wallet-api/configs"
	"go-wallet-api/features/shared/utils/enums"
	"go-wallet-api/features/wallet/business_logic/handlers"
	"go-wallet-api/features/wallet/business_logic/validations"
	"go-wallet-api/middlewares"
)

func RegisterWalletRoutes(api fiber.Router) {
	walletsRoutes := api.Group("/wallets").Use(configs.AddAuthentication()) // set authentication
	//Wallets routes
	walletsRoutes.Get("/", middlewares.RequiresAuthorization([]string{enums.RoleType_Admin, enums.RoleType_User}), handlers.GetAllWalletsHandler)
	walletsRoutes.Get("/:id", middlewares.RequiresAuthorization([]string{enums.RoleType_Admin, enums.RoleType_User}), handlers.GetWalletByIdHandler)
	walletsRoutes.Post("/", middlewares.RequiresAuthorization([]string{enums.RoleType_Admin, enums.RoleType_User}), validations.ValidateWalletRequest, handlers.CreateWalletHandler)
	walletsRoutes.Put("/:id", middlewares.RequiresAuthorization([]string{enums.RoleType_Admin, enums.RoleType_User}), validations.ValidateWalletRequest, handlers.UpdateWalletHandler)
	walletsRoutes.Delete("/:id", middlewares.RequiresAuthorization([]string{enums.RoleType_Admin, enums.RoleType_User}), handlers.DeleteWalletHandler)

}
