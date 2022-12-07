package main

import (
	"flag"

	"github.com/nicboul/shortd/internal/server"
	"github.com/nicboul/shortd/internal/store"
)

func main() {

	ipFlag := flag.String("ip", "127.0.0.1", "the IP address we listen to")
	portFlag := flag.String("port", "8080", "the port we listen to")
	flag.Parse()

	// temporary to test the logic
	store.StoreInit()

	listen := *ipFlag + ":" + *portFlag
	server := server.NewServer(listen)
	server.Serve()

}
