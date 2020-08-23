package main

import (
	EComApp "github.com/codedv8/go-ecom-app"
	EComBase "github.com/codedv8/go-ecom-base"
	EComStructsAPI "github.com/codedv8/go-ecom-structs/API"
	"github.com/gin-gonic/gin"
	"log"
)

func (api *API) SysInit(app *EComApp.Application) {

}

func (api *API) Init(app *EComApp.Application) {
	// Create apiRouter
	apiRouter := app.Router.Group("/api/v1")
	// Create basic auth object
	basicAuth := EComBase.NewBasicAuth(app, "API")
	// Add basic auth to qpiRouter
	apiRouter.Use(basicAuth.Use())
	// Add hook
	app.ListenToHook("API_CALL", func(payload interface{}) (bool, error) {
		switch v := payload.(type) {
		case *EComStructsAPI.Root:
			log.Printf("API_CALL: %+v\n", v)
			v.S = "Bye bye"
		}
		return true, nil
	})
	// Add handlers
	apiRouter.Handle("GET", "/", func(c *gin.Context) {
		payload := &EComStructsAPI.Root{S: "Hej"}
		app.CallHook("API_CALL", payload)
		log.Printf("Returned %+v\n", payload)
		c.JSON(200, payload)
	})
}
