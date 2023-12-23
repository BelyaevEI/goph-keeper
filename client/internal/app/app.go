package app

import (
	"net/http"

	"github.com/BelyaevEI/GophKeeper/client/internal/initialization"
	"github.com/BelyaevEI/GophKeeper/client/internal/logger"
	"github.com/go-chi/chi"
)

type app struct {
	host  string
	port  string
	route *chi.Mux
}

func NewApp() (*app, error) {

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

	return &app{
		route: init.Route,
		host:  init.Host,
		port:  init.Port,
	}, nil
}

func (a *app) RunServer() error {

	return http.ListenAndServe(a.host+":"+a.port, a.route)

}
