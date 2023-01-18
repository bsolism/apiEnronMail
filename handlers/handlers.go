package handlers

import (
	"log"
	"net/http"
	"os"

	"example.com/apiEnronMail/routers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func Manejador() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/index", routers.Document)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3030"
	}

	handler := cors.AllowAll().Handler(r)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
