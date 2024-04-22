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

	var startTimeDate, err = time.Parse("02-01-2006", startDate)
	if err != nil {
		fmt.Println("La date de début est invalide")
		return
	}
	var endTimeDate, err2 = time.Parse("02-01-2006", endDate)
	if err2 != nil {
		fmt.Println("La date de fin est invalide")
		return
	}
	if startTimeDate.After(endTimeDate) {
		fmt.Println("La date de début doit être avant la date de fin")
		return
	}
	if startTimeDate.Before(time.Now()) {
		fmt.Println("La date de début doit être après aujourd'hui")
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
