//go:build wireinject
// +build wireinject

package di

import (
	"ecommerce/identity/config"
	grpcDelivery "ecommerce/identity/internal/delivery/grpc"
	"ecommerce/identity/internal/delivery/http"
	"ecommerce/identity/internal/repository"
	"ecommerce/identity/internal/usecase"
	"ecommerce/identity/pkg/grpcserver"
	"ecommerce/identity/pkg/httpserver"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/uchin-mentorship/ecommerce-go/identity"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var useCaseSet = wire.NewSet(config.NewConfig, provideGormDB, provideIdentityUseCase, provideIdentityRepo)

func InitializeHttpServer() (*httpserver.Server, func(), error) {
	panic(wire.Build(
		useCaseSet,
		gin.New,
		provideHttpServer,
		http.NewRouter,
	))
}

func InitializeGRPCServer() (*grpcserver.GRPCServer, func(), error) {
	panic(wire.Build(
		useCaseSet,
		grpc.NewServer,
		provideGRPCServerOptions,
		provideGRPCServer,
		provideGRPCIdentityService,
	))
}

func provideGRPCIdentityService(u usecase.AuthUsecase) identity.IdentityServiceServer {
	return grpcDelivery.NewIdentityService(u)
}

func provideGRPCServerOptions() []grpc.ServerOption {
	return nil
}

func provideGRPCServer(cfg *config.Config, server *grpc.Server, delivery identity.IdentityServiceServer) *grpcserver.GRPCServer {
	identity.RegisterIdentityServiceServer(server, delivery)
	return grpcserver.New(server, cfg.GRPC.Address)
}

func provideAuthRepo(db *gorm.DB) usecase.AuthRepo {
	return repository.New(db)
}

func provideAuthUseCase(r usecase.AuthRepo) usecase.AuthUsecase {
	return usecase.NewIdentity(r)
}

func provideGormDB(cfg *config.Config) (*gorm.DB, func(), error) {
	db, err := gorm.Open(postgres.Open(cfg.PG.URL), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	return db, func() {
		conn, err := db.DB()
		if err != nil {
			log.Printf("failed to get db connection, %v", err)
			return
		}
		conn.Close()
	}, nil
}

func provideHttpServer(router *http.Router, handler *gin.Engine, cfg *config.Config) *httpserver.Server {
	router.Register()
	return httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
}
