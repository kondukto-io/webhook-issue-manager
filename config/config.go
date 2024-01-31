package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/webhook-issue-manager/model"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

func Config(configFile string) *model.Config {
	if _, err := os.Stat(configFile); err != nil {
		log.Fatalf("failed to reach config directory: %v", err)
	}

	viper.SetConfigName(filepath.Base(configFile))
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config file: %s", err)
	}

	var config model.Config
	config.Port = viper.GetInt("application.db.postgres.port")
	config.Host = viper.GetString("application.db.postgres.host")
	config.User = viper.GetString("application.db.postgres.user")
	config.Password = viper.GetInt("application.db.postgres.password")
	config.Database = viper.GetString("application.db.postgres.database")

	return &config
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
