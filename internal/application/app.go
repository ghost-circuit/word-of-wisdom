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

func NewApp() *fx.App {
	return fx.New(
		fxslog.WithLogger(),
		fx.Provide(func() context.Context {
			return context.Background()
		}),
		fx.Provide(config.LoadConfig),
		fx.Provide(NewLogger),
		fx.Provide(NewDatabaseClient),
		fx.Provide(NewRepository),
		fx.Provide(NewPoWManager),
		fx.Provide(NewQuoteService),
		fx.Provide(NewGRPCServer),
		fx.Provide(NewGRPCHandlers),
		fx.Invoke(invoke),
	)
}

func NewLogger(cfg *config.Config) *slog.Logger {
	log := logger.InitLogger(cfg.IsSugarLogger)
	slog.Debug("Configured logger", slog.Any("config", cfg))
	return log
}

func NewDatabaseClient(lc fx.Lifecycle, ctx context.Context, cfg *config.Config) (*postgres.DatabaseClient, error) {
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

func NewRepository(client *postgres.DatabaseClient) *repositoryQuote.Repository {
	return repositoryQuote.NewRepository(client.DB())
}

func NewPoWManager(cfg *config.Config) *powalgorithm.ProofOfWorkManager {
	return powalgorithm.NewProofOfWorkManager(cfg.Difficulty)
}

func NewQuoteService(repo *repositoryQuote.Repository, powManager *powalgorithm.ProofOfWorkManager) *serviceWordOfWisdom.WordOfWisdomService {
	return serviceWordOfWisdom.NewWordOfWisdomService(repo, powManager)
}

func NewGRPCHandlers(service *serviceWordOfWisdom.WordOfWisdomService) *grpchandlers.WordOfWisdomHandlers {
	return grpchandlers.NewWordOfWisdomHandlers(service)
}

func NewGRPCServer(lc fx.Lifecycle, cfg *config.Config, handlers *grpchandlers.WordOfWisdomHandlers) (*grpcserver.Server, error) {
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
