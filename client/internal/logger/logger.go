package logger

import "go.uber.org/zap"

type Logger struct {
	Log zap.SugaredLogger
}

func New() (*Logger, error) {
	// create installed registrator zap
	logg, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	defer logg.Sync()

	// create registrator SugaredLogger
	sugar := *logg.Sugar()

	return &Logger{Log: sugar}, nil
}
