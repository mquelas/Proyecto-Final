package service

import (
	. "backend/data_access"
	. "backend/model"
	"database/sql"
	"fmt"
)

func GetHotels() ([]Hotel, error) {
	var err error
	var db *sql.DB

	db, err = DataConnect()
	if err != nil {
		return nil, fmt.Errorf("getHotels %q", err)
		//return nil
	}
	defer db.Close()

	var hotels []Hotel

	rows, err := db.Query("SELECT * FROM hotel")
	if err != nil {
		return nil, fmt.Errorf("getHotels %q: %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.

	for rows.Next() {

		var hot Hotel
		if err := rows.Scan(&hot.ID, &hot.Name, &hot.Description, &hot.Price, &hot.Rooms); err != nil {
			return nil, fmt.Errorf("getHotels %q: %v", err)
		}
		hotels = append(hotels, hot)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getHotels %q: %v", err)
	}

	return hotels, err
}

func GetHotelById(id string) Hotel {

	// An hotel slice to hold data from returned rows.

	var hotel Hotel
	var err error
	var db *sql.DB

	rows, err := db.Query("SELECT * FROM hotel WHERE Id = ?", id)
	if err != nil {
		//return nil
		return hotel
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.

	for rows.Next() {

		//var hot Hotel
		if err := rows.Scan(&hotel.ID, &hotel.Name, &hotel.Description, &hotel.Price); err != nil {
			//return nil
			return hotel
		}
		//hotels = append(hotels, hot)
	}
	if err := rows.Err(); err != nil {
		return hotel
	}
	return hotel
}
