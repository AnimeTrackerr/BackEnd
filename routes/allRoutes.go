package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AnimeTrackerr/v2/backend/handlers"
	"github.com/gorilla/mux"
)

func AddHandlers(r *mux.Router) {
	AnimeRoutes(r)
	MangaRoutes(r)
}

func SetRoutes() {
	val, present := os.LookupEnv("PORT_NO")
	var PORT_NO string = ":80"
	handlers.SetCollection(os.Getenv("MONGODB_URI"))

	if present {
		PORT_NO = val
	}

	r := mux.NewRouter()
	AddHandlers(r)

	fmt.Printf("Starting server on PORT%s ....\n", PORT_NO)
	log.Fatal(http.ListenAndServe("127.0.0.1" + PORT_NO, r))
}