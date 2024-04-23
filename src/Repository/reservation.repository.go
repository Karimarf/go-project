package Repository

import (
	"src/Types"
)

func CreateRoom(name string, capacity int, description string) {
	db := createConnection()
	_, err := db.Exec("INSERT INTO room (name, capacity, description) VALUES (?, ?, ?)", name, capacity, description)
	if err != nil {
		panic(err.Error())
	}
	defer closeConnection(db)
}

func DeleteRoom(id int) {
	db := createConnection()
	_, err := db.Exec("DELETE FROM room WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer closeConnection(db)
}

func CreateReservation(roomId int, startTime string, endTime string) {
	db := createConnection()
	startTime = startTime[6:] + "-" + startTime[3:5] + "-" + startTime[:2]
	endTime = endTime[6:] + "-" + endTime[3:5] + "-" + endTime[:2]
	_, err := db.Exec("INSERT INTO reservation (room_id, start_date, end_date) VALUES (?, ?, ?)", roomId, startTime, endTime)
	if err != nil {
		panic(err.Error())
	}
	defer closeConnection(db)
}

func DeleteReservation(id int) {
	db := createConnection()
	_, err := db.Exec("DELETE FROM reservation WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer closeConnection(db)
}

func VerifyResponsibilityRoom(roomId int, startTime string, endTime string) bool {
	db := createConnection()
	rows, err := db.Query("SELECT start_date, end_date FROM reservation WHERE room_id = ?", roomId)
	if err != nil {
		panic(err.Error())
	}
	isAvailable := true
	for rows.Next() {
		var startDate string
		var endDate string
		err = rows.Scan(&startDate, &endDate)
		if err != nil {
			panic(err.Error())
		}
		if startTime < startDate && endTime > startDate {
			isAvailable = false
		}
		if startTime < endDate && endTime > endDate {
			isAvailable = false
		}
	}
	defer rows.Close()
	defer closeConnection(db)
	return isAvailable
}

func GetRooms() []Types.Room {
	db := createConnection()
	rows, err := db.Query("SELECT id, name, capacity, description FROM room")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var rooms []Types.Room
	for rows.Next() {
		var id int
		var name string
		var capacity int
		var description string
		err = rows.Scan(&id, &name, &capacity, &description)
		if err != nil {
			panic(err.Error())
		}
		rooms = append(rooms, Types.Room{id, name, capacity, description})
	}
	defer closeConnection(db)
	return rooms
}

func GetReservations() []Types.Reservation {
	db := createConnection()
	rows, err := db.Query("SELECT id, room_id, start_date, end_date FROM reservation")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var reservations []Types.Reservation
	for rows.Next() {
		var id int
		var roomId int
		var startTime string
		var endTime string
		err = rows.Scan(&id, &roomId, &startTime, &endTime)
		if err != nil {
			panic(err.Error())
		}
		reservations = append(reservations, Types.Reservation{id, roomId, startTime, endTime})
	}
	defer closeConnection(db)
	return reservations
}
