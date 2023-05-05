package writerole

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/mirito333/winc-team-role/config"
	"github.com/mirito333/winc-team-role/user"
)

func WriteRole(users []user.User, s *discordgo.Session) {
	guild_id := config.Env.Bot.GUILD_ID
	for _, u := range users{
		for _, id := range u.Roles{
			s.GuildMemberRoleAdd(guild_id,u.Id,id)
		}
		fmt.Println(u.Id ,"に", strings.Join(u.Roles, ","), "を追加しました。")
	}
}
