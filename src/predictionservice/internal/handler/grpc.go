package handler

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
	"time"
)

type liveMatchServiceServer struct {
	service  interfaces.Service
	interval time.Duration
}

func NewLiveMatchServiceServer(service interfaces.Service, interval time.Duration) (*liveMatchServiceServer, error) {
	log.Print("Registered liveMatchServiceServer handler")

	return &liveMatchServiceServer{
		service:  service,
		interval: interval,
	}, nil
}

func (l *liveMatchServiceServer) GetUpcomingMatches(req *empty.Empty, stream gen.LiveMatchService_GetUpcomingMatchesServer) error {
	matches, err := l.service.GetUpcomingMatches()
	if err != nil {
		return err
	}

	resp := model.ToUpcomingMatchesResponse(matches)

	if err := stream.Send(resp); err != nil {
		return err
	}

	ticker := time.NewTicker(l.interval)
	for {
		select {
		case <-ticker.C:
			matches, err := l.service.GetUpcomingMatches()
			if err != nil {
				return nil
			}

			resp := model.ToUpcomingMatchesResponse(matches)

			if err := stream.Send(resp); err != nil {
				return nil
			}
		}
	}
}

func (l *liveMatchServiceServer) GetMatchSummary(req *gen.PredictionRequest, stream gen.LiveMatchService_GetMatchSummaryServer) error {
	r := model.PredictionRequest{
		UserId:  req.UserId,
		MatchId: req.MatchId,
	}

	matchSummary, err := l.service.GetMatchSummary(stream.Context(), r)
	if err != nil {
		return err
	}

	resp := model.MatchSummaryToGrpc(matchSummary)

	if err := stream.Send(resp); err != nil {
		return err
	}

	ticker := time.NewTicker(l.interval)
	for {
		select {
		case <-ticker.C:
			match, err := l.service.GetMatchFacts(req.MatchId)
			if err != nil {
				return nil
			}

			resp.Match = common.MatchFactsToGrpc(match)

			if err := stream.Send(resp); err != nil {
				return nil
			}
		}
	}
}
