package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-wallet-api/config"
	"go-wallet-api/features/shared/utils/enums"
	"go-wallet-api/features/shared/utils/validations"
	"go-wallet-api/features/wallet/business_logic/handlers"
	"go-wallet-api/middlewares"
)

func RegisterWalletRoutes(api fiber.Router) {
	walletsRoutes := api.Group("/wallets").Use(config.AddAuthentication()) // set authentication
	//Wallets routes
	walletsRoutes.Get("/", middlewares.RequiresAuthorization([]string{enums.RoleType_Admin, enums.RoleType_User}), handlers.GetAllWalletsHandler)
	walletsRoutes.Get("/:id", middlewares.RequiresAuthorization([]string{enums.RoleType_Admin, enums.RoleType_User}), handlers.GetWalletByIdHandler)
	walletsRoutes.Post("/", middlewares.RequiresAuthorization([]string{enums.RoleType_Admin, enums.RoleType_User}), validations.ValidateWalletRequest, handlers.CreateWalletHandler)
	walletsRoutes.Put("/:id", middlewares.RequiresAuthorization([]string{enums.RoleType_Admin, enums.RoleType_User}), validations.ValidateWalletRequest, handlers.UpdateWalletHandler)
	walletsRoutes.Delete("/:id", middlewares.RequiresAuthorization([]string{enums.RoleType_Admin, enums.RoleType_User}), handlers.DeleteWalletHandler)

}
