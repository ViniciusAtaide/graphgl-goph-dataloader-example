package models

import (
	"github.com/graph-gophers/graphql-go"
)

type User struct {
	UserID   graphql.ID `db:"id"`
	Username string     `db:"name"`
	Emoji    string     `db:"emoji"`
	Notes    []*Note
}
