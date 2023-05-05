package config

import "github.com/joho/godotenv"

type Environment struct {
	Bot struct {
		TOKEN     string
		CLIENT_ID string
		GUILD_ID  string
	}

	Role struct {
		Dict map[string]string
	}
}

func GetEnvironment() Environment {
	var Environment = Environment{}
	err := godotenv.Load("tmp/.env")
	if err != nil {
		panic("Error loading .env file")
	}
	
	return Environment
}

var Env Environment = GetEnvironment()
