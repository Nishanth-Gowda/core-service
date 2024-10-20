package config

import (
	"log/slog"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/spf13/viper"
)

var (
	DynamoDBClient *dynamodb.DynamoDB
)

func LoanConfig() {

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// Reading configurations
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		slog.Any("Error reading config file", err)
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(viper.GetString("aws.region")),
		Credentials: credentials.NewStaticCredentials(viper.GetString("AWS_ACCESS_KEY_ID"), viper.GetString("AWS_SECRET_ACCESS_KEY"), ""),
	})
	if err != nil {
		slog.Any("Error creating session", err)
	}

	// Create DynamoDB client
	DynamoDBClient = dynamodb.New(sess)
}