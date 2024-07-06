package server

import "github.com/rs/cors"

type Config struct {
	CORS cors.Options
}

func NewConfig() *Config {
	return &Config{
		CORS: cors.Options{
			AllowedOrigins: []string{"http://localhost:3000", "http://0.0.0.0:3000"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"Content-Type", "Authorization"},
		},
	}
}
