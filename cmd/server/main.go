package main

import (
    "context"
    "flag"
    "github.com/zcubbs/grill/gen/openapi"
    "github.com/zcubbs/grill/internal/api"
    "github.com/zcubbs/grill/internal/config"
    migration "github.com/zcubbs/grill/internal/db/migration"
    db "github.com/zcubbs/grill/internal/db/sqlc"
    dbUtil "github.com/zcubbs/grill/internal/db/util"
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
    // Init context
    ctx := context.Background()

    // Database Migration
    err := migration.Run(cfg.Database)
    if err != nil {
        log.Fatal("failed to perform database migrations", "error", err)
    }

    // Connect to database
    conn, err := dbUtil.Connect(ctx, cfg.Database)
    if err != nil {
        log.Fatal("failed to connect to database", "error", err)
    }

    // Initialize store
    store := db.NewSQLStore(conn)

    // Initialize admin user
    err = dbUtil.InitAdminUser(store, *cfg)
    if err != nil {
        log.Fatal("failed to initialize admin user", "error", err)
    }

    // Create gRPC Server
    gs, err := api.NewServer(nil, cfg, api.EmbedAssetsOpts{
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
