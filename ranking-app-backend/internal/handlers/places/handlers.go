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
		places := []models.Place{}

		rows, err := db.Query("SELECT id, name, types, created_at, updated_at FROM places")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var place models.Place
			var typesStr string
			var createdAt, updatedAt time.Time
			if err := rows.Scan(&place.ID, &place.Name, &typesStr, &createdAt, &updatedAt); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			place.CreatedAt = createdAt
			place.UpdatedAt = updatedAt
			if typesStr != "" {
				rawTypes := strings.Split(typesStr, ",")
				place.Types = make([]string, 0, len(rawTypes))
				for _, t := range rawTypes {
					trimmed := strings.TrimSpace(t)
					if trimmed != "" {
						place.Types = append(place.Types, trimmed)
					}
				}
			} else {
				place.Types = []string{}
			}
			places = append(places, place)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(places)
	}
}

func CreatePlace(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		place := models.Place{}
		if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := db.Exec("INSERT INTO places (name, types) VALUES (?, ?)", place.Name, strings.Join(place.Types, ","))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		place.ID = int(id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(place)
	}
}

func GetPlace(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "id is required", http.StatusBadRequest)
			return
		}

		var typesStr string
		place := models.Place{}
		err := db.QueryRow("SELECT id, name, types, created_at, updated_at FROM places WHERE id = ?", id).Scan(
			&place.ID,
			&place.Name,
			&typesStr,
			&place.CreatedAt,
			&place.UpdatedAt,
		)
		if err == sql.ErrNoRows {
			http.Error(w, "Place not found", http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if typesStr != "" {
			rawTypes := strings.Split(typesStr, ",")
			place.Types = make([]string, 0, len(rawTypes))
			for _, t := range rawTypes {
				trimmed := strings.TrimSpace(t)
				if trimmed != "" {
					place.Types = append(place.Types, trimmed)
				}
			}
		} else {
			place.Types = []string{}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(place)
	}
}

func UpdatePlace(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "id is required", http.StatusBadRequest)
			return
		}

		// Decode into a map to check what fields are present
		var updates map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fields := []string{}
		args := []interface{}{}

		if name, ok := updates["name"].(string); ok {
			fields = append(fields, "name = ?")
			args = append(args, name)
		}
		if types, ok := updates["types"].([]interface{}); ok {
			// Convert []interface{} to []string
			strTypes := make([]string, 0, len(types))
			for _, t := range types {
				if s, ok := t.(string); ok {
					strTypes = append(strTypes, s)
				}
			}
			fields = append(fields, "types = ?")
			args = append(args, strings.Join(strTypes, ","))
		}

		if len(fields) == 0 {
			http.Error(w, "no valid fields to update", http.StatusBadRequest)
			return
		}

		// Update the updated_at field
		fields = append(fields, "updated_at = ?")
		args = append(args, time.Now().UTC().Truncate(time.Second))
		args = append(args, id)
		query := "UPDATE places SET " + strings.Join(fields, ", ") + " WHERE id = ?"
		_, err := db.Exec(query, args...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Fetch the updated place
		var updatedPlace models.Place
		var typesStr string
		err = db.QueryRow("SELECT id, name, types, created_at, updated_at FROM places WHERE id = ?", id).Scan(
			&updatedPlace.ID,
			&updatedPlace.Name,
			&typesStr,
			&updatedPlace.CreatedAt,
			&updatedPlace.UpdatedAt,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if typesStr != "" {
			rawTypes := strings.Split(typesStr, ",")
			updatedPlace.Types = make([]string, 0, len(rawTypes))
			for _, t := range rawTypes {
				trimmed := strings.TrimSpace(t)
				if trimmed != "" {
					updatedPlace.Types = append(updatedPlace.Types, trimmed)
				}
			}
		} else {
			updatedPlace.Types = []string{}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedPlace)
	}
}

func DeletePlace(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "id is required", http.StatusBadRequest)
			return
		}

		// Check if the place exists
		var exists bool
		err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM places WHERE id = ?)", id).Scan(&exists)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !exists {
			http.Error(w, "place not found", http.StatusNotFound)
			return
		}

		_, err = db.Exec("DELETE FROM places WHERE id = ?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
