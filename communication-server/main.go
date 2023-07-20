package main

import (
	"communication-server/config"
	"communication-server/infrastructure/api"
	"communication-server/infrastructure/cache"
	"communication-server/infrastructure/postgresql"
	"communication-server/internal/usecase"
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	runtime.GOMAXPROCS(1)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	config := config.Load(".")

	log.Println("Connecting to PostgreSQL...")
	pdb := postgresql.New(config.PostgreSQL)
	defer pdb.Close()
	log.Println("Connected to PostgreSQL...")

	router := gin.Default()
	apiGroup := router.Group("/api")

	apiServer := api.New(
		config.App,
		cache.New(nil),
		usecase.NewPlayerUsecase(pdb),
	)
	apiServer.Use(apiGroup)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.App.Port),
		Handler: router,
	}

	go func() {
		log.Printf("Server is listening on port %d", config.App.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panic(err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println(err.Error())
	}

	log.Println("Exited")
}
