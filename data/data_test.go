package data_test

import (
	"testing"
	"week-6-assingment-nidadinch/config"
	"week-6-assingment-nidadinch/data"

	"github.com/stretchr/testify/assert"
)

var conf = config.Get()

func TestGetAllWallets_Empty(t *testing.T) {

	tests := []struct {
		name    string
		d       *data.Data
		want    []*data.Wallet
		wantErr bool
	}{
		{"testCase01", &data.Data{}, []*data.Wallet{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := data.NewData()
			got, err := d.GetAllWallets()
			assert.Nil(t, err)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestGetAllWallets(t *testing.T) {

	tests := []struct {
		name    string
		d       *data.Data
		want    []*data.Wallet
		wantErr bool
	}{
		{"testCase01", &data.Data{}, []*data.Wallet{{Username: "Nida", Balance: 666}, {Username: "John", Balance: 999}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := data.NewData()
			wallet1 := data.Wallet{Username: "Nida", Balance: 666}
			wallet2 := data.Wallet{Username: "John", Balance: 999}
			d.NewWallet(&wallet1)
			d.NewWallet(&wallet2)

			got, err := d.GetAllWallets()
			assert.Nil(t, err)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestGetWallet(t *testing.T) {

	tests := []struct {
		name     string
		d        *data.Data
		username string
		want     *data.Wallet
		wantErr  bool
	}{
		{"testCase01", &data.Data{}, "Nida", &data.Wallet{Username: "Nida", Balance: 666}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := data.NewData()
			wallet1 := data.Wallet{Username: "Nida", Balance: 666}
			d.NewWallet(&wallet1)
			got, err := d.GetWallet(tt.username)
			assert.Nil(t, err)
			assert.Equal(t, got, tt.want)
		})
	}
}
func TestGetWallet_Empty(t *testing.T) {

	tests := []struct {
		name     string
		d        *data.Data
		username string
		want     *data.Wallet
		wantErr  bool
	}{
		{"testCase01", &data.Data{}, "Nida", &data.Wallet{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := data.NewData()
			got, err := d.GetWallet(tt.username)
			assert.Error(t, err, "user with ID '%v' not found", tt.username)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestData_AddWallet(t *testing.T) {

	tests := []struct {
		name                 string
		d                    *data.Data
		username             string
		initialBalanceAmount int
		want                 *data.Wallet
		wantErr              bool
	}{
		{"testCase01", &data.Data{}, "Nida", conf.InitialBalanceAmount, &data.Wallet{Username: "Nida", Balance: conf.InitialBalanceAmount}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := data.NewData()
			got, err := d.AddWallet(tt.username, tt.initialBalanceAmount)
			assert.Nil(t, err)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestAddWallet_Existing(t *testing.T) {

	tests := []struct {
		name                 string
		d                    *data.Data
		username             string
		initialBalanceAmount int
		want                 *data.Wallet
		wantErr              bool
	}{
		{"testCase01", &data.Data{}, "Nida", conf.InitialBalanceAmount, &data.Wallet{Username: "Nida", Balance: conf.InitialBalanceAmount}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := data.NewData()
			wallet1 := data.Wallet{Username: "Nida", Balance: conf.InitialBalanceAmount}
			d.NewWallet(&wallet1)
			got, err := d.AddWallet(tt.username, tt.initialBalanceAmount)
			assert.Nil(t, err)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestUpdateWallet(t *testing.T) {
	tests := []struct {
		name     string
		d        *data.Data
		username string
		amount   int
		want     *data.Wallet
		wantErr  bool
	}{
		{"testCase01", &data.Data{}, "Nida", 100, &data.Wallet{Username: "Nida", Balance: 100}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := data.NewData()
			wallet1 := data.Wallet{Username: "Nida", Balance: conf.InitialBalanceAmount}
			d.NewWallet(&wallet1)
			got, err := d.Update(tt.username, tt.amount)
			assert.Nil(t, err)
			assert.Equal(t, got, tt.want)
		})
	}
}
