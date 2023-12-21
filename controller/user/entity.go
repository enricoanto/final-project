package user

import (
	"time"

	model "github.com/enricoanto/final-project/repository"
)

type (
	RegisterRequest struct {
		FullName string `json:"full_name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
		Balance  int    `json:"balance" validate:"min=0,max=100000000"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	BalanceRequest struct {
		Balance int `json:"balance" validate:"min=0,max=100000000"`
	}

	RegisterResponse struct {
		ID        int       `json:"id"`
		FullName  string    `json:"full_name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		Balance   int       `json:"balance"`
		CreatedAt time.Time `json:"created_at"`
	}

	TokenResponse struct {
		Token string `json:"token"`
	}

	MessageResponse struct {
		Message string `json:"message"`
	}
)

func transformToRegisterResponse(data model.User) RegisterResponse {
	return RegisterResponse{
		ID:        data.ID,
		FullName:  data.FullName,
		Email:     data.Email,
		Password:  data.Password,
		Balance:   data.Balance,
		CreatedAt: data.CreatedAt,
	}
}
