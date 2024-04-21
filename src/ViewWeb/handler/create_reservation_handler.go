package handler

import (
	"html/template"
	"net/http"
)

func CreateReservationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("ViewWeb/templates/createReservation.html")

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
