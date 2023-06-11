package service

import (
	. "backend/data_access"
	. "backend/model"
	"context"
	"database/sql"
	"fmt"
	"log"
)

func GetHotels() ([]Hotel, error) {

	var err error
	var db *sql.DB

	db, err = DataConnect()
	if err != nil {

		return nil, fmt.Errorf("getHotels %q", err)
	}
	defer db.Close()

	var hotels []Hotel

	rows, err := db.Query("SELECT * FROM hotel")
	if err != nil {
		return nil, fmt.Errorf("getHotels %q", err)
	}
	defer rows.Close()

	for rows.Next() {

		var hot Hotel
		if err := rows.Scan(&hot.ID, &hot.Name, &hot.Description, &hot.Price, &hot.Rooms); err != nil {
			return nil, fmt.Errorf("getHotels %q", err)
		}
		hotels = append(hotels, hot)
	}
	if err := rows.Err(); err != nil {

		return nil, fmt.Errorf("getHotels %q", err)
	}

	return hotels, nil
}

func GetHotelById(id string) (Hotel, error) {

	var hotel Hotel
	var err error
	var db *sql.DB

	db, err = DataConnect()
	if err != nil {

		return hotel, fmt.Errorf("getHotelById %q", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM hotel WHERE id_hotel = ?", id)
	if err != nil {

		return hotel, fmt.Errorf("getHotelById %q", err)
	}
	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&hotel.ID, &hotel.Name, &hotel.Description, &hotel.Price, &hotel.Rooms); err != nil {
			return hotel, fmt.Errorf("getHotelById %q", err)
		}
	}

	if err := rows.Err(); err != nil {
		return hotel, fmt.Errorf("getHotelById %q", err)
	}

	return hotel, nil
}

func CreateHotel(hotel *Hotel) (*Hotel, error) {

	var err error
	var db *sql.DB

	db, err = DataConnect()

	if err != nil {
		return nil, fmt.Errorf("createHotel: %s", err)
	}
	defer db.Close()

	insertResult, err := db.ExecContext(

		context.Background(),

		"INSERT INTO hotels (id, name, description, price, rooms) VALUES (?, ?, ?, ?, ?)",
		hotel.ID, hotel.Name, hotel.Description, hotel.Price, hotel.Rooms)

	if err != nil {

		return nil, fmt.Errorf("impossible insert hotel: %s", err)
	}
	id, err := insertResult.LastInsertId()

	if err != nil {

		return nil, fmt.Errorf("impossible to retrieve last inserted id: %s", err)
	}

	log.Printf("inserted id: %d", id)

	return hotel, nil
}
