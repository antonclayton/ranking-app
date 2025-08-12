package products

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func ProductRoutes(db *sql.DB) chi.Router {
	router := chi.NewRouter()
	return router
}
