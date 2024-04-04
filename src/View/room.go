package View

import (
	"fmt"
	"src/Types"
)

func DisplayRooms(rooms []Types.Room) {
	fmt.Println("ID\tName\tCapacity\tDescription")
	for _, r := range rooms {
		fmt.Println(r.Id, r.Name, r.Capacity, r.Description)
	}
}
