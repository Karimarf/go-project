package ViewWeb

import (
	"fmt"
	"net/http"
	"src/ViewWeb/handler"
)

func Viewpage() {

	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/createReservation", handler.CreateReservationHandler)
	http.HandleFunc("/cancelReservation", handler.CancelReservationHandler)
	http.HandleFunc("/viewReservations", handler.ViewReservationsHandler)

	fmt.Println("Serveur de réservation démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}