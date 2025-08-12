package places

import (
	"database/sql"

	"ranking-app-backend/internal/handlers/places"

	"github.com/go-chi/chi/v5"
)

func PlaceRoutes(db *sql.DB) chi.Router {
	router := chi.NewRouter()
	router.Get("/", places.ListPlaces(db))
	router.Post("/", places.CreatePlace(db))
	router.Get("/{id}", places.GetPlace(db))
	router.Put("/{id}", places.UpdatePlace(db))
	router.Delete("/{id}", places.DeletePlace(db))
	return router
}
