package config

const (
	DefaultDbName = "postgres"
	Localhost     = "127.0.0.1"
	HttpPort      = 8000
	GrpcPort      = 9000
)

var (
	Defaults = map[string]interface{}{
		"debug":                           false,
		"dev_mode":                        false,
		"http_server.port":                HttpPort,
		"http_server.allow_origins":       "http://localhost:3000",
		"http_server.allow_headers":       "Origin, Content-Type, Accept",
		"http_server.tz":                  "europe/paris",
		"http_server.enable_print_routes": false,
		"http_server.read_header_timeout": "3s",
		"grpc_server.port":                GrpcPort,
		"grpc_server.enable_reflection":   true,
		"grpc_server.tls.enabled":         false,
		"grpc_server.tls.cert":            "",
		"grpc_server.tls.key":             "",
		"auth.token_symmetric_key":        "12345678901234567890123456789012",
		"auth.access_token_duration":      "30s",
		"auth.refresh_token_duration":     "15m",
		"database.auto_migration":         true,
		"database.postgres.enabled":       true,
		"database.postgres.host":          Localhost,
		"database.postgres.port":          5432,
		"database.postgres.username":      "postgres",
		"database.postgres.password":      "postgres",
		"database.postgres.db_name":       DefaultDbName,
		"database.postgres.ssl_mode":      false,
		"database.postgres.verbose":       false,
		"database.postgres.cert_pem":      "",
		"database.postgres.cert_key":      "",
		"database.postgres.max_conns":     10,
		"database.postgres.min_conns":     4,
	}

	EnvKeys = map[string]string{
		"debug":                           "DEBUG",
		"dev_mode":                        "DEV_MODE",
		"init_admin_password":             "INIT_ADMIN_PASSWORD",
		"http_server.port":                "HTTP_SERVER_PORT",
		"http_server.allow_origins":       "HTTP_SERVER_ALLOW_ORIGINS",
		"http_server.allow_headers":       "HTTP_SERVER_ALLOW_HEADERS",
		"http_server.tz":                  "HTTP_SERVER_TZ",
		"http_server.enable_print_routes": "HTTP_SERVER_ENABLE_PRINT_ROUTES",
		"http_server.read_header_timeout": "HTTP_SERVER_READ_HEADER_TIMEOUT",
		"grpc_server.port":                "GRPC_SERVER_PORT",
		"grpc_server.enable_reflection":   "GRPC_SERVER_ENABLE_REFLECTION",
		"grpc_server.tls.enabled":         "GRPC_SERVER_TLS_ENABLED",
		"grpc_server.tls.cert":            "GRPC_SERVER_TLS_CERT",
		"grpc_server.tls.key":             "GRPC_SERVER_TLS_KEY",
		"auth.token_symmetric_key":        "AUTH_TOKEN_SYMMETRIC_KEY",
		"auth.access_token_duration":      "AUTH_ACCESS_TOKEN_DURATION",
		"auth.refresh_token_duration":     "AUTH_REFRESH_TOKEN_DURATION",
		"database.auto_migration":         "DATABASE_AUTO_MIGRATION",
		"database.postgres.enabled":       "DATABASE_POSTGRES_ENABLED",
		"database.postgres.host":          "DATABASE_POSTGRES_HOST",
		"database.postgres.port":          "DATABASE_POSTGRES_PORT",
		"database.postgres.username":      "DATABASE_POSTGRES_USERNAME",
		"database.postgres.password":      "DATABASE_POSTGRES_PASSWORD",
		"database.postgres.database":      "DATABASE_POSTGRES_DATABASE",
		"database.postgres.ssl_mode":      "DATABASE_POSTGRES_SSL_MODE",
		"database.postgres.verbose":       "DATABASE_POSTGRES_VERBOSE",
		"database.postgres.cert_pem":      "DATABASE_POSTGRES_CERT_PEM",
		"database.postgres.cert_key":      "DATABASE_POSTGRES_CERT_KEY",
		"database.postgres.max_conns":     "DATABASE_POSTGRES_MAX_CONNS",
		"database.postgres.min_conns":     "DATABASE_POSTGRES_MIN_CONNS",
	}
)
