package store

import (
	"context"
	"errors"
)

// A Datastore is a storage mechanism for persisting urls mapping.
type Datastore interface {
	ReadURL(ctx context.Context, shortUrl string) (string, error)
	WriteURL(ctx context.Context, shortUrl, url string) error
}

type memoryDatastore struct {
	keyVal map[string]string
}

var _ Datastore = &memoryDatastore{}

// ErrNotFound is thrown when the given short URL was not found in the datastore.
var ErrNotFound = errors.New("short url not found in store")

// NewMemoryDatastore instantiates a in-memory datastore.
func NewMemoryDatastore() (Datastore, error) {
	return &memoryDatastore{keyVal: make(map[string]string)}, nil
}

// ReadURL reads the datastore in an attempt to retreive the URL for the provided shortURL.
func (ds *memoryDatastore) ReadURL(ctx context.Context, shortUrl string) (string, error) {
	val, ok := ds.keyVal[shortUrl]
	if !ok {
		return "", ErrNotFound
	}
	return val, nil
}

// WriteURL writes the datastore to persist the URL keyed with its shorter URL version.
func (ds *memoryDatastore) WriteURL(ctx context.Context, shortUrl, url string) error {
	ds.keyVal[shortUrl] = url
	return nil
}
