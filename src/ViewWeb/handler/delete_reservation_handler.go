package handler

import (
	"fmt"
	"net/http"
	"src/Repository"
	"strconv"
)

func DeleteReservationHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.Form.Get("id")
	id, err := strconv.Atoi(idStr)

	fmt.Printf("ID: %d\n", id)

	err = Repository.DeleteReservation(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/viewReservations", http.StatusSeeOther)
}
