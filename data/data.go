package data

import (
	"fmt"
)

type Wallet struct {
	Username string
	Balance  int
}

// inMemoryDB
var wallets []*Wallet

type Data struct{}

type IData interface {
	GetAllWallets() ([]*Wallet, error)
	GetWallet(username string) (*Wallet, error)
	AddWallet(username string, initialBalanceAmount int) (*Wallet, error)
	Update(username string, amount int) (*Wallet, error)
}

func (d *Data) GetAllWallets() ([]*Wallet, error) {
	nida := Wallet{}
	nida.Username = "Nida"
	nida.Balance = 2001
	wallets = append(wallets, &nida)
	return wallets, nil
}

func (d *Data) GetWallet(username string) (*Wallet, error) {
	for _, u := range wallets {
		if u.Username == username {
			return u, nil
		}
	}
	return &Wallet{}, fmt.Errorf("user with ID '%v' not found", username)
}

func (d *Data) AddWallet(username string, initialBalanceAmount int) (*Wallet, error) {
	newWallet := Wallet{Username: username, Balance: initialBalanceAmount}
	wallets = append(wallets, &newWallet)
	return d.GetWallet(username)
}

func (d *Data) Update(username string, amount int) (*Wallet, error) {
	for _, u := range wallets {
		if u.Username == username {
			u.Balance += amount
		}
	}

	return d.GetWallet(username)
}

func NewData() IData {
	return &Data{}
}
