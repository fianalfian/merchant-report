package auth_controller

import (
	"merchant-report/model"
	"merchant-report/service/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authController struct {
	AuthService auth.AuthService
}

func NewAuthController(authService *auth.AuthService) authController {
	return authController{AuthService: *authService}
}

func (a *authController) Route(e *echo.Echo) {
	e.POST("/api/login", a.Login)
}

func (a *authController) Login(c echo.Context) error {
	var request model.LoginRequest
	err := c.Bind(&request)
	if err != nil {
		return err
	}

	if err = c.Validate(request); err != nil {
		return err
	}

	response, err := a.AuthService.Login(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusOK, model.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
