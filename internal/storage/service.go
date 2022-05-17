package storage

import (
	"context"
	"time"

	"github.com/cod3rcarl/wwdatabase-go-backend/base/pgx"
	"github.com/cod3rcarl/wwdatabase-go-backend/internal/models"
	"go.uber.org/zap"
)

type ServiceInterface interface {
	WWDatabaseService
	Stop()
}

type WWDatabaseService interface {
	GetAllChampions(ctx context.Context) (models.Champions, error)
	GetChampionsByShow(ctx context.Context, show string) (models.Champions, error)
	GetChampionsByYear(ctx context.Context, input models.YearInput) (models.Champions, error)
	GetChampionListByName(ctx context.Context, name string) (models.Champions, error)
	GetCurrentChampion(ctx context.Context, cc bool) (models.Champion, error)
	AddChampion(ctx context.Context, input models.CreateChampionInput) (models.Champion, error)
	UpdateChampion(ctx context.Context, input models.UpdateChampionInput) (models.Champion, error)
	GetChampionByOrderNumber(ctx context.Context, tn int32) (models.Champion, error)
	GetChampionByID(ctx context.Context, id string) error
	GetChampionByDate(ctx context.Context, date time.Time) (models.Champion, error)
	DeleteChampion(ctx context.Context, id string) (string, error)
}

type Service struct {
	serviceOptions
}

type serviceOptions struct {
	baseOption
	loggerOption
}

func NewStorage(opts ...Option) *Service {
	svc := &Service{}

	return svc.withOptions(opts...)
}

func WithLogger(l *zap.Logger) Option {
	return loggerOption{logger: l}
}

func WithBase(b *pgx.Service) Option {
	return baseOption{b}
}
