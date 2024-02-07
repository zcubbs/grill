package openapi

import "embed"

//go:embed * agent/v1/* grill/v1/* user/v1/*
var OpenApiFs embed.FS
