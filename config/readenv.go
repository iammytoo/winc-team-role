package config

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	Bot struct {
		TOKEN     string
		CLIENT_ID string
		GUILD_ID  string
	}

	Role struct {
		Dict map[string]string
	}

	Team struct {
		Path string
	}
}

func GetEnvironment() Environment {
	var Environment = Environment{}
	err := godotenv.Load("tmp/.env")
	if err != nil {
		panic("Error loading .env file")
	}
	Environment.Bot.TOKEN = os.Getenv("DISCORD_TOKEN")
	Environment.Bot.CLIENT_ID = os.Getenv("DISCORD_CLIENT")
	Environment.Bot.GUILD_ID = os.Getenv("DISCORD_GUILD")

	Environment.Team.Path = os.Getenv("MEMBER_PATH")

	path := os.Getenv("TEAM_PATH")

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	Environment.Role.Dict = make(map[string]string)
	for _, v := range rows {
		Environment.Role.Dict[v[0]] = v[1]
	}
	return Environment
}

var Env Environment = GetEnvironment()
