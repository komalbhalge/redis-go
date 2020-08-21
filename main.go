package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	api "github.com/komalbhalge/redis-geo-go/api"
)

func main() {

	router := setupRoutes()
	log.Println("Server started to accept request in 8080 port.")
	http.ListenAndServe(":8080", router)

}
func setupRoutes() *httprouter.Router {
	router := httprouter.New()
	router.POST("/addlocation", api.AddLocation)
	router.POST("/searchLocation", api.SearchLocation)
	return router
}
