//go:build wireinject
// +build wireinject

package di

import (
	usecase "github.com/athunlal/bookNow-Account-Services/pkg/UseCase"
	"github.com/athunlal/bookNow-Account-Services/pkg/api"
	"github.com/athunlal/bookNow-Account-Services/pkg/api/handler"
	"github.com/athunlal/bookNow-Account-Services/pkg/config"
	"github.com/athunlal/bookNow-Account-Services/pkg/db"
	"github.com/athunlal/bookNow-Account-Services/pkg/repository"
	"github.com/google/wire"
)

func InitApi(cfg config.Config) (*api.ServerHttp, error) {
	wire.Build(
		db.ConnectDataBase,
		repository.NewUserRepo,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
		api.NewServerHttp)
	return &api.ServerHttp{}, nil
}
