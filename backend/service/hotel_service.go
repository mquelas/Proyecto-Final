package service

import (
	. "backend/data_access"
	. "backend/model"
	"database/sql"
)

func GetHotels() []Hotel {
	var err error
	var db *sql.DB

	db, err = DataConnect()
	if err != nil {
		//return nil, fmt.Errorf("getHotels %q: %v", err)
		return nil
	}

	var hotels []Hotel

	rows, err := db.Query("SELECT * FROM hotel")
	if err != nil {
		return nil
		//, fmt.Errorf("getHotels %q: %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.

	for rows.Next() {

		var hot Hotel
		if err := rows.Scan(&hot.ID, &hot.Name, &hot.Description, &hot.Price); err != nil {
			return nil
			//, fmt.Errorf("getHotels %q: %v", err)
		}
		hotels = append(hotels, hot)
	}
	if err := rows.Err(); err != nil {
		return nil
		//, fmt.Errorf("getHotels %q: %v", err)
	}
	return hotels
}

/*
func hotelsById(id string) ([]Hotel, error) {

	// An hotel slice to hold data from returned rows.

	var hotels []Hotel

	rows, err := db.Query("SELECT * FROM hotel WHERE Id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("hotelsById %q: %v", id, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.

	for rows.Next() {

		var hot Hotel
		if err := rows.Scan(&hot.ID, &hot.Name, &hot.Description, &hot.Price); err != nil {
			return nil, fmt.Errorf("hotelsById %q: %v", id, err)
		}
		hotels = append(hotels, hot)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("hotelsById %q: %v", id, err)
	}
	return hotels, nil
}
*/
