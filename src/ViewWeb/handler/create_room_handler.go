package handler

import (
	"html/template"
	"net/http"
	reservation "src/Repository"
	"strconv"
)

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Afficher le formulaire de création de chambre
		tmpl, err := template.ParseFiles("ViewWeb/templates/createRoom.html")
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
		// Analyser les données du formulaire
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Récupérer les données du formulaire
		name := r.Form.Get("name")
		capacityStr := r.Form.Get("capacity")
		description := r.Form.Get("description")

		// Convertir la capacité de chaîne de caractères en un entier
		capacity, err := strconv.Atoi(capacityStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Appeler la fonction CreateRoom avec les données du formulaire
		reservation.CreateRoom(name, capacity, description)

		// Rediriger vers la page de création de chambre
		http.Redirect(w, r, "/createRoom", http.StatusSeeOther)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
