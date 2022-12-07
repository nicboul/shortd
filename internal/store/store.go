package store

var KV map[string]string

func StoreInit() {
	KV = make(map[string]string)
}
