package store_test

import (
	"context"
	"testing"

	"github.com/nicboul/shortd/internal/store"
)

func TestMemoryDatastoreWriteURLAndReadURL(t *testing.T) {
	tests := []struct {
		name string

		// Read expectations
		rKey string
		rVal string
		rErr error

		// Write expectations
		wKey string
		wVal string
		wErr error
	}{
		{
			name: "no such key",
			rKey: "foo",
			rErr: store.ErrNotFound,
		},
		{
			name: "key foo returns value bar",
			rKey: "foo",
			rVal: "bar",
			wKey: "foo",
			wVal: "bar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds, _ := store.NewMemoryDatastore()

			if got := ds.WriteURL(context.TODO(), tt.wKey, tt.wVal); got != tt.wErr {
				t.Fatalf("unexpected write error:\n- want:%+v\n- got:%+v", tt.wErr, got)
			}

			got, err := ds.ReadURL(context.TODO(), tt.rKey)
			if err != tt.rErr {
				t.Fatalf("unexpected read error:\n- want:%+v\n- got:%+v", tt.rErr, err)
			}
			if got != tt.rVal {
				t.Fatalf("unexpected read value:\n- want:%+v\n- got:%+v", tt.rVal, got)
			}
		})
	}
}
