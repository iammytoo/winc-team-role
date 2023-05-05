package user

import (
	"strings"

	"github.com/mirito333/winc-team-role/config"
)

type User struct{
	Id string
	Teams []string
	Roles []string
}

func CreateUser(id, row_teams string) User{
	var user = User{}
	user.Id = id
	user.Teams = strings.Split(row_teams,",")
	user.setRoleIds()
	return user
}


func (u *User) setRoleIds(){
	var roles = []string{}
	roles = append(roles, config.Env.Role.Dict["member"])
	for _, k := range u.Teams{
		roles = append(roles, config.Env.Role.Dict[k])
	}
	u.Roles = roles
}