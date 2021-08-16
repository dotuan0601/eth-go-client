package main

import (
	"fmt"
	"log"
	"net/http"

	Handlers "github.com/LuisAcerv/goeth-api/handler"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

func main() {
	host := "http://localhost"
	networkId := 8501
	apiPort := 8080

	// connect to ETH
	client, err := ethclient.Dial(fmt.Sprintf("%s:%d", host, networkId))
	if err != nil {
		fmt.Println(err)
	}

	// Create a mux router
	r := mux.NewRouter()

	// start server
	r.Handle("/api/v1/eth/{module}", Handlers.ClientHandler{client})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", apiPort), r))
}
