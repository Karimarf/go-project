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
	if err != nil {
		http.Error(w, "Invalid reservation ID", http.StatusBadRequest)
		return
	}

	fmt.Printf("", id)

	Repository.DeleteReservation(id)

	http.Redirect(w, r, "/viewReservations", http.StatusSeeOther)
}
