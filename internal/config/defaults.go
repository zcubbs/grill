package config

const (
	HttpPort = 8000
	GrpcPort = 9000
)

var (
	defaults = map[string]interface{}{
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
	}

	envKeys = map[string]string{
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
	}
)
