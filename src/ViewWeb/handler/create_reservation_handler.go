package handler

import (
	"html/template"
	"net/http"
	reservation "src/Repository"
	"strconv"
)

func CreateReservationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("ViewWeb/templates/createReservation.html")
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		roomId := r.Form.Get("roomId")
		startTime := r.Form.Get("startTime")
		endTime := r.Form.Get("endTime")

		roomIdInt, _ := strconv.Atoi(roomId)
		reservation.CreateReservation(roomIdInt, startTime, endTime)

		http.Redirect(w, r, "/createReservation", http.StatusSeeOther)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
