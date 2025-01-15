package cmd

import (
	"easy-cooking/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "omni",
	Short: "Omni Channel Manager",
	Run: func(cmd *cobra.Command, args []string) {
		// Lấy cấu hình database từ Viper

		// Khoi tao config
		config.LoadConfig()

		//dbConn := connectDatabase()
		//s3Client, err := NewS3Client()
		//if err != nil {
		//	log.Fatal(err)
		//}
		//// Khoi tao handler
		//handler := createHandler(dbConn, s3Client)
		//go pushProductsToSalesChannel(dbConn, s3Client)
		//go deleteProductInSalesChannel(dbConn, s3Client)
		//go pushMediumPosts(dbConn, s3Client)
		//// Khoi tao router
		//routers.InitRouter(handler)
		//
		//// Lấy port từ Viper
		//port := config.Config.ServerPort
		//routers.StartRouter(port)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//func connectDatabase() *database.Database {
//	dbConn, err := database.NewDatabase(config.Config.DatabaseDSN)
//	if err != nil {
//		log.Fatalf("Failed to connect to database: %v", err)
//	}
//	return dbConn
//}
//
//func createHandler(dbConn *database.Database, s3Client *dto.S3Client) *handler.Handler {
//	return handler.NewHandler(dbConn.GetDB(), config.Config.JWTConfig, 5*time.Second, s3Client)
//}
