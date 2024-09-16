package main

import "fmt"

func getVehicleById(id int8) string {
	switch id {
	case 1:
		return "car"
	case 2:
		return "bike"
	case 3:
		return "truck"
	case 4:
		return "motorcycle"
	default:
		return "Vehicle not exist"
	}
}

func main() {
	vehicle := getVehicleById(2)

	fmt.Println(vehicle)
}