package config

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"github.com/webhook-issue-manager/model"
)

func Config(file string) *model.Config {
	var config model.Config
	vi := viper.New()
	vi.SetConfigFile(file)
	vi.ReadInConfig()
	config.Port = vi.GetInt("port")
	config.Hostname = vi.GetString("hostname")
	config.User = vi.GetString("postgres_user")
	config.Password = vi.GetInt("postgres_password")
	config.Database = vi.GetString("postgres_database")

	return &config
}

func MinioConnection() (*minio.Client, error) {
	endpoint := "http://127.0.0.1:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return minioClient, nil
}
