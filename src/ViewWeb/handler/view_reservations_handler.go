package handler

import (
	"html/template"
	"net/http"
)

func ViewReservationsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("ViewWeb/templates/viewReservations.html")

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
