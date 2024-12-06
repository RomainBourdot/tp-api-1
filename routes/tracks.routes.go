package routes

import (
	"api/controllers"
	"net/http"
)

func TrackRoutes() {
	http.HandleFunc("/tracks", controllers.PageTrackDetails)
}
