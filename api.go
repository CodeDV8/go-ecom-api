package main

import (
	"fmt"
	EComApp "github.com/codedv8/go-ecom-app"
	EComBase "github.com/codedv8/go-ecom-base"
)

type API struct {
	EComBase.Plugin
}

func Init(app *EComApp.Application) error {
	fmt.Printf("Let's rock and foul\n")
	return nil
}

func Done(app *EComApp.Application) error {
	fmt.Printf("Tackar för mig\n")
	return nil
}

var api API
