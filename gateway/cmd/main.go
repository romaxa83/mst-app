package main

import (
	"flag"
	"github.com/romaxa83/mst-app/gateway/config"
	"github.com/romaxa83/mst-app/gateway/internal/server"
	"github.com/romaxa83/mst-app/pkg/logger"
	"log"
)

// @contact.name Roman Rodomanov
// @contact.url https://github.com/romaxa83
// @contact.email romaxa83@ukr.net
func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.WithName("UltimateService")

	s := server.NewServer(appLogger, cfg)
	appLogger.Fatal(s.Run())
}
