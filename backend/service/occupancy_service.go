package service

import (
	. "backend/data_access"
	"database/sql"
	"fmt"
	"time"
)

func GetOccupancyByDate(idHotel int64, dateFrom time.Time, dateTo time.Time) (int64, error) {
	var err error
	var db *sql.DB

	db, err = DataConnect()
	if err != nil {
		return -1, fmt.Errorf("GetOccupancyByDate %q", err)
	}

	defer db.Close()

	rows, err := db.Query("select count(1) from ( "+
		"SELECT "+
		"distinct(O.id_reservation) "+
		"FROM "+
		"occupancy O "+
		"JOIN "+
		"reservation R ON (O.id_reservation = R.id_reservation) "+
		"WHERE "+
		"R.id_hotel = ? "+
		"AND O.date >= ? "+
		"AND O.date <= ? "+
		"group by id_reservation) T ",

		idHotel, dateFrom, dateTo)

	if err != nil {
		return -1, fmt.Errorf("GetOccupancyByDate %q", err)
	}

	defer rows.Close()

	var count int64
	for rows.Next() {

		if err := rows.Scan(&count); err != nil {
			return -1, fmt.Errorf("GetOccupancyByDate %q", err)
		}
	}

	if err := rows.Err(); err != nil {
		return -1, fmt.Errorf("GetOccupancyByDate %q", err)
	}

	return count, nil
}
