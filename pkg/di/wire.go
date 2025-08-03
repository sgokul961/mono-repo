//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	api "github.com/sgokul961/GO-PROJECT/pkg/api"
	"github.com/sgokul961/GO-PROJECT/pkg/config"
	"github.com/sgokul961/GO-PROJECT/pkg/db"
	"github.com/sgokul961/GO-PROJECT/pkg/handler"
	"github.com/sgokul961/GO-PROJECT/pkg/repository"
	"github.com/sgokul961/GO-PROJECT/pkg/usecases"
)

func InitializeApi(cfg config.Config) (*api.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabase,
		repository.NewAdminRepository,
		usecases.NewAdminUseCase,
		handler.NewAdminHandler,
		api.NewServerHTTP, // Use your custom server constructor
	)
	return nil, nil // wire will replace this
}
