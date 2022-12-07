package reader

import (
	"fmt"
	"net/http"

	"github.com/nicboul/shortd/internal/store"
)

type Reader struct {
}

func NewReader() *Reader {
	return &Reader{}
}

func (reader *Reader) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// strip the / out of the path
	short := req.URL.Path[1:]
	fmt.Printf("reader: %v\n", short)

	url := store.KV[short]
	fmt.Printf("url: %s\n", url)

	http.Redirect(w, req, url, http.StatusMovedPermanently)
}
