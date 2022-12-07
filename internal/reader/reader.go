package reader

import (
	"fmt"
	"net/http"
)

type Reader struct {
}

func NewReader() *Reader {
	return &Reader{}
}

func (reader *Reader) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("reader: %v\n", req.RequestURI)
}
