package client

import (
	"context"

	pb "github.com/cod3rcarl/wwd-protorepo-wwdatabase/v1"
	"github.com/cod3rcarl/wwdatabase-go-backend/internal/models"
	"github.com/pkg/errors"
)

func returnChampionList(champions models.Champions, err error) (*pb.ChampionsList, error) {
	if err != nil {
		return nil, errors.Wrap(err, "error in query")
	}
	count := len(champions)

	if count == 0 {
		return nil, ErrNoChampionsReturned
	}
	pbChampions := make([]*pb.Champion, len(champions))
	for i := range champions {
		pbChampions[i] = models.ModelToPBChampion(champions[i])
	}

	return &pb.ChampionsList{TotalCount: int32(count), Champions: pbChampions}, nil
}

func (s *Service) GetChampionByOrderNumber(ctx context.Context, tn *pb.ChampionNumber) (*pb.ChampionResponse, error) {
	champion, err := s.store.GetChampionByOrderNumber(ctx, tn.TitleHolderOrderNumber)
	if err != nil {
		if errors.Is(err, ErrNoChampionsReturned) {
			return nil, ErrNoChampionsReturned
		}

		return nil, errors.Wrap(err, "Error getting champions")

	}

	return &pb.ChampionResponse{
		Champion: models.ModelToPBChampion(&champion),
	}, nil
}

func (s *Service) GetChampionByDate(ctx context.Context, d *pb.GetChampionByDateRequest) (*pb.ChampionResponse, error) {
	champion, err := s.store.GetChampionByDate(ctx, models.TimestampToTime(d.Date))
	if err != nil {
		if errors.Is(err, ErrNoChampionsReturned) {
			return nil, ErrNoChampionsReturned
		}

		return nil, errors.Wrap(err, "Error getting champions")

	}

	return &pb.ChampionResponse{
		Champion: models.ModelToPBChampion(&champion),
	}, nil
}

func (s *Service) GetCurrentChampion(ctx context.Context, cc *pb.GetCurrentChampionRequest) (*pb.ChampionResponse, error) {
	champion, err := s.store.GetCurrentChampion(ctx, cc.CurrentChampion)
	if err != nil {
		if errors.Is(err, ErrNoChampionsReturned) {
			return nil, ErrNoChampionsReturned
		}

		return nil, errors.Wrap(err, "Error getting champions")

	}

	return &pb.ChampionResponse{
		Champion: models.ModelToPBChampion(&champion),
	}, nil
}

func (s *Service) GetAllChampions(ctx context.Context, _ *pb.GetChampionsRequest) (*pb.ChampionsList, error) {
	champions, err := s.store.GetAllChampions(ctx)
	return returnChampionList(champions, err)
}

func (s *Service) GetChampionsByShow(ctx context.Context, show string) (*pb.ChampionsList, error) {
	champions, err := s.store.GetChampionsByShow(ctx, show)

	return returnChampionList(champions, err)
}

func (s *Service) GetChampionsByYear(ctx context.Context, req *pb.GetChampionsByYearRequest) (*pb.ChampionsList, error) {
	champions, err := s.store.GetChampionsByYear(ctx, models.YearInput{
		StartDate: models.TimestampToNullTime(req.StartDate),
		EndDate:   models.TimestampToNullTime(req.EndDate),
	})
	if err != nil {
		if errors.Is(err, ErrNoChampionsReturned) {
			return nil, ErrNoChampionsReturned
		}

		return nil, errors.Wrap(err, "Error getting champions")
	}
	num := len(champions)

	if num < 1 {
		champion1, err2 := s.store.GetChampionByDate(ctx, models.TimestampToTime(req.StartDate))
		champion2, err2 := s.store.GetChampionByDate(ctx, models.TimestampToTime(req.EndDate))
		if champion1.DateLost.Unix() == champion2.DateLost.Unix() {
			return returnChampionList(models.Champions{&champion1}, err2)
		}
		return returnChampionList(models.Champions{&champion1, &champion2}, err2)
	}
	var prevChamp models.Champion

	if champions[0].TitleHolderNumber != 1 {
		prevChamp, err = s.store.GetChampionByOrderNumber(ctx, champions[0].TitleHolderOrderNumber-1)
		if err != nil {
			prevChamp = models.Champion{}
		}
	}

	nextChamp, err := s.store.GetChampionByOrderNumber(ctx, champions[num-1].TitleHolderOrderNumber+1)
	if err != nil {
		nextChamp = models.Champion{}
	}

	champions = append(models.Champions{&prevChamp}, champions...)
	champions = append(champions, &nextChamp)

	return returnChampionList(champions, err)
}

func (s *Service) GetChampionListByName(ctx context.Context, name string) (*pb.ChampionsList, error) {
	champions, err := s.store.GetChampionListByName(ctx, name)

	return returnChampionList(champions, err)
}
