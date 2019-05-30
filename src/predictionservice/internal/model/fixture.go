package model

import (
	"errors"
	"github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	_ "github.com/golang/protobuf/proto"
	"time"
)

type FixturePrediction struct {
	Id           string    `json:"id,omitempty"`
	PredictionId string    `json:"predictionId,omitempty"`
	HomeTeam     string    `json:"hTeam,omitempty"`
	AwayTeam     string    `json:"aTeam,omitempty"`
	HomeGoals    int       `json:"hGoals,omitempty"`
	AwayGoals    int       `json:"aGoals,omitempty"`
	HomeResult   int       `json:"hResult,omitempty"`
	AwayResult   int       `json:"aResult,omitempty"`
	Played       int       `json:"played,omitempty"`
	DateTime     time.Time `json:"dateTime,omitempty"`
	Matchday     int       `json:"matchday,omitempty"`
}

type TeamForm struct {
	Forms []*TeamMatchResult `json:"forms,omitempty"`
}

type TeamMatchResult struct {
	Result   Result   `json:"result,omitempty"`
	Score    string   `json:"score,omitempty"`
	Opponent string   `json:"opponent,omitempty"`
	Location Location `json:"location,omitempty"`
}

type Result string

const (
	WIN  Result = "WIN"
	DRAW Result = "DRAW"
	LOSS Result = "LOSS"
)

type Location string

const (
	HOME Location = "HOME"
	AWAY Location = "AWAY"
)

func TeamFormFromGrpc(form *model.Forms) (map[string]TeamForm, error) {
	inputForms := form.Teams

	teamForms := make(map[string]TeamForm)

	for k, v := range inputForms {
		var forms []*TeamMatchResult

		for _, form := range v.Forms {
			teamMatchResult, err := teamMatchResultFromGrpc(form)
			if err != nil {
				return nil, err
			}

			forms = append(forms, teamMatchResult)
		}

		teamForms[k] = TeamForm{
			Forms: forms,
		}
	}

	return teamForms, nil
}

func teamMatchResultFromGrpc(teamMatchResult *model.TeamMatchResult) (*TeamMatchResult, error) {
	result, err := resultFromGrpc(teamMatchResult.Result)
	if err != nil {
		return nil, err
	}

	location, err := locationFromGrpc(teamMatchResult.Location)
	if err != nil {
		return nil, err
	}

	return &TeamMatchResult{
		Result:   result,
		Score:    teamMatchResult.Score,
		Opponent: teamMatchResult.Opponent,
		Location: location,
	}, nil
}

func resultFromGrpc(result model.TeamMatchResult_Result) (Result, error) {
	switch result {
	case model.TeamMatchResult_WIN:
		return WIN, nil
	case model.TeamMatchResult_DRAW:
		return DRAW, nil
	case model.TeamMatchResult_LOSS:
		return LOSS, nil
	default:
		return "", errors.New("unsupported result")
	}
}

func locationFromGrpc(location model.TeamMatchResult_Location) (Location, error) {
	switch location {
	case model.TeamMatchResult_HOME:
		return HOME, nil
	case model.TeamMatchResult_AWAY:
		return AWAY, nil
	default:
		return "", errors.New("unsupported location")
	}
}

type PredictorData struct {
	Predictions []FixturePrediction  `json:"predictions,omitempty"`
	Forms       map[string]*TeamForm `json:"forms,omitempty"`
}
