package routes

import (
	"github.com/AnimeTrackerr/v2/backend/handlers"
	"github.com/gorilla/mux"
)

func AnimeRoutes(r *mux.Router) {
	r.HandleFunc("/", handlers.Default).Methods("GET")
	r.HandleFunc("/anime/{id:[0-9]+}", handlers.GetAnime).Methods("GET")
	r.HandleFunc("/anime/random", handlers.GetRandomAnime).Methods("GET")
	r.HandleFunc("/animelist", handlers.GetAllAnime).Methods("GET")
	r.HandleFunc("/anime/search", handlers.SearchAnime).Methods("GET")
}