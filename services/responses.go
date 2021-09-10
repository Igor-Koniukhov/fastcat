package services

import (
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/model"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const (
	RefreshSecret         = "refresh_secret"
	AccessLifetimeMinutes = 5
	RefreshLifetimeMinutes = 500
	AccessSecret          = "access_secret"
)

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken  string `json:"refresh_token"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func TokenResponder(w http.ResponseWriter,logReq *model.LoginRequest) (*LoginResponse, error) {

	user, err := repository.Repo.GetUserByEmail(logReq.Email)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return nil, err
	}
	fmt.Println(user)

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logReq.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return nil, err
	}

	accessString, err := GenerateToken(user.ID, AccessLifetimeMinutes, AccessSecret)
	refreshString, err := GenerateToken(user.ID, RefreshLifetimeMinutes, RefreshSecret)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil, err
	}
	resp := &LoginResponse{
		accessString,
		refreshString,
	}
	return resp, err
}
