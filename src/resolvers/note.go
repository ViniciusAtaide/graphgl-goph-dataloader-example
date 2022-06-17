package resolvers

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/viniciusataide/graphql-go-example/src/models"
)

type NoteResolver struct{ N *models.Note }
type NoteInput struct{ Data string }
type CreateNoteArgs struct {
	UserID graphql.ID
	Note   NoteInput
}

func (r *NoteResolver) NoteID() graphql.ID {
	return r.N.NoteID
}

func (r *NoteResolver) Data() string {
	return r.N.Data
}
