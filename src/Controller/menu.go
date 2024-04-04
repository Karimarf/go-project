package Controller

import (
	"fmt"
	"os"
	"src/Repository"
	"src/View"
	"src/ViewModel"
)

func MenuController(menu int) {
	switch menu {
	case 1:
		rooms := Repository.GetRooms()
		View.DisplayRooms(rooms)
	case 2:
		ViewModel.AddReservation()
	case 3:
		ViewModel.DeleteReservation()
	case 4:
		reservations := Repository.GetReservations()
		View.DisplayReservations(reservations)
	case 5:
		ViewModel.AddRooms()
	case 6:
		ViewModel.DeleteRoom()
	case 7:
		ViewModel.ExportReservation()
	case 8:
		ViewModel.ExportRoom()
	case 9:
		println("Voulez-vous vraiment quitter le programme ? ( oui/non) ")
		var leave string = ""
		fmt.Scan(&leave)
		if leave == "oui" {
			os.Exit(1)
		}
	}

}

func DisplayMenu() {
	fmt.Println("Bienvenue dans le Service de Réservation en Ligne")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("")
	fmt.Println("1. Lister les salles disponibles")
	fmt.Println("2. Créer une réservation")
	fmt.Println("3. Annuler une réservation")
	fmt.Println("4. Visualiser les réservations")
	fmt.Println("5. Créer une salle")
	fmt.Println("6. Supprimer une salle")
	fmt.Println("7. Exporter les réservations")
	fmt.Println("8. Exporter les salles")
	fmt.Println("9. Quitter")
	fmt.Print("Choisissez une option : ")

	var menu int
	fmt.Scan(&menu)

	for menu < 0 || menu > 9 {
		fmt.Println("Choisissez un nombre entre 1 et 9")
		DisplayMenu()
		fmt.Scan(&menu)
	}

	MenuController(menu)

	DisplayMenu()

}
