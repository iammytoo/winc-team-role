package readrole

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/mirito333/winc-team-role/user"
)

func ReadRole(path string) []user.User {
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
		var user = user.CreateUser(v[0], v[1])
		users = append(users, user)
	}
	return users
}
