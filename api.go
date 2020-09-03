package main

import (
	EComApp "github.com/codedv8/go-ecom-app"
)

// API - The main struct for this module
type API struct {
	App     *EComApp.Application
	Message string
}

var apiList []API

// Exports

// SysInit - Pre initialization of this module
func SysInit(app *EComApp.Application) error {
	api := &API{
		App:     app,
		Message: "Welcome to the API plugin",
	}
	api.SysInit(app)

	apiList = append(apiList, *api)

	return nil
}

// Init - Initialization of this module
func Init(app *EComApp.Application) error {
	for _, api := range apiList {
		api.Init(app)
	}

	return nil
}

// Done - Shut down of this module
func Done(app *EComApp.Application) error {
	for _, api := range apiList {
		api.Done(app)
	}

	return nil
}

var api API
