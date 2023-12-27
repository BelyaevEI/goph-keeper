package initialization

import (
	"github.com/BelyaevEI/GophKeeper/client/internal/config"
	"github.com/BelyaevEI/GophKeeper/client/internal/data/datarepository/bankrepository"
	"github.com/BelyaevEI/GophKeeper/client/internal/data/datarepository/binrepository"
	"github.com/BelyaevEI/GophKeeper/client/internal/data/datarepository/passwordrepository"
	"github.com/BelyaevEI/GophKeeper/client/internal/data/datarepository/textrepository"
	"github.com/BelyaevEI/GophKeeper/client/internal/data/dataservice"
	"github.com/BelyaevEI/GophKeeper/client/internal/logger"
	"github.com/BelyaevEI/GophKeeper/client/internal/middlewares"
	"github.com/BelyaevEI/GophKeeper/client/internal/route"
	"github.com/BelyaevEI/GophKeeper/client/internal/storage"
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

	// Connect to database for entity
	db, err := storage.Connect2DB(cfg.DSN)
	if err != nil {
		log.Log.Error("connection to database is failed: ", err)
		return Init{}, err
	}

	// Entity for data
	passwordrepository := passwordrepository.New(db.PassDB)                                                // Init client password repository
	textrepository := textrepository.New(db.TextsDB)                                                       // Init client text repository
	binrepository := binrepository.New(db.Bindb)                                                           // Init client bin repository
	bankrepository := bankrepository.New(db.Bankdb)                                                        // Init client bank repository
	dataservice := dataservice.New(log, passwordrepository, textrepository, binrepository, bankrepository) // Init data service

	// Entity for user
	userrepository := userrepository.New(db.UserDB)     // Init client user repository
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
