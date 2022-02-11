package main

import (
	"fmt"
	"net/http"
	"week-6-assingment-nidadinch/config"
	"week-6-assingment-nidadinch/controller"
	"week-6-assingment-nidadinch/data"
	"week-6-assingment-nidadinch/service"
)

func main() {
	config := config.Get()
	data := data.NewData()
	service := service.NewService(data)
	c := controller.NewController(service)

	http.HandleFunc("/", c.Handle)

	err := http.ListenAndServe(config.ServerAddr, nil)
	if err != nil {
		fmt.Println(err)
	}
}
