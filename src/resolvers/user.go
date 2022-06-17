package resolvers

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/viniciusataide/graphql-go-example/src/models"
)

type UserResolver struct{ U *models.User }

func (r *UserResolver) UserID() graphql.ID {
	return r.U.UserID
}

func (r *UserResolver) Username() string {
	return r.U.Username
}

func (r *UserResolver) Emoji() string {
	return r.U.Emoji
}

func (r *UserResolver) Notes() ([]*NoteResolver, error) {
	rootRx := &RootResolver{}

	return rootRx.Notes(struct{ UserID graphql.ID }{UserID: r.U.UserID})
}
