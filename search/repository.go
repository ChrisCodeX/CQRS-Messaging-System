package search

import (
	"context"

	"github.com/ChrisCodeX/CQRS-Messaging-System/models"
)

/*Abstract implementation of search*/

// Search Repository interface
type SearchRepository interface {
	Close()
	IndexFeed(ctx context.Context, feed models.Feed) error
	SearchFeed(ctx context.Context, query string) ([]models.Feed, error)
}

var repo SearchRepository

func SetSearchRepository(r SearchRepository) {
	repo = r
}

func Close() {
	repo.Close()
}

// Index feeds
func IndexFeed(ctx context.Context, feed models.Feed) error {
	return repo.IndexFeed(ctx, feed)
}

// Search feeds
func SearchFeed(ctx context.Context, query string) ([]models.Feed, error) {
	return repo.SearchFeed(ctx, query)
}
