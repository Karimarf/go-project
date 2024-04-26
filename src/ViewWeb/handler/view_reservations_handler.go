package handler

import (
	"html/template"
	"net/http"
	"src/Repository"
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
