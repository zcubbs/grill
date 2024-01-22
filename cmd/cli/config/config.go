package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"os/user"
	"path/filepath"
	"time"
)

const (
	ViperConfigName = ".grill"
	ViperConfigType = "yaml"
)

type HostIsRequiredError struct{}

func (e *HostIsRequiredError) Error() string {
	return fmt.Sprintf(`"host" is required.
please run "grill config set --host <host>"
or set it manually in the config file: %s`, viper.ConfigFileUsed())
}

func Load() *Config {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal("failed to load config", "err", err.Error())
	}

	return cfg
}

func loadConfig() (*Config, error) {
	cfgFileDir, err := getConfigDir()
	if err != nil {
		return nil, err
	}
	cfgFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}

	initViperPresets(cfgFileDir)

	err = initConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to init config err=%s", err)
	}

	err = viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("unable to load config file path=%s err=%s", cfgFile, err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("could not decode config into struct err=%s", err)
	}

	return &cfg, nil
}

func initViperPresets(cfgFileDir string) {
	viper.AddConfigPath(cfgFileDir)
	viper.SetConfigType(ViperConfigType)
	viper.SetConfigName(ViperConfigName)
}

func initConfig() error {
	// checks is config file exists
	// if not, creates it
	err := viper.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			cfgFile, err := getConfigFile()
			if err != nil {
				return err
			}

			err = viper.WriteConfigAs(cfgFile)
			if err != nil {
				return err
			}
		}
	}

	// checks if config file is empty
	// if so, sets default values
	initConfigViperValues()

	err = viper.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}

func initConfigViperValues() {
	if viper.GetString(debugPropertyName) == "" {
		viper.Set(debugPropertyName, false)
	}

	hostP := fmt.Sprintf("%s.%s", grpcClientPropertyName, grpcClientHostPropertyName)
	if viper.GetString(hostP) == "" {
		viper.Set(hostP, "")
	}

	sessionIdP := fmt.Sprintf("%s.%s", grpcClientPropertyName, grpcClientSessionIdPropertyName)
	if viper.GetString(sessionIdP) == "" {
		viper.Set(sessionIdP, "")
	}

	accessTokenP := fmt.Sprintf("%s.%s", grpcClientPropertyName, grpcClientAccessTokenPropertyName)
	if viper.GetString(accessTokenP) == "" {
		viper.Set(accessTokenP, "")
	}

	refreshTokenP := fmt.Sprintf("%s.%s", grpcClientPropertyName, grpcClientRefreshTokenPropertyName)
	if viper.GetString(refreshTokenP) == "" {
		viper.Set(refreshTokenP, "")
	}

	accessTokenExpiresAtP := fmt.Sprintf("%s.%s", grpcClientPropertyName, grpcClientAccessTokenExpiresAtPropertyName)
	if viper.GetString(accessTokenExpiresAtP) == "" || !isValidDate(viper.GetString(accessTokenExpiresAtP)) {
		viper.Set(accessTokenExpiresAtP, time.Time{})
	}

	refreshTokenExpiresAtP := fmt.Sprintf("%s.%s", grpcClientPropertyName, grpcClientRefreshTokenExpiresAtPropertyName)
	if viper.GetString(refreshTokenExpiresAtP) == "" || !isValidDate(viper.GetString(refreshTokenExpiresAtP)) {
		viper.Set(refreshTokenExpiresAtP, time.Time{})
	}
}

func isValidDate(date string) bool {
	// List of potential date-time formats.
	formats := []string{
		time.RFC3339,
		"2006-01-02 15:04:05.999999 -0700 MST",
		"2006-01-02 15:04:05",
		// ...
	}

	// parse the date string using each format.
	for _, format := range formats {
		if _, err := time.Parse(format, date); err == nil {
			return true // Successfully parsed with this format.
		}
	}

	return false
}

func SaveConfig(cfg *Config) error {
	return saveConfig(cfg, false)
}

func ResetConfig(cfg *Config) error {
	return saveConfig(cfg, true)
}

func saveConfig(cfg *Config, reset bool) error {
	cfgFile, err := getConfigFile()
	if err != nil {
		return fmt.Errorf("failed to get config file err=%w", err)
	}

	initViperPresets(filepath.Dir(cfgFile))

	err = viperSetValues(cfg, reset)
	if err != nil {
		return fmt.Errorf("failed to set values err=%w", err)
	}

	err = viper.WriteConfigAs(cfgFile)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFile() (string, error) {
	cfgFileDir, err := getConfigDir()
	if err != nil {
		return "", fmt.Errorf("unable to get config directory err=%w", err)
	}
	cfgFile := filepath.Join(cfgFileDir, ViperConfigName)
	return cfgFile, nil
}

func getConfigDir() (string, error) {
	usrHomeDir, err := getUserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to get user home directory err=%w", err)
	}

	return usrHomeDir, nil
}

func getUserHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}

func viperSetValues(cfg *Config, reset bool) error {
	viper.Set("debug", cfg.Debug)

	grpcCfg := cfg.GrpcClient
	if reset || grpcCfg.Host != "" {
		viper.Set("grpc_client.host", grpcCfg.Host)
	}

	if reset || grpcCfg.SessionId != "" {
		viper.Set("grpc_client.session_id", grpcCfg.SessionId)
	}

	if reset || grpcCfg.AccessToken != "" {
		viper.Set("grpc_client.access_token", grpcCfg.AccessToken)
	}

	if reset || grpcCfg.RefreshToken != "" {
		viper.Set("grpc_client.refresh_token", grpcCfg.RefreshToken)
	}

	if reset || !grpcCfg.AccessTokenExpiresAt.IsZero() {
		viper.Set("grpc_client.access_token_expires_at", grpcCfg.AccessTokenExpiresAt)
	}

	if reset || !grpcCfg.RefreshTokenExpiresAt.IsZero() {
		viper.Set("grpc_client.refresh_token_expires_at", grpcCfg.RefreshTokenExpiresAt)
	}

	return nil
}

func PrintConfig(cfg *Config) {
	jsonConfig, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("configpath: %s\n", viper.ConfigFileUsed())
	fmt.Printf("%v\n", string(jsonConfig))
}
