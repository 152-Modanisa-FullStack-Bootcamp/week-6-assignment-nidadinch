package service

import (
	"fmt"
	"week-6-assingment-nidadinch/config"
	"week-6-assingment-nidadinch/data"
	"week-6-assingment-nidadinch/model"
)

type IWalletsService interface {
	Wallets() (*model.WalletsResponse, error)
	GetWalletByUsername(username string) (*model.WalletsResponse, error)
	CreateWalletByUsername(username string) (*model.WalletsResponse, error)
	UpdateWalletByUsername(username string, amount int) (*model.WalletsResponse, error)
}

type WalletsService struct {
	Data data.IData
}

var conf = config.Get()

func (s *WalletsService) Wallets() (*model.WalletsResponse, error) {
	wallets, err := s.Data.GetAllWallets()

	m := model.WalletsResponse{}
	for _, v := range wallets {
		m[v.Username] = v.Balance
	}

	return &m, err
}

func (s *WalletsService) GetWalletByUsername(username string) (*model.WalletsResponse, error) {

	wallet, _ := s.Data.GetWallet(username)

	m := model.WalletsResponse{}
	m[wallet.Username] = wallet.Balance

	return &m, nil
}
func (s *WalletsService) CreateWalletByUsername(username string) (*model.WalletsResponse, error) {
	wallet, _ := s.Data.AddWallet(username, conf.InitialBalanceAmount)

	m := model.WalletsResponse{}
	m[wallet.Username] = wallet.Balance

	return &m, nil
}

func (s *WalletsService) UpdateWalletByUsername(username string, amount int) (*model.WalletsResponse, error) {

	wallet, _ := s.Data.GetWallet(username)

	if conf.MinumumBalanceAmount <= wallet.Balance+amount {
		wallet, _ = s.Data.Update(username, amount)
	} else {
		return nil, fmt.Errorf("wallet balance can not be less than '%v'", conf.MinumumBalanceAmount)
	}

	m := model.WalletsResponse{}
	m[wallet.Username] = wallet.Balance

	return &m, nil
}

func NewService(data data.IData) IWalletsService {
	return &WalletsService{Data: data}
}
