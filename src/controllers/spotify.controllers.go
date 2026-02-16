package controllers

import (
	"guide/helper"
	"guide/services"
	"net/http"
	"strings"
)

func DisplayHome(w http.ResponseWriter, r *http.Request) {
	helper.RenderTemplate(w, r, "index", nil)
}

func DisplayCollection(w http.ResponseWriter, r *http.Request) {
	allTracks := services.GetMyPlaylist()
	query := strings.ToLower(r.URL.Query().Get("search"))

	var results []services.Track
	for _, t := range allTracks {
		if query == "" || strings.Contains(strings.ToLower(t.Name), query) || 
		   strings.Contains(strings.ToLower(t.Artist), query) {
			results = append(results, t)
		}
	}
	helper.RenderTemplate(w, r, "collection", results)
}

func DisplayDetails(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	track := services.GetTrackByID(id)
	helper.RenderTemplate(w, r, "list_tracks", track) 
}

func DisplayFavorites(w http.ResponseWriter, r *http.Request) {
	favs := services.LoadFavorites()

	helper.RenderTemplate(w, r, "collection", favs) 
}

func AddToFavorites(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	services.SaveToFavorites(id)
	http.Redirect(w, r, "/favoris", http.StatusSeeOther)
}

func DisplayAbout(w http.ResponseWriter, r *http.Request) {
	helper.RenderTemplate(w, r, "about", nil)
}