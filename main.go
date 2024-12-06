package main

import (
	"api/routes"
	"net/http"
)

func main() {

	routes.AlbumRoutes()
	routes.TrackRoutes()

	http.ListenAndServe(":8080", nil)

}
