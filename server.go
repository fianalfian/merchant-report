package main

import (
	"merchant-report/config/db"
	auth_controller "merchant-report/controller/auth"
	trx_controller "merchant-report/controller/transaction"
	"merchant-report/model"
	auth_service "merchant-report/service/auth"
	merchant_service "merchant-report/service/merchant"
	outlet_service "merchant-report/service/outlet"
	trx_service "merchant-report/service/transaction"
	"merchant-report/utils"
	"net/http"

	"github.com/go-rel/rel"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	v, trans := utils.NewValidator()
	e := echo.New()
	d := db.Init()
	dbRepo := rel.New(d)
	authService := auth_service.NewAuthService(dbRepo)
	authController := auth_controller.NewAuthController(&authService)

	trxService := trx_service.NewTransactionService(dbRepo)
	merchantService := merchant_service.NewMerchantService(dbRepo)
	outletService := outlet_service.NewOutletService(dbRepo)
	trxController := trx_controller.NewTransactionController(&trxService, &merchantService, &outletService)

	IsLoggedIn := middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &model.JWTPayload{},
		SigningKey: []byte("secret"),
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.Validator = &utils.Validator{Validator: v, Trans: trans}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK v1.0.0")
	})

	trxController.Route(e, IsLoggedIn)
	authController.Route(e)

	e.Logger.Fatal(e.Start(":8000"))
}
