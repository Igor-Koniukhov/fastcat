package services

import (
	"fmt"
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	web "github.com/igor-koniukhov/webLogger/v3"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
)



type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func TokenResponder(w http.ResponseWriter, logReq *models.LoginRequest) (*LoginResponse, error) {
	RefreshLifetimeMinutes, err := strconv.Atoi(os.Getenv("RefreshLifetimeMinutes"))
	web.Log.Error(err, "message: ", err)
	RefreshAccess := os.Getenv("RefreshAccess")
	AccessSecret := os.Getenv("AccessSecret")
	AccessLifetimeMinutes, err := strconv.Atoi(os.Getenv("AccessLifetimeMinutes"))
	web.Log.Error(err, "message:", err)
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
	refreshString, err := GenerateToken(user.ID, RefreshLifetimeMinutes, RefreshAccess )
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
