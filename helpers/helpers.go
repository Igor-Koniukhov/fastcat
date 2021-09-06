package helpers

import (
	"github.com/igor-koniukhov/fastcat/internal/config"
)

var app *config.AppConfig

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

