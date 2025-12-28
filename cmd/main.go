package main

import (
	"log"

	"github.com/AnupamSingh2004/Go_Ecom_Backend/cmd/api"
)

func main(){
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}