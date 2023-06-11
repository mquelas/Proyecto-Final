package model

// estructura hotel

type Hotel struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Rooms       int64   `json:"rooms"`
}
