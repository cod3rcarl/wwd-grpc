package server

import (
	"context"

	pb "github.com/cod3rcarl/wwd-protorepo-wwdatabase/v1"
	"github.com/pkg/errors"
)

func (s *Service) GetChampions(ctx context.Context, req *pb.GetChampionsRequest) (
	*pb.ChampionsList, error,
) {
	s.logger.Info("gRPC call: GetAllChampions")
	pbMc, err := s.wwdatabase.GetAllChampions(ctx, req)
	if err != nil {
		return nil, errors.Errorf("error in grpc call", err)
	}

	return pbMc, nil
}

func (s *Service) GetChampionsByShow(ctx context.Context, req *pb.GetChampionsByShowRequest) (
	*pb.ChampionsList, error,
) {
	s.logger.Info("gRPC call: GetChampionsByShow")

	show := req.GetShow()
	if show == "" {
		return nil, errors.New("must provide a show name")
	}

	pbMc, err := s.wwdatabase.GetChampionsByShow(ctx, show)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting champ")
	}

	return pbMc, nil
}

func (s *Service) GetChampionsByYear(ctx context.Context, req *pb.GetChampionsByYearRequest) (
	*pb.ChampionsList, error,
) {
	s.logger.Info("gRPC call: GetChampionsByYear")

	pbMc, err := s.wwdatabase.GetChampionsByYear(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting champ")
	}

	return pbMc, nil
}

func (s *Service) GetCurrentChampion(ctx context.Context, req *pb.GetCurrentChampionRequest) (
	*pb.ChampionResponse, error,
) {
	s.logger.Info("gRPC call: GetCurrentChampions")
	pbMc, err := s.wwdatabase.GetCurrentChampion(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting champ")
	}

	return pbMc, nil
}

func (s *Service) GetChampionByDate(ctx context.Context, req *pb.GetChampionByDateRequest) (
	*pb.ChampionResponse, error,
) {
	s.logger.Info("gRPC call: GetChampionByDate")
	pbMc, err := s.wwdatabase.GetChampionByDate(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting champ")
	}

	return pbMc, nil
}

func (s *Service) GetChampionByOrderNumber(ctx context.Context, req *pb.ChampionNumber) (
	*pb.ChampionResponse, error,
) {
	s.logger.Info("gRPC call: GetPreviousChampions")
	pbMc, err := s.wwdatabase.GetChampionByOrderNumber(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting champ")
	}

	return pbMc, nil
}

func (s *Service) GetChampionByName(ctx context.Context, req *pb.GetChampionByNameRequest) (
	*pb.ChampionsList, error,
) {
	s.logger.Info("gRPC call: GetChampionsByName")

	name := req.GetName()
	if name == "" {
		return nil, errors.New("must provide a name")
	}

	pbMc, err := s.wwdatabase.GetChampionListByName(ctx, name)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting champ")
	}

	return pbMc, nil
}

func (s *Service) AddChampion(ctx context.Context, req *pb.NewChampionData) (
	*pb.CreateChampionPayload, error,
) {
	s.logger.Info("gRPC call: AddChampion")
	pbMc, err := s.wwdatabase.AddChampion(ctx, req)
	if err != nil {
		return nil, errors.New("Cannot add champs")
	}

	return pbMc, nil
}

func (s *Service) DeleteChampion(ctx context.Context, req *pb.DeleteChampionRequest) (
	*pb.DeleteChampionResponse, error,
) {
	s.logger.Info("gRPC call: DeleteChampion")
	i, err := s.wwdatabase.DeleteChampion(ctx, req.Id)
	if err != nil {
		return nil, errors.Errorf("error in grpc call", err)
	}

	return i, nil
}
