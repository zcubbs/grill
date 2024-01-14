package main

import (
	"flag"
	"github.com/zcubbs/grill/gen/openapi"
	"github.com/zcubbs/grill/internal/api"
	"github.com/zcubbs/grill/internal/config"
	"github.com/zcubbs/grill/internal/utils"
	"github.com/zcubbs/log"
	"github.com/zcubbs/log/structuredlogger"
	"os"
)

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

var cfg *config.Configuration

var configPath = flag.String("config", "", "Path to the configuration file")

func init() {
	flag.Parse()

	log.SetLoggerType(structuredlogger.CharmLoggerType)
	log.SetLevel(structuredlogger.InfoLevel)

	// Load configuration
	log.Info("loading configuration...")
	var err error
	cfg, err = config.Load(*configPath)
	if err != nil {
		log.Fatal("failed to load configuration", "error", err)
	}

	cfg.Version = Version
	cfg.Commit = Commit
	cfg.Date = Date

	if cfg.Debug {
		log.SetLevel(structuredlogger.DebugLevel)
		config.PrintConfiguration(*cfg)
	}

	if cfg.DevMode {
		log.SetFormat(structuredlogger.TextFormat)
	}

	// Set the timezone
	err = os.Setenv("TZ", cfg.HttpServer.TZ)
	if err != nil {
		log.Error("failed to set timezone", "error", err)
	}
	utils.CheckTimeZone()

	log.Info("loaded configuration")
}

func main() {
	// Create gRPC Server
	gs, err := api.NewServer(cfg, api.EmbedAssetsOpts{
		Dir:    openapi.OpenApiFs,
		Path:   "/swagger/",
		Prefix: ".",
	})
	if err != nil {
		log.Fatal("failed to create grpc server", "error", err)
	}

	// Start gRPC Server
	go gs.StartGrpcServer()

	// Start HTTP Gateway
	gs.StartHttpGateway()
}
