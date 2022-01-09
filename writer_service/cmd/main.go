package main

import (
	"flag"
	"github.com/romaxa83/mst-app/pkg/logger"
	"github.com/romaxa83/mst-app/writer_service/config"
	"github.com/romaxa83/mst-app/writer_service/internal/server"
	"log"
)

func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.WithName("WriterService")

	s := server.NewServer(appLogger, cfg)
	appLogger.Fatal(s.Run())
}
