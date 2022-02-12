package controller_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"week-6-assingment-nidadinch/controller"
	mock "week-6-assingment-nidadinch/mock"
	"week-6-assingment-nidadinch/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUsernameFromPath(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/:Nida", nil)
	username := controller.GetUsernameFromPath(r)
	assert.Equal(t, r.URL.Path[2:], username)
}

func TestControllerServiceReturnWallets(t *testing.T) {
	t.Run("should return wallets correctly", func(t *testing.T) {
		service := mock.NewMockIWalletsService(gomock.NewController(t))
		serviceReturn := &model.WalletsResponse{}

		service.EXPECT().Wallets().Return(serviceReturn, nil).Times(1)

		controller := controller.NewController(service)

		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		controller.Handle(w, r)

		actual := &model.WalletsResponse{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, serviceReturn, actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
		assert.Equal(t, "application/json", w.Header().Get("content-type"))
	})

	t.Run("should return error if service fails", func(t *testing.T) {
		service := mock.NewMockIWalletsService(gomock.NewController(t))
		serviceErr := errors.New("test err")
		service.EXPECT().Wallets().Return(nil, serviceErr).Times(1)

		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		controller := controller.NewController(service)

		controller.Handle(w, r)

		actual := &model.WalletsResponse{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
		assert.Equal(t, w.Body.String(), serviceErr.Error())
	})
}

func TestControllerServiceGetWallet(t *testing.T) {
	t.Run("should return user wallet correctly", func(t *testing.T) {
		service := mock.NewMockIWalletsService(gomock.NewController(t))
		username := "Nida"
		serviceReturn := &model.WalletsResponse{username: 200}

		service.EXPECT().GetWalletByUsername("Nida").Return(serviceReturn, nil).Times(1)

		controller := controller.NewController(service)
		r := httptest.NewRequest(http.MethodGet, "/:"+username, nil)
		w := httptest.NewRecorder()
		controller.Handle(w, r)

		actual := &model.WalletsResponse{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, serviceReturn, actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
		assert.Equal(t, "application/json", w.Header().Get("content-type"))
	})

	t.Run("should return error if wallet doesn't exist", func(t *testing.T) {
		service := mock.NewMockIWalletsService(gomock.NewController(t))
		username := "Nida"
		serviceErr := errors.New("test err")

		service.EXPECT().GetWalletByUsername("Nida").Return(nil, serviceErr).Times(1)

		controller := controller.NewController(service)
		r := httptest.NewRequest(http.MethodGet, "/:"+username, nil)
		w := httptest.NewRecorder()
		controller.Handle(w, r)

		actual := &model.WalletsResponse{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
		assert.Equal(t, w.Body.String(), serviceErr.Error())
	})
}

func TestControllerServiceCreateWallet(t *testing.T) {
	t.Run("should create new wallet", func(t *testing.T) {
		service := mock.NewMockIWalletsService(gomock.NewController(t))
		username := "Nida"
		serviceReturn := &model.WalletsResponse{username: 200}

		service.EXPECT().CreateWalletByUsername("Nida").Return(serviceReturn, nil).Times(1)

		controller := controller.NewController(service)
		r := httptest.NewRequest(http.MethodPut, "/:"+username, nil)
		w := httptest.NewRecorder()
		controller.Handle(w, r)

		actual := &model.WalletsResponse{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, serviceReturn, actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
		assert.Equal(t, "application/json", w.Header().Get("content-type"))
	})
	t.Run("should return error if service fails", func(t *testing.T) {
		service := mock.NewMockIWalletsService(gomock.NewController(t))
		username := "Nida"
		err := errors.New("test error")

		service.EXPECT().CreateWalletByUsername("Nida").Return(nil, err).Times(1)

		controller := controller.NewController(service)
		r := httptest.NewRequest(http.MethodPut, "/:"+username, nil)
		w := httptest.NewRecorder()
		controller.Handle(w, r)

		actual := &model.WalletsResponse{}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
		assert.Equal(t, w.Body.String(), err.Error())
	})

}

func TestControllerServiceUpdateWallet(t *testing.T) {
	t.Run("should update wallet correctly", func(t *testing.T) {

		service := mock.NewMockIWalletsService(gomock.NewController(t))
		username := "Nida"
		amount := 200
		serviceReturn := &model.WalletsResponse{username: amount}

		service.EXPECT().UpdateWalletByUsername("Nida", amount).Return(serviceReturn, nil).Times(1)

		controller := controller.NewController(service)
		w := httptest.NewRecorder()
		bodyReader := strings.NewReader(`{"balance": 200}`)
		r := httptest.NewRequest(http.MethodPost, "/:"+username, bodyReader)
		r.Header.Set("Content-Type", "application/json")

		controller.Handle(w, r)

		actual := &model.WalletsResponse{username: amount}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, serviceReturn, actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
		assert.Equal(t, "application/json", w.Header().Get("content-type"))

	})
	t.Run("should return error if service fails", func(t *testing.T) {

		service := mock.NewMockIWalletsService(gomock.NewController(t))
		username := "Nida"
		amount := 200
		serviceErr := errors.New("test err")
		service.EXPECT().UpdateWalletByUsername("Nida", amount).Return(nil, serviceErr).Times(1)

		controller := controller.NewController(service)
		w := httptest.NewRecorder()
		bodyReader := strings.NewReader(`{"balance": 200}`)
		r := httptest.NewRequest(http.MethodPost, "/:"+username, bodyReader)
		r.Header.Set("Content-Type", "application/json")

		controller.Handle(w, r)

		actual := &model.WalletsResponse{username: amount}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
		assert.Equal(t, w.Body.String(), serviceErr.Error())

	})
	t.Run("should return error when request body is wrong", func(t *testing.T) {
		service := mock.NewMockIWalletsService(gomock.NewController(t))
		username := "Nida"
		amount := 200

		controller := controller.NewController(service)
		w := httptest.NewRecorder()
		bodyReader := strings.NewReader(`{"balance": "200"}`)
		r := httptest.NewRequest(http.MethodPost, "/:"+username, bodyReader)
		r.Header.Set("Content-Type", "application/json")

		controller.Handle(w, r)

		actual := &model.WalletsResponse{username: amount}
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
	})
}

func TestHandlerRouterErrors(t *testing.T) {
	t.Run("should return error if request method is PUT and path doesn't contain any username", func(t *testing.T) {
		mockService := mock.NewMockIWalletsService(gomock.NewController(t))
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPut, "/", nil)
		controller := controller.NewController(mockService)
		controller.Handle(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
		assert.Equal(t, w.Body.String(), "please provide a username to create a wallet")
	})
	t.Run("should return error if request method is POST and path doesn't contain any username", func(t *testing.T) {
		mockService := mock.NewMockIWalletsService(gomock.NewController(t))
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", nil)
		controller := controller.NewController(mockService)
		controller.Handle(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
		assert.Equal(t, w.Body.String(), "please provide a username to update the balance of wallet")
	})
}
