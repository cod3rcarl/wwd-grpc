package client

import (
	"context"

	pb "github.com/cod3rcarl/wwd-protorepo-wwdatabase/v1"
	"github.com/cod3rcarl/wwdatabase-go-backend/internal/storage"

	"go.uber.org/zap"
)

const (
	ServiceName = "wwdatabase"
)

type Service struct {
	serviceOptions
}

type ServiceInterface interface {
	GetAllChampions(ctx context.Context, req *pb.GetChampionsRequest) (*pb.ChampionsList, error)
	GetChampionsByShow(ctx context.Context, show string) (*pb.ChampionsList, error)
	GetChampionsByYear(ctx context.Context, req *pb.GetChampionsByYearRequest) (*pb.ChampionsList, error)
	GetCurrentChampion(ctx context.Context, req *pb.GetCurrentChampionRequest) (*pb.ChampionResponse, error)
	GetChampionByOrderNumber(ctx context.Context, req *pb.ChampionNumber) (*pb.ChampionResponse, error)
	GetChampionByDate(ctx context.Context, req *pb.GetChampionByDateRequest) (*pb.ChampionResponse, error)
	GetChampionListByName(ctx context.Context, name string) (*pb.ChampionsList, error)
	AddChampion(ctx context.Context, req *pb.NewChampionData) (*pb.CreateChampionPayload, error)
	DeleteChampion(ctx context.Context, id string) (*pb.DeleteChampionResponse, error)
}

type serviceOptions struct {
	loggerOption
	storageOption
}

func NewService(opts ...Option) *Service {
	svc := &Service{}

	return svc.withOptions(opts...)
}

func WithLogger(l *zap.Logger) Option {
	return loggerOption{logger: l}
}

func WithStorage(s storage.ServiceInterface) Option {
	return storageOption{store: s}
}
