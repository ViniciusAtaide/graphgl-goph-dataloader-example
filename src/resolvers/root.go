package resolvers

import (
	"errors"
	"log"

	"github.com/graph-gophers/graphql-go"
	"github.com/jmoiron/sqlx"
	"github.com/viniciusataide/graphql-go-example/src/models"
)

type RootResolver struct {
	Db sqlx.DB
}

func (r *RootResolver) Users() ([]*UserResolver, error) {
	userRxs := []*UserResolver{}
	users := []*models.User{}

	err := r.Db.Select(&users, `SELECT * FROM Users`)

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		userRxs = append(userRxs, &UserResolver{user})
	}
	return userRxs, nil
}

func (r *RootResolver) User(args struct{ UserID graphql.ID }) (*UserResolver, error) {
	user := models.User{}

	err := r.Db.Select(&user, "SELECT * FROM Users where id = $1", args.UserID)

	if err != nil {
		return nil, err
	}

	return &UserResolver{&user}, nil
}

func (r *RootResolver) Notes(args struct{ UserID graphql.ID }) ([]*NoteResolver, error) {
	var noteRxs []*NoteResolver
	var notes []*models.Note

	err := r.Db.Select(&notes, "SELECT * FROM Notes WHERE user_id=$1", args.UserID)

	if err != nil {
		return nil, err
	}

	log.Println(notes)
	for _, note := range notes {
		noteRxs = append(noteRxs, &NoteResolver{note})
	}
	return noteRxs, nil
}

func (r *RootResolver) Note(args struct{ NoteID graphql.ID }) (*NoteResolver, error) {
	var note *models.Note

	err := r.Db.Select(&note, "SELECT * FROM Notes WHERE id=$1", args.NoteID)

	if err != nil {
		return nil, err
	}

	return &NoteResolver{note}, nil
}

func (r *RootResolver) CreateNote(args CreateNoteArgs) (*NoteResolver, error) {
	var note *models.Note
	var user *models.User

	err := r.Db.Select(&user, "SELECT * FROM Users WHERE id=$1", args.UserID)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("User not Found")
	}

	note = &models.Note{NoteID: "n-010", Data: args.Note.Data}
	user.Notes = append(user.Notes, note)

	return &NoteResolver{note}, nil
}
