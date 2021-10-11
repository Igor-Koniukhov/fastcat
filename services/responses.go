package services

import (
	"github.com/igor-koniukhov/fastcat/internal/models"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	web "github.com/igor-koniukhov/webLogger/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func TokenResponder(w http.ResponseWriter, logReq *models.LoginRequest) (*models.LoginResponse, int, error) {

	user, err := repository.Repo.GetUserByEmail(logReq.Email)
	if err != nil {
		web.Log.Error(err)
		return nil, 0, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logReq.Password)); err != nil {
		web.Log.Error(err)
		return nil, 0, err
	}
	resp, err := TokenGenerator(w, user.ID)
	if err != nil {
		if err != nil {
			log.Println(err)
			return nil, 0, err
		}
	}
	return resp, user.ID, nil
}

func TokenGenerator(w http.ResponseWriter, id int) (*models.LoginResponse, error) {
	RefreshAccess := os.Getenv("RefreshAccess")
	RefreshLifetimeMinutes, err := strconv.Atoi(os.Getenv("RefreshLifetimeMinutes"))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	AccessSecret := os.Getenv("AccessSecret")
	AccessLifetimeMinutes, err := strconv.Atoi(os.Getenv("AccessLifetimeMinutes"))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	accessString, err := GenerateToken(id, AccessLifetimeMinutes, AccessSecret)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	refreshString, err := GenerateToken(id, RefreshLifetimeMinutes, RefreshAccess)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil, err
	}
	resp := &models.LoginResponse{
		AccessToken:  accessString,
		RefreshToken: refreshString,
	}
	return resp, nil
}
