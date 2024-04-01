package ViewModel

import (
	"fmt"
	"src/Repository"
	"time"
)

func AddReservation() {
	fmt.Println("Entrez l'ID de la chambre : ")
	var roomId int
	fmt.Scan(&roomId)

	fmt.Println("Entrez la date de début (DD-MM-YYYY) : ")
	var startDate string
	fmt.Scan(&startDate)

	fmt.Println("Entrez la date de fin (DD-MM-YYYY) : ")
	var endDate string
	fmt.Scan(&endDate)

	now := time.Now()
	if startDate < now.Format("02-01-2006") || endDate < now.Format("02-01-2006") {
		fmt.Println("La date de début et de fin doivent être supérieures à la date d'aujourd'hui")
		return
	}
	if startDate > endDate {
		fmt.Println("La date de début doit être avant la date de fin")
		return
	}

	if Repository.VerifyResponsibilityRoom(roomId, startDate, endDate) {
		Repository.CreateReservation(roomId, startDate, endDate)
	} else {
		fmt.Println("La chambre n'est pas disponible à cette période")
	}
}

func DeleteReservation() {
	fmt.Println("Entrez l'ID de la réservation à supprimer : ")
	var id int
	fmt.Scan(&id)

	Repository.DeleteReservation(id)
}

func ExportReservation() {
	println("1. Exporter en CSV")
	println("2. Exporter en JSON")
	println("3. Retour")

	var menu int
	fmt.Scan(&menu)

	for menu < 1 || menu > 3 {
		fmt.Println("Choisissez un nombre entre 1 et 3")
		ExportReservation()
		fmt.Scan(&menu)
	}

	switch menu {
	case 1:
		exportReservationCSV()
	case 2:
		exportReservationJSON()
	}
}
