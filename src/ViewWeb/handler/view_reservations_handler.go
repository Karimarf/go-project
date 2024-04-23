package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"src/Repository"
	"strconv"
)

func ViewReservationsHandler(w http.ResponseWriter, r *http.Request) {
	reservations := Repository.GetReservations()
	tmpl, err := template.ParseFiles("ViewWeb/templates/viewReservations.html")

	err = tmpl.Execute(w, reservations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteReservationHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := r.Form.Get("id")

	IdReservation := strconv.Atoi(id)

	fmt.Printf("ID ---> %s\n", id)
	fmt.Printf("ID ---> %s\n", IdReservation)

	err = Repository.DeleteReservation(IdReservation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/viewReservations", http.StatusSeeOther)
}
