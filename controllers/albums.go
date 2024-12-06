package controllers

import (
	"api/spotify"
	"fmt"
	"html/template"
	"net/http"
)

func PageListAlbums(w http.ResponseWriter, r *http.Request) {
	albums, err := spotify.GetAlbums()
	fmt.Println(albums)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des albums", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("./templates/albums.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, albums)
}
