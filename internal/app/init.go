package app

import (
	"github.com/kevinnaserwan/lrt-realtime-service/internal/util/env"
)

type app struct {
	config *env.Config
}

//	@contact.name	LRT SUMSEL
//	@contact.email	indraadha24@gmail.com

// @externalDocs.description	OpenAPI
//
// @externalDocs.url			https://swagger.io/resources/open-api/

func NewApp(config *env.Config) *app {
	return &app{
		config: config,
	}
}
