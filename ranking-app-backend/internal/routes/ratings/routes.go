package ratings

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func RatingRoutes(db *sql.DB) chi.Router {
	router := chi.NewRouter()
	return router
}
