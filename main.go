package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/pavva91/golang-api-structure/api"
	"github.com/pavva91/golang-api-structure/storage"
)

func main()  {
	listenAddr := flag.String("listenaddr", ":3000", "the server address")
	flag.Parse()

	store := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddr, store)
	fmt.Println("server running on port:", *listenAddr)
	log.Fatal(server.Start())
	
}
