package utils

import (
	"merchant-report/model"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetJWTPayload(c echo.Context) *model.JWTPayload {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JWTPayload)
	return claims
}
