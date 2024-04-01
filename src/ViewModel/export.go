package ViewModel

import (
	"encoding/json"
	"log"
	"os"
	"src/Repository"
)

func exportRoomCSV() {
	roms := Repository.GetRooms()
	text := "ID,Name,Capacity,Description\n"
	for _, r := range roms {
		text += string(r.Id) + "," + r.Name + "," + string(r.Capacity) + "," + r.Description + "\n"

	}
	err := os.WriteFile("export.csv", []byte(text), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func exportReservationCSV() {
	reservations := Repository.GetReservations()
	text := "ID,RoomID,StartDate,EndDate\n"
	for _, r := range reservations {
		text += string(r.Id) + "," + string(r.RoomId) + "," + r.StartTime + "," + r.EndTime + "\n"
	}
	err := os.WriteFile("export.csv", []byte(text), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func exportRoomJSON() {
	rooms := Repository.GetRooms()
	jsonData, _ := json.Marshal(rooms)
	err := os.WriteFile("export.json", []byte(string(jsonData)), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func exportReservationJSON() {
	reservations := Repository.GetReservations()
	jsonData, _ := json.Marshal(reservations)
	err := os.WriteFile("export.json", []byte(string(jsonData)), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
