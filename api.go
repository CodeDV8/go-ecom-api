package main

import (
	EComApp "github.com/codedv8/go-ecom-app"
)

type API struct {
}

var apiList []API

// Exports
func Init(app *EComApp.Application) error {
	api := &API{}
	api.Init(app)

	apiList = append(apiList, *api)

	return nil
}

func Done(app *EComApp.Application) error {
	for _, api := range apiList {
		api.Done(app)
	}
	return nil
}

var api API
