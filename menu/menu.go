package menu

import (
	"fmt"
	"os"
)

func PrintMenu() {
	fmt.Println("Bienvenue dans le Service de Réservation en Ligne")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("")
	fmt.Println("1. Lister les salles disponibles")
	fmt.Println("2. Créer une réservation")
	fmt.Println("3. Annuler une réservation")
	fmt.Println("4. Visualiser les réservations")
	fmt.Println("5. Quitter")
	fmt.Println("")
	fmt.Print("Choisissez une option : ")

	var menu int
	fmt.Scan(&menu)

	for menu < 0 || menu > 5 {
		fmt.Println("Choisissez un nombre entre 1 et 5")
		PrintMenu()
		fmt.Scan(&menu)
	}

	menuRedirect(menu)

}

func menuRedirect(menu int) {
	switch menu {
	case 1:
		fmt.Println("Fonction 1")
		PrintMenu()
	case 2:
		//fonction 2
		PrintMenu()
	case 3:
		//fonction 3
		PrintMenu()
	case 4:
		//fonction 4
		PrintMenu()
	case 5:
		println("Voulez-vous vraiment quitter le programme ? ( oui/non) ")
		var leave string = ""
		fmt.Scan(&leave)
		if leave == "oui" {
			os.Exit(1)
		}
		PrintMenu()
	}

}
