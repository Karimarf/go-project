package main

import (
	"fmt"
	"net/http"
	menu "src/Controller"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bienvenue sur le serveur de réservation !")
	})

	fmt.Println("http://localhost:8085")
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}

	menu.DisplayMenu()

}
