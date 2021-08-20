package handlers

/*
import (
	"github.com/igor-koniukhov/fastcat/internal/config"
	"github.com/igor-koniukhov/fastcat/internal/repository"
	"net/http"
)
var Repo *UserDBRepository

type UserDBRepository struct {
	App *config.AppConfig
}

func NewHandlers(r *UserDBRepository)  {
	Repo = r

}

func NewUserDBRepostitory(a *config.AppConfig) *UserDBRepository {
	return &UserDBRepository{App: a}
}
var email = "Jone@gmail.com"

func (u *UserDBRepository) GetUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
 u :=repository.MethodRepo.Get()
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
*/