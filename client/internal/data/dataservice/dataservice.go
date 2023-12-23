package dataservice

import (
	"github.com/BelyaevEI/GophKeeper/client/internal/data/datarepository"
	"github.com/BelyaevEI/GophKeeper/client/internal/logger"
)

type DataService struct {
	log            *logger.Logger
	datarepository *datarepository.DataRepository
}

func New(log *logger.Logger, datarepository *datarepository.DataRepository) *DataService {
	return &DataService{
		datarepository: datarepository,
		log:            log,
	}
}
