package main

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"go_api_tempalte/db"
	"go_api_tempalte/route"

	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"
)

const (
	maxprocs    = 1
	defaultPort = "1323"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))
	viper.SetDefault("Port", defaultPort)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	runtime.GOMAXPROCS(maxprocs)
}

func main() {
	ctx := context.Background()
	db := db.NewMongoDB()
	e := route.InitRouter(db)

	// Start server
	s := &http.Server{
		Addr:         ":1323",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := e.StartServer(s); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
