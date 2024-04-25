package handler

import (
	"html/template"
	"net/http"
	reservation "src/Repository"
	"strconv"
)

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("ViewWeb/templates/createRoom.html")
		tmpl.Execute(w, nil)

	} else if r.Method == "POST" {
		r.ParseForm()

		name := r.Form.Get("name")
		capacityStr := r.Form.Get("capacity")
		description := r.Form.Get("description")
		capacity, _ := strconv.Atoi(capacityStr)

		reservation.CreateRoom(name, capacity, description)

		http.Redirect(w, r, "/createRoom", http.StatusSeeOther)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
