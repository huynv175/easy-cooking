package config

import (
	"easy-cooking/internal/models/dto"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var Config dto.AppConfig

// LoadConfig tải cấu hình từ file
func LoadConfig() {
	viper.SetConfigName("env") // Tên file cấu hình, không cần phần mở rộng
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // Thư mục hiện tại

	// Đọc file cấu hình
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Gán giá trị vào cấu hình ứng dụng
	Config = dto.AppConfig{
		DatabaseDSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
			viper.GetString("database.user"),
			viper.GetString("database.password"),
			viper.GetString("database.host"),
			viper.GetString("database.port"),
			viper.GetString("database.name")),
		ServerPort: viper.GetString("server.port"),
		JWTConfig: dto.JWTConfig{
			SecretKey:        viper.GetString("jwt.secretKey"),
			AccessExpiration: viper.GetDuration("jwt.accessExpiration"),
		},
	}

	log.Println("Loaded config:", Config)
}
