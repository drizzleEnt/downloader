package app

import (
	"context"
	"downloader/internal/api"
	"downloader/internal/config"
	"downloader/internal/config/env"
	"os"

	"go.uber.org/zap"
)

type serviceProvider struct {
	httpConfig config.HTTPConfig
	controller api.Controller

	logger *zap.Logger
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHttpConfig()
		if err != nil {
			s.Logger(context.Background()).Error("failed create new http config", zap.Error(err))
			os.Exit(1)
		}
		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) Logger(_ context.Context) *zap.Logger {
	if s.logger == nil {
		l := zap.New(getCore(getAtomicLevel()))
		s.logger = l
	}

	return s.logger
}

func (s *serviceProvider) Controller() api.Controller {
	if s.controller == nil {
		s.controller = api.New()
	}

	return s.controller
}
