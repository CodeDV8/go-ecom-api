package main

import (
	ecomapp "github.com/codedv8/go-ecom-app"
)

// API - The main struct for this module
type API struct {
	App     *ecomapp.Application
	Message string
}

var apiList []API

// Exports

// SysInit - Pre initialization of this module
func SysInit(app *ecomapp.Application) error {
	api := &API{
		App:     app,
		Message: "Welcome to the API plugin",
	}
	api.SysInit(app)

	apiList = append(apiList, *api)

	return nil
}

// Init - Initialization of this module
func Init(app *ecomapp.Application) error {
	for _, api := range apiList {
		api.Init(app)
	}

	return nil
}

// Done - Shut down of this module
func Done(app *ecomapp.Application) error {
	for _, api := range apiList {
		api.Done(app)
	}

	return nil
}

var api API
