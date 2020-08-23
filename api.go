package main

import (
	EComApp "github.com/codedv8/go-ecom-app"
)

type API struct {
	App     *EComApp.Application
	Message string
}

var apiList []API

// Exports
func SysInit(app *EComApp.Application) error {
	api := &API{
		App:     app,
		Message: "Welcome to the API plugin",
	}
	api.SysInit(app)

	apiList = append(apiList, *api)

	return nil
}

func Init(app *EComApp.Application) error {
	for _, api := range apiList {
		api.Init(app)
	}

	return nil
}

func Done(app *EComApp.Application) error {
	for _, api := range apiList {
		api.Done(app)
	}

	return nil
}

var api API
