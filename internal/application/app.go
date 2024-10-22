package application

import (
	"context"
	"log/slog"

	"github.com/robbert229/fxslog"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/alisher-baizhumanov/word-of-wisdom/internal/application/config"
	serviceWordOfWisdom "github.com/alisher-baizhumanov/word-of-wisdom/internal/domain/service/word-of-wisdom"
	repositoryQuote "github.com/alisher-baizhumanov/word-of-wisdom/internal/infrastructure/postgres/quote"
	grpchandlers "github.com/alisher-baizhumanov/word-of-wisdom/internal/presentation/grpc-handlers"
	"github.com/alisher-baizhumanov/word-of-wisdom/pkg/adapter/postgres"
	desc "github.com/alisher-baizhumanov/word-of-wisdom/pkg/generated/wisdom"
	grpcserver "github.com/alisher-baizhumanov/word-of-wisdom/pkg/grpc-server"
	"github.com/alisher-baizhumanov/word-of-wisdom/pkg/grpc-server/interceptor"
	powalgorithm "github.com/alisher-baizhumanov/word-of-wisdom/pkg/pow-algorithm"
	"github.com/alisher-baizhumanov/word-of-wisdom/pkg/system/logger"
)

// NewApp creates a new application.
func NewApp() *fx.App {
	return fx.New(
		fxslog.WithLogger(),
		fx.Provide(func() context.Context {
			return context.Background()
		}),
		fx.Provide(config.LoadConfig),
		fx.Provide(newLogger),
		fx.Provide(newDatabaseClient),
		fx.Provide(newRepository),
		fx.Provide(newPoWManager),
		fx.Provide(newQuoteService),
		fx.Provide(newGRPCServer),
		fx.Provide(newGRPCHandlers),
		fx.Invoke(invoke),
	)
}

// newLogger creates a new logger.
func newLogger(cfg *config.Config) *slog.Logger {
	log := logger.InitLogger(cfg.IsSugarLogger)
	slog.Debug("Configured logger", slog.Any("config", cfg))
	return log
}

// newDatabaseClient creates a new database client.
func newDatabaseClient(ctx context.Context, lc fx.Lifecycle, cfg *config.Config) (*postgres.DatabaseClient, error) {
	client, err := postgres.NewClient(ctx, cfg.DatabaseDSN())
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if errPing := client.Ping(ctx); errPing != nil {
				return errPing
			}
			slog.Info("Database connection established")
			return nil
		},
		OnStop: func(_ context.Context) error {
			if errClose := client.Close(); errClose != nil {
				return errClose
			}
			slog.Info("Database connection closed")
			return nil
		},
	})

	return client, nil
}

// newRepository creates a new repository.
func newRepository(client *postgres.DatabaseClient) *repositoryQuote.Repository {
	return repositoryQuote.NewRepository(client.DB())
}

// newPoWManager creates a new PoW manager.
func newPoWManager(cfg *config.Config) *powalgorithm.ProofOfWorkManager {
	return powalgorithm.NewProofOfWorkManager(cfg.Difficulty)
}

// newQuoteService creates a new quote service.
func newQuoteService(repo *repositoryQuote.Repository, powManager *powalgorithm.ProofOfWorkManager) *serviceWordOfWisdom.WordOfWisdomService {
	return serviceWordOfWisdom.NewWordOfWisdomService(repo, powManager)
}

// newGRPCHandlers creates new gRPC handlers.
func newGRPCHandlers(service *serviceWordOfWisdom.WordOfWisdomService) *grpchandlers.WordOfWisdomHandlers {
	return grpchandlers.NewWordOfWisdomHandlers(service)
}

// newGRPCServer creates a new gRPC server.
func newGRPCServer(lc fx.Lifecycle, cfg *config.Config, handlers *grpchandlers.WordOfWisdomHandlers) (*grpcserver.Server, error) {
	srv, err := grpcserver.NewGRPCServer(
		cfg.Addr,
		[]grpcserver.Service{
			{
				ServiceDesc: &desc.WordOfWisdomService_ServiceDesc,
				Handler:     handlers,
			},
		},
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(
			interceptor.Recover,
			interceptor.Logger,
			grpchandlers.ConvertErrorInterceptor,
		),
	)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			srv.Start()
			slog.Info("gRPC server started", slog.String("address", cfg.Addr))
			return nil
		},
		OnStop: func(_ context.Context) error {
			srv.Stop()
			slog.Info("gRPC server stopped")
			return nil
		},
	})

	return srv, nil
}

func invoke(*grpcserver.Server) {
	slog.Info("Application started")
}
