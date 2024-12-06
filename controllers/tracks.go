package controllers

import (
	"api/spotify"
	"html/template"
	"net/http"
)

func PageTrackDetails(w http.ResponseWriter, r *http.Request) {
	track, err := spotify.GetTrackDetails()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des détails de la chanson", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("./templates/track.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, track)
}
