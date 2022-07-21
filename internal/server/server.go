package server

import (
	"tm/internal/conf"
	"tm/internal/database"
	"tm/internal/store"
)

func Start(cfg conf.Config) {

	jwtSetup(cfg)

	store.SetDBConnection(database.NewDBOptions(cfg))

	router := setRouter()

	// Start listening and serving requests
	router.Run(":8080")
}
