package places

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"ranking-app-backend/internal/models"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

func ListPlaces(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := `
			SELECT
				p.id, p.name, p.created_at, p.updated_at,
				GROUP_CONCAT(t.name) AS tags
			FROM
				places p
			LEFT JOIN
				place_tags pt ON p.id = pt.place_id
			LEFT JOIN
				tags t ON pt.tag_id = t.id
			GROUP BY
				p.id
		`
		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		places := []models.Place{}
		for rows.Next() {
			var place models.Place
			var tags sql.NullString // Use sql.NullString for tags

			if err := rows.Scan(&place.ID, &place.Name, &place.CreatedAt, &place.UpdatedAt, &tags); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if tags.Valid {
				place.Tags = strings.Split(tags.String, ",")
			} else {
				place.Tags = []string{}
			}
			places = append(places, place)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(places)
	}
}

func CreatePlace(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var place models.Place
		if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		tx, err := db.Begin()
		if err != nil {
			http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
			return
		}

		// Insert the place and get its ID
		var placeID int
		err = tx.QueryRow("INSERT INTO places (name) VALUES (?) RETURNING id", place.Name).Scan(&placeID)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to create place", http.StatusInternalServerError)
			return
		}
		place.ID = placeID

		// Handle tags
		if len(place.Tags) > 0 {
			for _, tagName := range place.Tags {
				// Find or create the tag
				var tagID int
				err := tx.QueryRow("INSERT INTO tags (name) VALUES (?) ON CONFLICT(name) DO UPDATE SET name=name RETURNING id", tagName).Scan(&tagID)
				if err != nil {
					tx.Rollback()
					http.Error(w, "Failed to create or find tag", http.StatusInternalServerError)
					return
				}
				// Associate tag with place
				_, err = tx.Exec("INSERT INTO place_tags (place_id, tag_id) VALUES (?, ?)", placeID, tagID)
				if err != nil {
					tx.Rollback()
					http.Error(w, "Failed to associate tag with place", http.StatusInternalServerError)
					return
				}
			}
		}

		if err := tx.Commit(); err != nil {
			http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(place)
	}
}

func GetPlace(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		query := `
			SELECT
				p.id, p.name, p.created_at, p.updated_at,
				GROUP_CONCAT(t.name) AS tags
			FROM
				places p
			LEFT JOIN
				place_tags pt ON p.id = pt.place_id
			LEFT JOIN
				tags t ON pt.tag_id = t.id
			WHERE
				p.id = ?
			GROUP BY
				p.id
		`
		var place models.Place
		var tags sql.NullString
		err := db.QueryRow(query, id).Scan(&place.ID, &place.Name, &place.CreatedAt, &place.UpdatedAt, &tags)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Place not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if tags.Valid {
			place.Tags = strings.Split(tags.String, ",")
		} else {
			place.Tags = []string{}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(place)
	}
}

func UpdatePlace(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		var place models.Place
		if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		tx, err := db.Begin()
		if err != nil {
			http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
			return
		}

		// Update place name
		_, err = tx.Exec("UPDATE places SET name = ?, updated_at = ? WHERE id = ?", place.Name, time.Now(), id)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to update place", http.StatusInternalServerError)
			return
		}

		// Delete old tags for the place
		_, err = tx.Exec("DELETE FROM place_tags WHERE place_id = ?", id)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to delete old tags", http.StatusInternalServerError)
			return
		}

		// Add new tags
		if len(place.Tags) > 0 {
			for _, tagName := range place.Tags {
				var tagID int
				err := tx.QueryRow("INSERT INTO tags (name) VALUES (?) ON CONFLICT(name) DO UPDATE SET name=name RETURNING id", tagName).Scan(&tagID)
				if err != nil {
					tx.Rollback()
					http.Error(w, "Failed to create or find tag", http.StatusInternalServerError)
					return
				}
				_, err = tx.Exec("INSERT INTO place_tags (place_id, tag_id) VALUES (?, ?)", id, tagID)
				if err != nil {
					tx.Rollback()
					http.Error(w, "Failed to associate tag with place", http.StatusInternalServerError)
					return
				}
			}
		}

		if err := tx.Commit(); err != nil {
			http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
			return
		}

		// Fetch and return the updated place
		GetPlace(db)(w, r)
	}
}

func DeletePlace(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		tx, err := db.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// First delete associations in place_tags
		_, err = tx.Exec("DELETE FROM place_tags WHERE place_id = ?", id)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to delete tag associations", http.StatusInternalServerError)
			return
		}

		// Then delete the place
		_, err = tx.Exec("DELETE FROM places WHERE id = ?", id)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to delete place", http.StatusInternalServerError)
			return
		}

		if err := tx.Commit(); err != nil {
			http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
