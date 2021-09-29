package services

import (
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func TokenResponder(w http.ResponseWriter, logReq *models.LoginRequest) (*models.LoginResponse, int, error) {
	RefreshAccess := os.Getenv("RefreshAccess")
	RefreshLifetimeMinutes, err := strconv.Atoi(os.Getenv("RefreshLifetimeMinutes"))
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	AccessSecret := os.Getenv("AccessSecret")
	AccessLifetimeMinutes, err := strconv.Atoi(os.Getenv("AccessLifetimeMinutes"))
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	user, err := repository.Repo.GetUserByEmail(logReq.Email)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return nil, 0, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logReq.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return nil, 0, err
	}
	accessString, err := GenerateToken(user.ID, AccessLifetimeMinutes, AccessSecret)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
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
