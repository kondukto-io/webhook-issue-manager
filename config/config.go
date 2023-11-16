package config

import (
	"log"

	"github.com/webhook-issue-manager/model"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

func Config(file string) (*model.Config, error) {
	var config model.Config
	var vi = viper.New()
	vi.SetConfigFile(file)
	if err := vi.ReadInConfig(); err != nil {
		return nil, err
	}

	config.Port = vi.GetInt("port")
	config.Hostname = vi.GetString("hostname")
	config.User = vi.GetString("postgres_user")
	config.Password = vi.GetInt("postgres_password")
	config.Database = vi.GetString("postgres_database")

	return &config, nil
}

func MinioConnection() (*minio.Client, error) {
	var endpoint = "192.168.2.224:9000"
	var accessKeyID = "minioadmin"
	var secretAccessKey = "minioadmin"

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatalln(err)
	}

	return minioClient, nil
}
