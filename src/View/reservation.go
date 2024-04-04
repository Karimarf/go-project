package View

import (
	"fmt"
	"src/Types"
)

func DisplayReservations(reservations []Types.Reservation) {
	fmt.Println("ID\tRoom ID\tDate\tStart Time\tEnd Time")
	for _, r := range reservations {
		fmt.Println(r.Id, r.RoomId, r.StartTime, r.EndTime)
	}
}
