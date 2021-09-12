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

func TokenResponder(w http.ResponseWriter, logReq *models.LoginRequest) (*models.LoginResponse, int, error) {
	RefreshLifetimeMinutes, err := strconv.Atoi(os.Getenv("RefreshLifetimeMinutes"))
	web.Log.Error(err, "message: ", err)
	RefreshAccess := os.Getenv("RefreshAccess")
	AccessSecret := os.Getenv("AccessSecret")
	AccessLifetimeMinutes, err := strconv.Atoi(os.Getenv("AccessLifetimeMinutes"))
	web.Log.Error(err, "message:", err)
	user, err := repository.Repo.GetUserByEmail(logReq.Email)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return nil, 0, err
	}
	fmt.Println(user)

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logReq.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return nil, 0, err
	}

	accessString, err := GenerateToken(user.ID, AccessLifetimeMinutes, AccessSecret)
	refreshString, err := GenerateToken(user.ID, RefreshLifetimeMinutes, RefreshAccess)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil, 0, err
	}
	resp := &models.LoginResponse{
		AccessToken:  accessString,
		RefreshToken: refreshString,
	}

	return resp, user.ID, nil
}
