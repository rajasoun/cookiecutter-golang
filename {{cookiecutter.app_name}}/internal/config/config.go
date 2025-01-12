package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

// Config holds the application configuration values
type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		DSN string
	}
	JSONLogs bool
	LogLevel string
}

// Provider defines a set of read-only methods for accessing the application
// configuration params as defined in one of the config files.
type Provider interface {
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

var defaultConfig *viper.Viper

// LoadConfig initializes the configuration
func LoadConfig(appName string) *Config {
	defaultConfig = readViperConfig(appName)

	var config Config
	if err := defaultConfig.Unmarshal(&config); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}
	return &config
}

func readViperConfig(appName string) *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(appName)
	v.AutomaticEnv()

	// Set global defaults
	v.SetDefault("json_logs", false)
	v.SetDefault("loglevel", "debug")

	// Read the config file (if available)
	v.AddConfigPath("config")
	v.SetConfigName("app")
	v.SetConfigType("env")
	if err := v.ReadInConfig(); err != nil {
		log.Printf("No config file found or failed to read config: %v", err)
	}

	return v
}
