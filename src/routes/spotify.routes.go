package routes

import (
	"guide/controllers"
	"net/http"
)

func spotifyRoutes(router *http.ServeMux) {
	router.HandleFunc("/", controllers.DisplayHome)
	router.HandleFunc("/collection", controllers.DisplayCollection)
	router.HandleFunc("/details", controllers.DisplayDetails)
	router.HandleFunc("/favoris", controllers.DisplayFavorites)
	router.HandleFunc("/add-favorite", controllers.AddToFavorites)
	router.HandleFunc("/about", controllers.DisplayAbout)
}
