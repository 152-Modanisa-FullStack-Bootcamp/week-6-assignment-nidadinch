package data

import (
	"fmt"
)

type Wallet struct {
	Username string
	Balance  int
}

type Data struct {
	wallets []*Wallet
}

type IData interface {
	GetAllWallets() ([]*Wallet, error)
	GetWallet(username string) (*Wallet, error)
	AddWallet(username string, initialBalanceAmount int) (*Wallet, error)
	Update(username string, amount int) (*Wallet, error)
	NewWallet(wallet *Wallet)
}

func (d *Data) NewWallet(wallet *Wallet) {
	d.wallets = append(d.wallets, wallet)
}

func (d *Data) GetAllWallets() ([]*Wallet, error) {
	return d.wallets, nil
}

func (d *Data) GetWallet(username string) (*Wallet, error) {
	for _, u := range d.wallets {
		if u.Username == username {
			return u, nil
		}
	}
	return &Wallet{}, fmt.Errorf("user with ID '%v' not found", username)
}

func (d *Data) AddWallet(username string, initialBalanceAmount int) (*Wallet, error) {
	newWallet := Wallet{Username: username, Balance: initialBalanceAmount}
	d.wallets = append(d.wallets, &newWallet)
	return d.GetWallet(username)
}

func (d *Data) Update(username string, amount int) (*Wallet, error) {
	for _, u := range d.wallets {
		if u.Username == username {
			u.Balance += amount
		}
	}
	return d.GetWallet(username)
}

func NewData() IData {
	// inMemoryDB
	DB := []*Wallet{}
	return &Data{wallets: DB}
}
