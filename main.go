package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AsdGroup8/ASD_QRCodeCheckIn/conf"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/db"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/log"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/service"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var engine *gin.Engine

func main() {
	rootCmd := &cobra.Command{
		Use:     "http server",
		PreRunE: Initialize,
		Run:     Run,
	}
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("root cmd execute error. %v", err)
	}
}

// Initialize ...
func Initialize(cmd *cobra.Command, _ []string) error {
	conf.Init(cmd)
	if err := log.Init(conf.LogFile, conf.LogLevel); err != nil {
		return err
	}
	if err := db.Init(conf.DBConnStr); err != nil {
		return err
	}
	if err := service.Init(); err != nil {
		return err
	}
	initEngine()
	log.Info("app initialize complete")
	return nil
}

// Run ...
func Run(_ *cobra.Command, _ []string) {
	server := http.Server{
		Addr:    conf.Addr,
		Handler: engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("fail to start http server. %v", err)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
	<-sigint
	log.Info("closing server in 3 seconds...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Errorf("server shutdown error. %v", err)
		return
	}
	<-ctx.Done()
	log.Info("server shutdown")
}

func initEngine() {
	engine = gin.New()
	engine.Use(ginzap.Ginzap(log.L(), time.RFC3339, false))
	engine.Use(ginzap.RecoveryWithZap(log.L(), true))
}
