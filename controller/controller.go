package controller

import (
	"encoding/json"
	"net/http"
	"regexp"
	"week-6-assingment-nidadinch/model"
	"week-6-assingment-nidadinch/service"
)

type IController interface {
	Wallets(w http.ResponseWriter, r *http.Request)
	GetWallet(w http.ResponseWriter, r *http.Request)
	Handle(w http.ResponseWriter, r *http.Request)
	CreateWallet(w http.ResponseWriter, r *http.Request)
	UpdateWallet(w http.ResponseWriter, r *http.Request)
}

type Controller struct {
	service service.IWalletsService
}

var param = regexp.MustCompile(`/:`)

// handle function for routing
func (c *Controller) Handle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch {
	case http.MethodGet == r.Method:
		switch {
		case param.MatchString(path):
			c.GetWallet(w, r)
		default:
			c.Wallets(w, r)
		}
	case http.MethodPut == r.Method:
		switch {
		case param.MatchString(path):
			c.CreateWallet(w, r)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("please provide a username to create a wallet"))
		}
	case http.MethodPost == r.Method:
		switch {
		case param.MatchString(path):
			c.UpdateWallet(w, r)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("please provide a username to update the balance of wallet"))
		}
	}
}

func GetUsernameFromPath(r *http.Request) string {
	return r.URL.Path[2:]
}

func (c *Controller) Wallets(w http.ResponseWriter, r *http.Request) {
	response, err := c.service.Wallets()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(json)
}

func (c *Controller) GetWallet(w http.ResponseWriter, r *http.Request) {
	username := GetUsernameFromPath(r)
	response, err := c.service.GetWalletByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(json)
}

func (c *Controller) CreateWallet(w http.ResponseWriter, r *http.Request) {
	username := GetUsernameFromPath(r)

	response, err := c.service.CreateWalletByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(json)
}

func (c *Controller) UpdateWallet(w http.ResponseWriter, r *http.Request) {
	username := GetUsernameFromPath(r)

	// parse json body
	decoder := json.NewDecoder(r.Body)
	req := &model.WalletUpdateRequest{}
	err := decoder.Decode(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	amount := req.Balance

	response, err := c.service.UpdateWalletByUsername(username, amount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(json)
}

func NewController(service service.IWalletsService) IController {
	return &Controller{service: service}
}
