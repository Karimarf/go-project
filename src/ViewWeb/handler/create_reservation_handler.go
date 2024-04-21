package handler

import (
	"html/template"
	"net/http"
	reservation "src/Repository"
	"strconv"
	"time"
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

		roomIdStr := r.Form.Get("roomId")
		startTime := r.Form.Get("startTime")
		endTime := r.Form.Get("endTime")

		roomId, err := strconv.Atoi(roomIdStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		parsedStartTime, err := time.Parse("2006-01-02", startTime)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		formattedStartTime := parsedStartTime.Format("02-01-2006")

		parsedEndTime, err := time.Parse("2006-01-02", endTime)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		formattedEndTime := parsedEndTime.Format("02-01-2006")

		reservation.CreateReservation(roomId, formattedStartTime, formattedEndTime)

		http.Redirect(w, r, "/createReservation", http.StatusSeeOther)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
