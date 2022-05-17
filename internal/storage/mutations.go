package storage

import (
	"context"
	"fmt"

	"github.com/cod3rcarl/wwdatabase-go-backend/internal/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (s *Service) AddChampion(ctx context.Context, input models.CreateChampionInput) (models.Champion, error) {
	var createdChamp models.Champion
	id := uuid.New()
	fmt.Println(input)
	query := `
	INSERT INTO "champion" (
	"id",
	"title_holder",
	"title_holder_number",
	"title_holder_order_number",
	"date_won",
	"show",
	"wrestler_id",
	"current_champion"
)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id,title_holder, COALESCE(title_holder_number, 0)
	"title_holder_number", date_won, date_lost, show, current_champion,
	COALESCE(title_holder_order_number, 0)"title_holder_order_number"`

	if err := pgxscan.Get(
		ctx,
		s.Pool,
		&createdChamp,
		query,
		id.String(),
		input.TitleHolder,
		input.TitleHolderNumber,
		input.TitleHolderOrderNumber,
		input.DateWon,
		input.Show,
		input.WrestlerID,
		true,
	); err != nil {
		s.logger.Error(err.Error())

		return models.Champion{}, errors.New(err.Error())
	}

	return createdChamp, nil
}

func (s *Service) UpdateChampion(ctx context.Context,
	input models.UpdateChampionInput,
) (models.Champion, error) {
	updatedChampion := models.Champion{}
	query := `
	UPDATE champion
	SET date_lost = $2,
	current_champion = $3
	WHERE id = $1
	RETURNING id,title_holder, COALESCE(title_holder_number, 0)
	"title_holder_number", date_won, date_lost, show, wrestler_id,current_champion,
	COALESCE(title_holder_order_number, 0)"title_holder_order_number"
;`

	if err := pgxscan.Get(
		ctx,
		s.Pool,
		&updatedChampion,
		query,
		input.ID,
		input.DateLost,
		input.CurrentChampion,
	); err != nil {
		s.logger.Error(err.Error())

		return models.Champion{}, errors.New(err.Error())
	}

	return updatedChampion, nil
}

func (s *Service) DeleteChampion(ctx context.Context, id string) (string, error) {
	query := `
			DELETE FROM "champion"
			WHERE id = $1
			RETURNING id
		`

	_, err := s.Pool.Exec(ctx, query, id)
	if err != nil {
		s.logger.Error(err.Error())

		return "", errors.New(err.Error())
	}

	return id, nil
}
