package reader

import (
	"fmt"
	"net/http"

	"github.com/nicboul/shortd/internal/store"
)

type Reader struct {
	ds store.Datastore
}

func NewReader(ds store.Datastore) *Reader {
	return &Reader{ds: ds}
}

func (reader *Reader) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// strip the / out of the path
	short := req.URL.Path[1:]
	fmt.Printf("reader: %v\n", short)

	url, err := reader.ds.ReadURL(req.Context(), short)
	if err != nil {
		fmt.Printf("not found: %s\n", url)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Printf("url: %s\n", url)

	http.Redirect(w, req, url, http.StatusMovedPermanently)
}
