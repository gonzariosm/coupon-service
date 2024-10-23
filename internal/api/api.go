package api

import (
	"context"
	"coupon_service/internal/entity"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Service interface {
	ApplyCoupon(entity.Basket, string) (*entity.Basket, error)
	CreateCoupon(int, string, int) (*entity.Coupon, error)
	GetCoupons([]string) ([]entity.Coupon, error)
}

type Config struct {
	Host string
	Port int
}

type API struct {
	srv *http.Server
	MUX *gin.Engine
	svc Service
	CFG Config
}

func New[T Service](cfg Config, svc T) *API {
	r := gin.New()
	r.Use(gin.Recovery())

	api := &API{
		MUX: r,
		CFG: cfg,
		svc: svc,
	}

	api.withRoutes() // Routes

	api.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", api.CFG.Port),
		Handler: api.MUX,
	}

	return api
}

func (a *API) withRoutes() {
	apiGroup := a.MUX.Group("/api")
	apiGroup.POST("/apply", a.Apply)
	apiGroup.POST("/create", a.Create)
	apiGroup.GET("/coupons", a.Get)
}

func (a *API) Start() {
	fmt.Printf("Starting server on port %d\n", a.CFG.Port)
	if err := a.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", a.srv.Addr, err)
	}
}

func (a *API) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v\n", err)
	} else {
		fmt.Println("Server shut down gracefully")
	}
}
