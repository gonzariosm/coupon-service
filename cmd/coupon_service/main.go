package main

import (
	"coupon_service/internal/api"
	"coupon_service/internal/config"
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.New()
	repo := memdb.New()

	// Check system requirements
	config.CheckSystemRequirements(cfg)

	svc := service.New(repo)
	couponAPI := api.New(cfg.API, svc)

	// Start the coupon service
	go func() {
		fmt.Println("Starting Coupon service server")
		couponAPI.Start()
	}()

	// Wait for a terminal signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Block until receive a term signal
	// TODO: IMHO is not the best practice set a "timer" here to shutdown the service
	sig := <-signalChan
	fmt.Printf("Received signal: %v. Shutting down the server...\n", sig)

	couponAPI.Close()
	fmt.Println("Coupon service server shut down gracefully")
}
