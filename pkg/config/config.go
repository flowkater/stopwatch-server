package config

import (
	"os"
)

// Config 애플리케이션 설정
type Config struct {
	Port string
}

// NewConfig 기본 설정으로 Config 생성
func NewConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // 기본 포트
	}

	return &Config{
		Port: port,
	}
}
