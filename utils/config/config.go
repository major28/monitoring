package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	env string
)

// LoadConfig function loads static configuration with a help of viper.
func LoadConfig(configPath string, userEnv string) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error while reading config file: %s", err.Error())
	}

	env = userEnv
}

func Set(key string, value interface{}) { viper.Set(key, value) }

// GetEnvConfig returns the configuration string based on the current env
func GetEnvConfig(configKey string) string {
	config := fmt.Sprintf("%s.%s", env, configKey)
	val := viper.GetString(config)
	return val
}

// GetEnvConfigInt returns the configuration integer based on the current env
func GetEnvConfigInt(configKey string) int {
	config := fmt.Sprintf("%s.%s", env, configKey)
	return viper.GetInt(config)
}

// GetEnvConfigInt64 returns the configuration integer based on the current env
func GetEnvConfigInt64(configKey string) int64 {
	config := fmt.Sprintf("%s.%s", env, configKey)
	return viper.GetInt64(config)
}

// GetEnvConfigBool returns the configuration boolean value based on the current env
func GetEnvConfigBool(configKey string) bool {
	config := fmt.Sprintf("%s.%s", env, configKey)
	return viper.GetBool(config)
}

// GetEnvConfigSub returns the sub-configuration based on the current env
func GetEnvConfigSub(configKey string) *viper.Viper {
	config := fmt.Sprintf("%s.%s", env, configKey)

	return viper.Sub(config)
}

// GetEnvConfigT unmarshals the configuration into the passed data structure
func GetEnvConfigT(configKey string, data interface{}) error {
	config := fmt.Sprintf("%s.%s", env, configKey)

	return viper.UnmarshalKey(config, data)
}

// GetFromEnvOrConfig tries to upload config from environment first and then from the viper.
func GetFromEnvOrConfig(envKey, configKey string) string {
	value := os.Getenv(envKey)

	if value == "" {
		value = viper.GetString(configKey)
	}

	return value
}
