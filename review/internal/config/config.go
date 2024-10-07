package config

import (
	"coupon_service/internal/api"
	"log"
	"runtime"

	"github.com/brumhard/alligotor"
)

type Config struct {
	API           api.Config
	RequiredCores int `env:"REQUIRED_CORES"`
}

func New() Config {
	cfg := Config{}
	if err := alligotor.Get(&cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}

func CheckSystemRequirements(cfg Config) {
	requiredCores := cfg.RequiredCores
	if requiredCores == 0 {
		log.Println("REQUIRED_CORES is not set, defaulting to 32 cores")
		requiredCores = 32
	}

	if requiredCores != runtime.NumCPU() {
		log.Printf("this API is meant to be run on %d core machines\n", requiredCores)
	}
}
