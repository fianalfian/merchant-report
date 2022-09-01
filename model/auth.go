package model

type (
	LoginRequest struct {
		Username string `json:"user_name" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		Token string `json:"token"`
	}
)
