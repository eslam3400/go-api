package app_config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type SEnvironment struct {
	Port   string
	DBHost string
	DBName string
}

var Env SEnvironment

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	Env = SEnvironment{
		Port:   os.Getenv("PORT"),
		DBHost: os.Getenv("DB_HOST"),
		DBName: os.Getenv("DB_NAME"),
	}

	Env.PrintValues()
}

func (env SEnvironment) PrintValues() {
	fmt.Printf(`
		ENV Variables Loaded!
		***************************
		port: %v,
		db host: %v,
		db name: %v,
	`,
		Env.Port,
		Env.DBHost,
		Env.DBName,
	)
}
