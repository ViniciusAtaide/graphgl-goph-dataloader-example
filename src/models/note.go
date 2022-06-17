package models

import "github.com/graph-gophers/graphql-go"

type Note struct {
	NoteID graphql.ID `db:"id"`
	Data   string     `db:"data"`
	UserID graphql.ID `db:"user_id"`
}
