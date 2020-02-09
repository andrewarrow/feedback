package models

import "fmt"
import "github.com/jmoiron/sqlx"
import "encoding/base64"
import "encoding/json"

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Flavor    string `json:"flavor"`
	CreatedAt int64  `json:"created_at"`
}

func (u *User) Encode() string {
	b, _ := json.Marshal(u)
	s := string(b)
	sEnc := base64.StdEncoding.EncodeToString([]byte(s))
	return sEnc
}

func DecodeUser(s string) *User {
	var user User
	decoded, _ := base64.StdEncoding.DecodeString(s)
	err := json.Unmarshal([]byte(decoded), &user)
	if err != nil {
		return nil
	}
	return &user
}

func SelectUsers(db *sqlx.DB) ([]User, string) {
	users := []User{}
	sql := fmt.Sprintf("SELECT id, email, UNIX_TIMESTAMP(created_at) as createdat from users order by created_at desc")
	err := db.Select(&users, sql)
	s := ""
	if err != nil {
		s = err.Error()
	}

	return users, s
}
