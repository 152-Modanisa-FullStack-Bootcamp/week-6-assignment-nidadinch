package service_test

import (
	"fmt"
	"testing"
	"week-6-assingment-nidadinch/config"
	"week-6-assingment-nidadinch/data"
	mock "week-6-assingment-nidadinch/mock"
	"week-6-assingment-nidadinch/model"
	"week-6-assingment-nidadinch/service"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var conf = config.Get()

func TestWalletsService_Wallets(t *testing.T) {

	mockController := gomock.NewController(t)
	mockData := mock.NewMockIData(mockController)

	mockWallets := []*data.Wallet{{Username: "Nida", Balance: 666}}
	mockData.EXPECT().GetAllWallets().Return(mockWallets, nil).Times(1)

	walletService := service.NewService(mockData)
	wallets, err := walletService.Wallets()

	assert.Equal(t, &model.WalletsResponse{"Nida": 666}, wallets)
	assert.Nil(t, err)
}

func TestWalletsService_GetWallet(t *testing.T) {

	mockController := gomock.NewController(t)
	mockData := mock.NewMockIData(mockController)

	username := "Nida"
	mockWallet := &data.Wallet{Username: username, Balance: 100}
	mockData.EXPECT().GetWallet(username).Return(mockWallet, nil).Times(1)

	walletService := service.NewService(mockData)
	wallet, err := walletService.GetWalletByUsername(username)

	assert.Equal(t, &model.WalletsResponse{"Nida": 100}, wallet)
	assert.Nil(t, err)
}

func TestWalletsService_AddWallet(t *testing.T) {

	mockController := gomock.NewController(t)
	mockData := mock.NewMockIData(mockController)

	username := "Nida"
	mockWallet := &data.Wallet{Username: username, Balance: conf.InitialBalanceAmount}
	mockData.EXPECT().AddWallet(username, conf.InitialBalanceAmount).Return(mockWallet, nil).Times(1)

	walletService := service.NewService(mockData)
	wallet, err := walletService.CreateWalletByUsername(username)

	assert.Equal(t, &model.WalletsResponse{"Nida": conf.InitialBalanceAmount}, wallet)
	assert.Nil(t, err)
}

func TestWalletsService_UpdateWallet(t *testing.T) {
	t.Run("should update if amount is not invalid", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockData := mock.NewMockIData(mockController)

		username := "Nida"
		balance := 0
		amount := 100
		mockWalletInitial := &data.Wallet{Username: username, Balance: balance}
		mockWalletForUpdate := &data.Wallet{Username: username, Balance: amount}

		mockData.EXPECT().GetWallet(username).Return(mockWalletInitial, nil).Times(1)
		mockData.EXPECT().Update(username, amount).Return(mockWalletForUpdate, nil).Times(1)

		walletService := service.NewService(mockData)
		wallet, err := walletService.UpdateWalletByUsername(username, amount)

		assert.Equal(t, &model.WalletsResponse{"Nida": amount}, wallet)
		assert.Nil(t, err)
	})
	t.Run("should return error if amount is invalid", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockData := mock.NewMockIData(mockController)

		username := "Nida"
		balance := 0
		amount := -150
		mockWalletInitial := &data.Wallet{Username: username, Balance: balance}

		mockData.EXPECT().GetWallet(username).Return(mockWalletInitial, nil).Times(1)

		walletService := service.NewService(mockData)
		wallet, err := walletService.UpdateWalletByUsername(username, amount)

		assert.Nil(t, nil, wallet)
		assert.Error(t, err)
		assert.Equal(t, err, fmt.Errorf("wallet balance can not be less than '%v'", conf.MinumumBalanceAmount))
	})
}
