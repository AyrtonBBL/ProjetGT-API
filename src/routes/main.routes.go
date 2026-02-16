package routes

import "net/http"

func MainRouter() *http.ServeMux {
	mainRouter := http.NewServeMux()

	spotifyRoutes(mainRouter)

	fileServer := http.FileServer(http.Dir("./../assets"))
	mainRouter.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mainRouter
}