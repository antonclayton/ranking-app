package routes

import (
	"database/sql"
	"ranking-app-backend/internal/routes/places"
	"ranking-app-backend/internal/routes/products"
	"ranking-app-backend/internal/routes/ratings"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(db *sql.DB) chi.Router {
	router := chi.NewRouter()
	router.Mount("/api/v1/places", places.PlaceRoutes(db))
	router.Mount("/api/v1/products", products.ProductRoutes(db))
	router.Mount("/api/v1/ratings", ratings.RatingRoutes(db))
	return router
}
