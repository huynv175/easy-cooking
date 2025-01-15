package config

import (
	"fmt"
	"omni/internal/models/dto"

	"github.com/spf13/viper"
)

var Config dto.AppConfig

// LoadConfig tải cấu hình từ file
func LoadConfig() {
	viper.AutomaticEnv()
	//Config = dto.AppConfig{
	//	DatabaseDSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
	//		viper.GetString("DATABASE_USER"),
	//		viper.GetString("DATABASE_PASSWORD"),
	//		viper.GetString("DATABASE_HOST"),
	//		viper.GetString("DATABASE_PORT"),
	//		viper.GetString("DATABASE_NAME")),
	//	ServerPort: viper.GetString("SERVER_PORT"),
	//	JWTConfig: dto.JWTConfig{
	//		SecretKey:        viper.GetString("JWT_SECRET_KEY"),
	//		AccessExpiration: viper.GetDuration("JWT_ACCESS_EXPIRATION"),
	//	},
	//	S3Config: dto.S3Config{
	//		AccessKey: viper.GetString("S3_ACCESS_KEY"),
	//		SecretKey: viper.GetString("S3_SECRET_KEY"),
	//		Region:    viper.GetString("S3_REGION"),
	//		Bucket:    viper.GetString("S3_BUCKET"),
	//	},
	//	ChatGPTKey: viper.GetString("CHATGPT_KEY"),
	//}

	//log.Println(Config)
}
