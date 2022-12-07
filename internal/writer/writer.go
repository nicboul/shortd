package writer

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Writer struct {
}

type Url struct {
	Str string `json:"url"`
}

func NewWriter() *Writer {
	return &Writer{}
}

var base66 = []byte{'-', '.', '_', '~', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func base66Convert(in []byte) string {
	var str string
	var num uint64

	length := len(in)
	if length > 8 {
		length = 8
	}

	for i := length; i > 0; i-- {
		num ^= uint64(in[i-1]) << ((length - i) * 8)
	}

	for num != 0 {
		mod := num % 66
		str = str + string(base66[mod])
		num = num / 66
	}

	return str
}

func (writer *Writer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	var url Url
	err = json.Unmarshal(body, &url)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	hash := sha1.New()
	hash.Write([]byte(url.Str))
	h := hash.Sum(nil)

	fmt.Printf("%x :: %s\n", h[:4], base66Convert(h[:4]))

}
