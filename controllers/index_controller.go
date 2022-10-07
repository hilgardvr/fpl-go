package controllers

import (
	"hilgardvr/go-fpl/service"
	"log"
	"net/http"
	"text/template"
)

type IndexData struct {
	Players []service.PlayerStatsAllTime
}

func Index(w http.ResponseWriter, r *http.Request) {
	stats := service.GetAllPlayerStats()
	tmpl, err := template.ParseFiles("./public/index.html")
	if err != nil {
		log.Fatalln("Could not parse index.html: ", err)
	}
	tmpl.Execute(w, IndexData{Players: stats})
}

func Filter(w http.ResponseWriter, r *http.Request) {
	playerType := r.FormValue("playerType")
	sortBy := r.FormValue("sortBy")
	filterSortRequest := service.FilterSortRequest{
		PlayerType: playerType,
		SortBy:     sortBy,
	}
	stats := service.FilterSortByRequest(filterSortRequest)
	tmpl, err := template.ParseFiles("./public/index.html")
	if err != nil {
		log.Fatalln("Could not parse index.html: ", err)
	}
	tmpl.Execute(w, IndexData{Players: stats})
}
