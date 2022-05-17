package storage

import (
	"context"
	"time"

	"github.com/cod3rcarl/wwdatabase-go-backend/internal/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/pkg/errors"
)

func (s *Service) GetAllChampions(ctx context.Context) (models.Champions, error) {
	champions := models.Champions{}
	query := `
	SELECT
		"id",
		"title_holder",
		COALESCE("title_holder_number", 0) "title_holder_number",
		COALESCE("title_holder_order_number", 0) "title_holder_order_number",
		"date_won",
		"date_lost",
		"show",
		"wrestler_id",
		COALESCE("current_champion", false) "current_champion"
	FROM "champion"
	ORDER BY "date_lost" DESC
	;`

	if err := pgxscan.Select(ctx, s.Pool, &champions, query); err != nil {
		s.logger.Error(err.Error())

		return models.Champions{}, errors.New(err.Error())
	}

	return champions, nil
}

func (s *Service) GetChampionByOrderNumber(ctx context.Context, tn int32) (models.Champion, error) {
	champion := models.Champion{}
	query := `
	SELECT
		"id",
		"title_holder",
		COALESCE("title_holder_number", 0) "title_holder_number",
		COALESCE("title_holder_order_number", 0) "title_holder_order_number",
		"date_won",
		"date_lost",
		"show",
		"wrestler_id",
		COALESCE("current_champion", false) "current_champion"
	FROM "champion"
	WHERE  "title_holder_order_number" = $1
	;`

	if err := pgxscan.Get(ctx, s.Pool, &champion, query, tn); err != nil {
		s.logger.Error(err.Error())

		return models.Champion{}, errors.New(err.Error())
	}

	return champion, nil
}

func (s *Service) GetChampionByID(ctx context.Context, id string) error {
	champion := models.Champion{}
	query := `
	SELECT
		"id",
		"title_holder",
		COALESCE("title_holder_number", 0) "title_holder_number",
		COALESCE("title_holder_order_number", 0) "title_holder_order_number",
		"date_won",
		"date_lost",
		"show",
		"wrestler_id",
		COALESCE("current_champion", false) "current_champion"
	FROM "champion"
	WHERE  "id" = $1
	;`

	if err := pgxscan.Get(ctx, s.Pool, &champion, query, id); err != nil {
		s.logger.Error(err.Error())

		return errors.New(err.Error())
	}

	return nil
}

func (s *Service) GetCurrentChampion(ctx context.Context, cc bool) (models.Champion, error) {
	champion := models.Champion{}
	query := `
	SELECT
		"id",
		"title_holder",
		COALESCE("title_holder_number", 0) "title_holder_number",
		COALESCE("title_holder_order_number", 0) "title_holder_order_number",
		"date_won",
		"date_lost",
		"show",
		"wrestler_id",
		COALESCE("current_champion", false) "current_champion"
	FROM "champion"
	WHERE  "current_champion" = $1
	;`

	if err := pgxscan.Get(ctx, s.Pool, &champion, query, cc); err != nil {
		s.logger.Error(err.Error())

		return models.Champion{}, errors.New(err.Error())
	}

	return champion, nil
}

func (s *Service) GetChampionByDate(ctx context.Context, date time.Time) (models.Champion, error) {
	champion := models.Champion{}
	query := `
	SELECT
		"id",
		"title_holder",
		COALESCE("title_holder_number", 0) "title_holder_number",
		COALESCE("title_holder_order_number", 0) "title_holder_order_number",
		"date_won",
		"date_lost",
		"show",
		"wrestler_id",
		COALESCE("current_champion", false) "current_champion"
	FROM "champion"
	WHERE  "date_won"  <= $1 AND "date_lost" > $1
	;`

	if err := pgxscan.Get(ctx, s.Pool, &champion, query, date); err != nil {
		s.logger.Error(err.Error())

		return models.Champion{}, errors.New(err.Error())
	}

	return champion, nil
}

func (s *Service) GetChampionListByName(ctx context.Context, name string) (models.Champions, error) {
	champions := models.Champions{}
	query := `
	SELECT
		"id",
		"title_holder",
		COALESCE("title_holder_number", 0) "title_holder_number",
		COALESCE("title_holder_order_number", 0) "title_holder_order_number",
		"date_won",
		"date_lost",
		"show",
		"wrestler_id",
		COALESCE("current_champion", false) "current_champion"
	FROM "champion"
	WHERE  "title_holder" = $1
	;`

	if err := pgxscan.Select(ctx, s.Pool, &champions, query, name); err != nil {
		s.logger.Error(err.Error())

		return models.Champions{}, errors.New(err.Error())
	}

	return champions, nil
}

func (s *Service) GetChampionsByShow(ctx context.Context, show string) (models.Champions, error) {
	champions := models.Champions{}
	query := `
	SELECT
		"id",
		"title_holder",
		COALESCE("title_holder_number", 0) "title_holder_number",
		COALESCE("title_holder_order_number", 0) "title_holder_order_number",
		"date_won",
		"date_lost",
		"show",
		"wrestler_id",
		COALESCE("current_champion", false) "current_champion"
	FROM "champion"
	WHERE  "show" = $1
	;`

	if err := pgxscan.Select(ctx, s.Pool, &champions, query, show); err != nil {
		s.logger.Error(err.Error())

		return models.Champions{}, errors.New(err.Error())
	}

	return champions, nil
}

func (s *Service) GetChampionsByYear(ctx context.Context, year models.YearInput) (models.Champions, error) {
	champions := models.Champions{}
	query := `
	SELECT
		"id",
		"title_holder",
		COALESCE("title_holder_number", 0) "title_holder_number",
		COALESCE("title_holder_order_number", 0) "title_holder_order_number",
		"date_won",
		"date_lost",
		"show",
		"wrestler_id",
		COALESCE("current_champion", false) "current_champion"
	FROM "champion"
	WHERE  "date_won"  >= $1 AND "date_lost" <= $2
	ORDER BY "title_holder_order_number" ASC
	;`

	if err := pgxscan.Select(ctx, s.Pool, &champions, query, year.StartDate, year.EndDate); err != nil {
		s.logger.Error(err.Error())

		return models.Champions{}, errors.New(err.Error())
	}

	return champions, nil
}
