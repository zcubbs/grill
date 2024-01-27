package config

type AgentConfig struct {
	Host  string `json:"host"`
	Debug bool   `json:"debug"`
	Token string `json:"token"`
}
