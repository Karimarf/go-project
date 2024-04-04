package ViewModel

import (
	"fmt"
	"src/Repository"
)

func AddRooms() {
	fmt.Println("Entrez le nom de la chambre : ")
	var name string
	fmt.Scan(&name)

	fmt.Println("Entrez la capacité de la chambre : ")
	var capacity int
	fmt.Scan(&capacity)

	fmt.Println("Entrez la description de la chambre : ")
	var description string
	fmt.Scan(&description)

	Repository.CreateRoom(name, capacity, description)
}

func DeleteRoom() {
	fmt.Println("Entrez l'ID de la chambre à supprimer : ")
	var id int
	fmt.Scan(&id)

	Repository.DeleteRoom(id)
}

func ExportRoom() {
	println("1. Exporter en CSV")
	println("2. Exporter en JSON")
	println("3. Retour")

	var menu int
	fmt.Scan(&menu)

	for menu < 1 || menu > 3 {
		fmt.Println("Choisissez un nombre entre 1 et 3")
		ExportRoom()
		fmt.Scan(&menu)

	}

	switch menu {
	case 1:
		exportRoomCSV()
	case 2:
		exportRoomJSON()
	}

}
