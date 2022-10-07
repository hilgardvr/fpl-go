package main

import (
	"fmt"
	"hilgardvr/go-fpl/controllers"
	"hilgardvr/go-fpl/service"
	"log"
	"net/http"
)

func main() {
	// bootstrapData, err := clients.BootstrapData()
	err := service.InitAllPlayerStats()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/filter", controllers.Filter)
	fmt.Println("server listing on port 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
