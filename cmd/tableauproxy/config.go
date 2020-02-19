package main

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Domain string
	Port   int

	AzureSQLServer   string
	AzureSQLPort     int
	AzureSQLUser     string
	AzureSQLPassword string
	AzureSQLDatabase string
}

func GetConfig() *Config {
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		log.Fatal("Please specify the environment variable DOMAIN")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	portNumber, _ := strconv.Atoi(port)

	azureSQLServer := os.Getenv("AZURE_SQL_SERVER")
	if azureSQLServer == "" {
		log.Fatal("Please specify the environment variable AZURE_SQL_SERVER")
	}

	azureSQLPort := os.Getenv("AZURE_SQL_PORT")
	if azureSQLPort == "" {
		azureSQLPort = "1433"
	}
	azureSQLPortNumber, _ := strconv.Atoi(azureSQLPort)

	azureSQLUser := os.Getenv("AZURE_SQL_USER")
	if azureSQLUser == "" {
		log.Fatal("Please specify the environment variable AZURE_SQL_USER")
	}

	azureSQLPassword := os.Getenv("AZURE_SQL_PASSWORD")
	if azureSQLPassword == "" {
		log.Fatal("Please specify the environment variable AZURE_SQL_PASSWORD")
	}

	azureSQLDatabase := os.Getenv("AZURE_SQL_DATABASE")
	if azureSQLDatabase == "" {
		log.Fatal("Please specify the environment variable AZURE_SQL_DATABASE")
	}

	return &Config{
		Domain: domain,
		Port:   portNumber,

		AzureSQLServer:   azureSQLServer,
		AzureSQLPort:     azureSQLPortNumber,
		AzureSQLUser:     azureSQLUser,
		AzureSQLPassword: azureSQLPassword,
		AzureSQLDatabase: azureSQLDatabase,
	}
}
