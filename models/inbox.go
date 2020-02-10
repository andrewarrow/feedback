package models

import "github.com/jmoiron/sqlx"

var SELECT_INBOX = "SELECT id, sent_to as sentto, sent_from as sentfrom, body, subject, UNIX_TIMESTAMP(created_at) as createdat from inboxes order by created_at desc limit 1000"

type Inbox struct {
	Id        int    `json:"id"`
	SentTo    string `json:"sent_to"`
	SentFrom  string `json:"sent_from"`
	Body      string `json:"body"`
	Subject   string `json:"subject"`
	CreatedAt int64  `json:"created_at"`
}

func SelectInboxes(db *sqlx.DB) ([]Inbox, string) {
	items := []Inbox{}
	err := db.Select(&items, SELECT_INBOX)
	s := ""
	if err != nil {
		s = err.Error()
	}

	return items, s
}
