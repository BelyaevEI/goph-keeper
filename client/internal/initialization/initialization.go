package initialization

import (
	"github.com/BelyaevEI/GophKeeper/client/internal/config"
	"github.com/BelyaevEI/GophKeeper/client/internal/data/datarepository"
	"github.com/BelyaevEI/GophKeeper/client/internal/data/dataservice"
	"github.com/BelyaevEI/GophKeeper/client/internal/logger"
	"github.com/BelyaevEI/GophKeeper/client/internal/middlewares"
	"github.com/BelyaevEI/GophKeeper/client/internal/route"
	"github.com/BelyaevEI/GophKeeper/client/internal/storage/postgresql"
	"github.com/BelyaevEI/GophKeeper/client/internal/user/userrepository"
	"github.com/BelyaevEI/GophKeeper/client/internal/user/userservice"
	"github.com/go-chi/chi"
)

type Init struct {
	Host  string
	Port  string
	Route *chi.Mux
}

func Initialization(log *logger.Logger) (Init, error) {

	// Read config file
	cfg, err := config.LoadConfig("..")
	if err != nil {
		log.Log.Error("read config file is fail: ", err)
		return Init{}, err
	}

	// Connect to postgresql
	postgresql, err := postgresql.NewConnect(cfg.DSN)
	if err != nil {
		log.Log.Error("connect to postgresql is failed: ", err)
		return Init{}, err
	}

	// Entity for data
	datarepository := datarepository.New(postgresql)    // Init client data repository
	dataservice := dataservice.New(log, datarepository) // Init client service

	// Entity for user
	userrepository := userrepository.New(postgresql)    // Init client user repository
	userservice := userservice.New(log, userrepository) // Init client user service

	// Create middleware
	middlewares := middlewares.New(userservice)

	// Create new router
	route := route.New(userservice, dataservice, middlewares)

	return Init{
		Route: route,
		Host:  cfg.Host,
		Port:  cfg.Port,
	}, nil
}
