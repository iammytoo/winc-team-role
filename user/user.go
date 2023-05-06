package user

import (
	"strings"

	"github.com/mirito333/winc-team-role/config"
)

type User struct {
	Name  string
	Id    string
	Teams []string
	Roles []string
}

func CreateUser(name, id, row_teams string) User {
	var user = User{}
	user.Name = name
	user.Id = id
	if row_teams != ""{
		user.Teams = strings.Split(row_teams, ", ")
	}
	user.setRoleIds()
	return user
}

func (u *User) setRoleIds() {
	var roles = []string{}
	roles = append(roles, config.Env.Role.Dict["member"])
	if len(u.Teams) != 0{
		for _, k := range u.Teams {
			roles = append(roles, config.Env.Role.Dict[k])
		}
	}
	u.Roles = roles
}
