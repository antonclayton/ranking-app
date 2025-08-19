package models

import "time"

// Place represents a restaurant, cafe, or other establishment.
type Place struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Tags      []string  `json:"tags"` // e.g., ["restaurant", "cafe", "bar"]
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Product represents an item that can be rated, associated with a Place.
type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	PlaceID     int       `json:"place_id"` // Foreign key to Place
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Rating represents a user's rating for a Place or a Product.
type Rating struct {
	ID         int       `json:"id"`
	TargetID   int       `json:"target_id"`   // ID of the Place or Product being rated
	TargetType string    `json:"target_type"` // "place" or "product"
	Score      int       `json:"score"`       // e.g., 1-5
	Comment    string    `json:"comment,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}
