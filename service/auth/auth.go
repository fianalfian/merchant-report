package auth

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"log"
	"merchant-report/model"
	"merchant-report/schema"
	"time"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/golang-jwt/jwt/v4"
)

type authService struct {
	DB rel.Repository
}

func NewAuthService(db rel.Repository) AuthService {
	return &authService{DB: db}
}

type AuthService interface {
	Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error)
}

func (a *authService) Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error) {
	userRes := schema.User{}
	err := a.DB.Find(ctx, &userRes, where.Eq("user_name", req.Username))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(userRes)

	merchantRes := schema.Merchant{}
	err = a.DB.Find(ctx, &merchantRes, where.Eq("user_id", userRes.ID))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(merchantRes)

	var (
		reqPass  = fmt.Sprintf("%x", md5.Sum([]byte(req.Password)))
		userPass = userRes.Password
	)
	if reqPass != userPass {
		return nil, errors.New("invalid password")
	}

	claims := model.JWTPayload{
		UserID:     userRes.ID,
		MerchantID: merchantRes.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			Issuer:    "ReportingAPI",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(token)

	return &model.LoginResponse{
		Token: token,
	}, nil
}
