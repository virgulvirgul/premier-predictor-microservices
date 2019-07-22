package fixture

import (
	"context"
	"errors"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
)

type fixtureService struct {
	fixtureClient gen.FixtureServiceClient
}

func NewFixtureService(fixtureClient gen.FixtureServiceClient) (interfaces.FixtureService, error) {
	return &fixtureService{
		fixtureClient: fixtureClient,
	}, nil
}

func (f *fixtureService) GetMatches() ([]common.Fixture, error) {
	resp, err := f.fixtureClient.GetMatches(context.Background(), &empty.Empty{})
	if err != nil {
		return nil, errors.New("error getting matches")
	}

	var fixtures []common.Fixture
	for _, m := range resp.Matches {
		log.Println(m.DateTime)
		fixtures = append(fixtures, common.FixtureFromGrpc(m))
	}

	return fixtures, nil
}

func (f *fixtureService) GetTeamForm() (map[string]model.TeamForm, error) {
	resp, err := f.fixtureClient.GetTeamForm(context.Background(), &empty.Empty{})
	if err != nil {
		return nil, errors.New("error getting team forms")
	}

	forms, err := model.TeamFormFromGrpc(resp)
	if err != nil {
		return nil, errors.New("error converting team forms")
	}

	return forms, nil
}
