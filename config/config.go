package config

import (
	"os"
	"strconv"
	"strings"
)

// TrefleConfig holds trefle object structure
type TrefleConfig struct {
	Token    string
	UrlIds   string
	UrlPlant string
}

// MongoConfig holds mongo object structure
type MongoConfig struct {
	Username string
	Host     string
	PWD      string
}

// Config holds app config structure
type Config struct {
	Trefle          TrefleConfig
	Mongo           MongoConfig
	DebugMode       bool
	TLSVerification bool
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		Trefle: TrefleConfig{
			Token:    getEnv("TREFLE_TOKEN", ""),
			UrlIds:   getEnv("TREFLE_IDS_URL", ""),
			UrlPlant: getEnv("TREFLE_PLANT_URL", ""),
		},
		Mongo: MongoConfig{
			Username: getEnv("MONGO_USER", ""),
			Host:     getEnv("MONGO_HOST", ""),
			PWD:      getEnv("MONGO_PWD", ""),
		},
		DebugMode:       getEnvAsBool("DEBUG_MODE", true),
		TLSVerification: getEnvAsBool("TLS_VERIFICATION", true),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
