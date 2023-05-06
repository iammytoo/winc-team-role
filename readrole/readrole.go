package readrole

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/mirito333/winc-team-role/config"
	"github.com/mirito333/winc-team-role/user"
)

func ReadRole(path string, s *discordgo.Session) []user.User {
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
	var users = []user.User{}
	for _, v := range rows {
		members, err := s.GuildMembersSearch(config.Env.Bot.GUILD_ID, strings.Split(v[0], "#")[0], 1)
		if err != nil || len(members)==0{
			fmt.Println(v[0],"がない")
			continue
		}
		id := members[0].User.ID
		var user = user.CreateUser(v[0], id, v[1])
		users = append(users, user)
	}
	return users
}
