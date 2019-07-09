package handler

import (
	"bytes"
	"encoding/json"
	authmocks "github.com/cshep4/premier-predictor-microservices/src/common/auth/mocks"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
	leaguemocks "github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

const (
	pin        = int64(12345)
	userId     = "🆔"
	leagueName = "🏆🏆🏆🏆🏆🏆"
)

func TestHttpHandler_getLeague(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := leaguemocks.NewMockService(ctrl)
	auth := authmocks.NewMockAuthenticator(ctrl)

	handler, err := NewHttpHandler(service, auth)
	require.NoError(t, err)

	t.Run("it should return internal server error if an error occurred", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/"+userId,
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": userId,
		})

		service.EXPECT().GetUsersLeagueList(userId).Return(nil, errors.New("some error"))

		handler.getUsersLeagueList(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "some error", responseBody.Message)
	})

	t.Run("it should return ok with league overview in body", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/"+userId,
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": userId,
		})

		leagueOverview := &model.StandingsOverview{
			OverallLeagueOverview: model.OverallLeagueOverview{
				UserCount: 50,
			},
		}

		service.EXPECT().GetUsersLeagueList(userId).Return(leagueOverview, nil)

		handler.getUsersLeagueList(rr, req)

		var responseBody model.StandingsOverview
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, leagueOverview, &responseBody)
	})
}

func TestHttpHandler_addLeague(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := leaguemocks.NewMockService(ctrl)
	auth := authmocks.NewMockAuthenticator(ctrl)

	handler, err := NewHttpHandler(service, auth)
	require.NoError(t, err)

	t.Run("it should return bad request if the request body is invalid", func(t *testing.T) {
		b, err := json.Marshal("invalid body")
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPost,
			"/",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		handler.addLeague(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, invalidRequestBody, responseBody.Message)
	})

	t.Run("it should return internal server error if another error occurred", func(t *testing.T) {
		b, err := json.Marshal(addLeagueRequest{
			Id:   userId,
			Name: leagueName,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPost,
			"/",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().AddUserLeague(userId, leagueName).Return(nil, errors.New("some error"))

		handler.addLeague(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "some error", responseBody.Message)
	})

	t.Run("it should return created with league in body", func(t *testing.T) {
		b, err := json.Marshal(addLeagueRequest{
			Id:   userId,
			Name: leagueName,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPost,
			"/",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		league := &model.League{
			Name:  leagueName,
			Users: []string{userId},
			Pin:   pin,
		}

		service.EXPECT().AddUserLeague(userId, leagueName).Return(league, nil)

		handler.addLeague(rr, req)

		var responseBody model.League
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, league, &responseBody)
	})
}

func TestHttpHandler_joinLeague(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := leaguemocks.NewMockService(ctrl)
	auth := authmocks.NewMockAuthenticator(ctrl)

	handler, err := NewHttpHandler(service, auth)
	require.NoError(t, err)

	t.Run("it should return bad request if the request body is invalid", func(t *testing.T) {
		b, err := json.Marshal("invalid body")
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/join",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		handler.joinLeague(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, invalidRequestBody, responseBody.Message)
	})

	t.Run("it should return not found if league not found", func(t *testing.T) {
		b, err := json.Marshal(leagueRequest{
			Id:   userId,
			Pin: pin,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/join",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().JoinUserLeague(userId, pin).Return(nil, model.ErrLeagueNotFound)

		handler.joinLeague(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, model.ErrLeagueNotFound.Error(), responseBody.Message)
	})

	t.Run("it should return internal server error if another error occurred", func(t *testing.T) {
		b, err := json.Marshal(leagueRequest{
			Id:   userId,
			Pin: pin,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/join",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().JoinUserLeague(userId, pin).Return(nil, errors.New("some error"))

		handler.joinLeague(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "some error", responseBody.Message)
	})

	t.Run("it should return ok with league overview in body", func(t *testing.T) {
		b, err := json.Marshal(leagueRequest{
			Id:   userId,
			Pin: pin,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/join",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		leagueOverview := &model.LeagueOverview{
			Pin: pin,
			LeagueName: leagueName,
			Rank: 1,
		}

		service.EXPECT().JoinUserLeague(userId, pin).Return(leagueOverview, nil)

		handler.joinLeague(rr, req)

		var responseBody model.LeagueOverview
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, leagueOverview, &responseBody)
	})
}

func TestHttpHandler_leaveLeague(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := leaguemocks.NewMockService(ctrl)
	auth := authmocks.NewMockAuthenticator(ctrl)

	handler, err := NewHttpHandler(service, auth)
	require.NoError(t, err)

	t.Run("it should return bad request if the request body is invalid", func(t *testing.T) {
		b, err := json.Marshal("invalid body")
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/leave",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		handler.leaveLeague(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, invalidRequestBody, responseBody.Message)
	})

	t.Run("it should return not found if league not found", func(t *testing.T) {
		b, err := json.Marshal(leagueRequest{
			Id:   userId,
			Pin: pin,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/leave",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().LeaveUserLeague(userId, pin).Return(model.ErrLeagueNotFound)

		handler.leaveLeague(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, model.ErrLeagueNotFound.Error(), responseBody.Message)
	})

	t.Run("it should return internal server error if another error occurred", func(t *testing.T) {
		b, err := json.Marshal(leagueRequest{
			Id:   userId,
			Pin: pin,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/leave",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().LeaveUserLeague(userId, pin).Return(errors.New("some error"))

		handler.leaveLeague(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "some error", responseBody.Message)
	})

	t.Run("it should return ok", func(t *testing.T) {
		b, err := json.Marshal(leagueRequest{
			Id:   userId,
			Pin: pin,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/leave",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().LeaveUserLeague(userId, pin).Return(nil)

		handler.leaveLeague(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestHttpHandler_renameLeague(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := leaguemocks.NewMockService(ctrl)
	auth := authmocks.NewMockAuthenticator(ctrl)

	handler, err := NewHttpHandler(service, auth)
	require.NoError(t, err)

	t.Run("it should return bad request if the request body is invalid", func(t *testing.T) {
		b, err := json.Marshal("invalid body")
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/rename",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		handler.renameLeague(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, invalidRequestBody, responseBody.Message)
	})

	t.Run("it should return not found if league not found", func(t *testing.T) {
		b, err := json.Marshal(renameRequest{
			Pin: pin,
			Name: leagueName,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/rename",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().RenameUserLeague(pin, leagueName).Return(model.ErrLeagueNotFound)

		handler.renameLeague(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, model.ErrLeagueNotFound.Error(), responseBody.Message)
	})

	t.Run("it should return internal server error if another error occurred", func(t *testing.T) {
		b, err := json.Marshal(renameRequest{
			Pin: pin,
			Name: leagueName,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/rename",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().RenameUserLeague(pin, leagueName).Return(errors.New("some error"))

		handler.renameLeague(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "some error", responseBody.Message)
	})

	t.Run("it should return ok", func(t *testing.T) {
		b, err := json.Marshal(renameRequest{
			Pin: pin,
			Name: leagueName,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/rename",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().RenameUserLeague(pin, leagueName).Return(nil)

		handler.renameLeague(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestHttpHandler_getLeagueTable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := leaguemocks.NewMockService(ctrl)
	auth := authmocks.NewMockAuthenticator(ctrl)

	handler, err := NewHttpHandler(service, auth)
	require.NoError(t, err)

	t.Run("it should return bad request if pin is invalid", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/standings/invalidPin",
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": invalidPin,
		})

		handler.getLeagueTable(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, invalidPin, responseBody.Message)
	})

	t.Run("it should return not found if league not found", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/standings/" + strconv.FormatInt(pin, 10),
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": strconv.FormatInt(pin, 10),
		})

		service.EXPECT().GetLeagueTable(pin).Return(nil, model.ErrLeagueNotFound)

		handler.getLeagueTable(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, model.ErrLeagueNotFound.Error(), responseBody.Message)
	})

	t.Run("it should return internal server error if another error occurred", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/standings/" + strconv.FormatInt(pin, 10),
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": strconv.FormatInt(pin, 10),
		})

		service.EXPECT().GetLeagueTable(pin).Return(nil, errors.New("some error"))

		handler.getLeagueTable(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "some error", responseBody.Message)
	})

	t.Run("it should return ok", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/standings/" + strconv.FormatInt(pin, 10),
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": strconv.FormatInt(pin, 10),
		})

		leagueTable := []*model.LeagueUser{
			{
				Id: userId,
				Score: 1,
			},
		}

		service.EXPECT().GetLeagueTable(pin).Return(leagueTable, nil)

		handler.getLeagueTable(rr, req)

		var responseBody []*model.LeagueUser
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, leagueTable, responseBody)
	})
}

func TestHttpHandler_getOverallTable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := leaguemocks.NewMockService(ctrl)
	auth := authmocks.NewMockAuthenticator(ctrl)

	handler, err := NewHttpHandler(service, auth)
	require.NoError(t, err)

	t.Run("it should return internal server error if another error occurred", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/standings",
			nil,
		)
		rr := httptest.NewRecorder()

		service.EXPECT().GetOverallLeagueTable().Return(nil, errors.New("some error"))

		handler.getOverallTable(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "some error", responseBody.Message)
	})

	t.Run("it should return ok", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/standings",
			nil,
		)
		rr := httptest.NewRecorder()

		leagueTable := []*model.LeagueUser{
			{
				Id: userId,
				Score: 1,
			},
		}

		service.EXPECT().GetOverallLeagueTable().Return(leagueTable, nil)

		handler.getOverallTable(rr, req)

		var responseBody []*model.LeagueUser
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, leagueTable, responseBody)
	})
}