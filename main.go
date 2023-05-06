package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mirito333/winc-team-role/config"
	"github.com/mirito333/winc-team-role/readrole"
	"github.com/mirito333/winc-team-role/writerole"
)

func main() {
	s, err := discordgo.New("Bot " + config.Env.Bot.TOKEN)
	users := readrole.ReadRole(config.Env.Team.Path, s)
	if err != nil {
		fmt.Println("DiscordのBotと接続できんかったyo >< :", err)
		return
	}
	defer s.Close()
	writerole.WriteRole(users, s)
	fmt.Println("できたわよ")
}
