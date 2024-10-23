package config

import (
	"coupon_service/internal/api"
	"log"
	"runtime"

	"github.com/brumhard/alligotor"
)

type Config struct {
	API           api.Config
	RequiredCores int `config:"env=cores"`
}

func New() Config {
	cfg := Config{
		// Default configuration
		API: api.Config{
			Port: 8080,
		},
		RequiredCores: 1,
	}
	if err := alligotor.Get(&cfg); err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded Config: %+v\n", cfg)
	return cfg
}

func CheckSystemRequirements(cfg Config) {
	requiredCores := cfg.RequiredCores
	if requiredCores != runtime.NumCPU() {
		log.Printf("this API is meant to be run on %d core machines\n", requiredCores)
	}
}
