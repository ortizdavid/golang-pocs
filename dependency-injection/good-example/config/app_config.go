package config

import "os"

type AppConfig struct {
    DbURL    string
    RedisURL string
    LogPath  string
}

// LoadAppConfig simula o carregamento de variáveis de ambiente (ENVs)
func LoadAppConfig() *AppConfig {
    return &AppConfig{
        // Usamos valores default ou pegamos do sistema
        DbURL:    getEnv("DB_URL", "postgres://user:pass@localhost:5432/db"),
        RedisURL: getEnv("REDIS_URL", "localhost:6379"),
        LogPath:  getEnv("LOG_PATH", "./logs/app.log"),
    }
}

// Função auxiliar simples para o POC
func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}