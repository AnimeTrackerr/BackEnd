package main

import (
	"github.com/AnimeTrackerr/v2/backend/config"
	"github.com/AnimeTrackerr/v2/backend/routes"
)

func setup() {
	// load environment variables
	config.Load(".env")

	//setup routes and handlers
	routes.SetRoutes()
}

func main() {
		setup()
}