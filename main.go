package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mirito333/winc-team-role/config"
	"github.com/mirito333/winc-team-role/readrole"
	"github.com/mirito333/winc-team-role/writerole"
)

func main() {

	users := readrole.ReadRole("tmp/hoge.csv")
	print(users)
	s, err := discordgo.New("Bot " + config.Env.Bot.TOKEN)
	if err != nil {
		fmt.Println("DiscordのBotと接続できんかったyo >< :", err)
		return
	}
	defer s.Close()

	writerole.WriteRole(users, s)
	fmt.Println("できたわよ")

}
