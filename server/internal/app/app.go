package app

import (
	"net/http"

	"github.com/BelyaevEI/GophKeeper/server/internal/initialization"
	"github.com/BelyaevEI/GophKeeper/server/internal/logger"
)

func NewApp() (*http.Server, error) {

	// Create connect to logger
	log, err := logger.New()
	if err != nil {
		return nil, err
	}

	// Initialization additional entities
	init, err := initialization.Initialization(log)
	if err != nil {
		return nil, err
	}

	server := &http.Server{
		Addr:    init.Host + ":" + init.Port,
		Handler: init.Route,
	}

	return server, nil
}
