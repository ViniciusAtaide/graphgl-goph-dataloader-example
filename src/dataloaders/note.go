package dataloaders

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/jmoiron/sqlx"
	"github.com/viniciusataide/graphql-go-example/src/models"
)

func Config(db *sqlx.DB) {

	batchFunc := func(_ context.Context, keys []int) []*dataloader.Result[*Note] {
		var results []*dataloader.Result[*Note]
		var notes []*models.Note

		db.Select(&notes, "SELECT * from Notes WHERE user_id in [$1]", keys)
	}

	cache := &dataloader.NoCache[int, *models.Note]{}
	loader := dataloader.NewBatchedLoader(batchFunc, dataloader.WithCache[int, *models.Note](cache))

	result, err := loader.Load(context.Background(), 5)
}
