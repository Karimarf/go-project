// view_rooms_handler.go

package handler

import (
	"html/template"
	"net/http"
	"src/Repository"
)

func ViewRoomsHandler(w http.ResponseWriter, r *http.Request) {
	rooms := Repository.GetRooms()
	tmpl, err := template.ParseFiles("ViewWeb/templates/viewRooms.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, rooms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
