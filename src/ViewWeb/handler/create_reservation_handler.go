package handler

import (
	"html/template"
	"net/http"
	reservation "src/Repository"
	"strconv"
)

func CreateReservationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("ViewWeb/templates/createReservation.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		roomId := r.Form.Get("roomId")
		startTime := r.Form.Get("startTime")
		endTime := r.Form.Get("endTime")

		roomIdStr, err := strconv.Atoi(roomId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		reservation.CreateReservation(roomIdStr, startTime, endTime)

		http.Redirect(w, r, "/createReservation", http.StatusSeeOther)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
