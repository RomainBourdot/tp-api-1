package routes

import (
	"api/controllers"
	"net/http"
)

func AlbumRoutes() {
	http.HandleFunc("/albums", controllers.PageListAlbums)
}
