package main

import (
	"fmt"

	"encoding/json"
)

func main() {

	// Vessel
	boat := NewVessel("Renaissance")

	// Create a voyage, one or more passages
	v := NewVoyage(boat)

	v.AddPassage("Jervis Bay to Ulladulla")
	v.AddPassage("Ulladulla to Batemans Bay")
	v.AddPassage("Batemans Bay to Eden")

	// Passage 0 waypoints
	w1 := NewWaypoint(-35.24, 151.34)
	w2 := NewWaypoint(-35.45, 151.37)
	w3 := NewWaypoint(-35.75, 151.43)
	w4 := NewWaypoint(-36.75, 154.43)

	v.Passages[0].AddWaypoint(w1, -1)
	v.Passages[0].AddWaypoint(w2, -1)
	v.Passages[0].AddWaypoint(w3, -1)
	v.Passages[0].AddWaypoint(w4, 2)


	xb, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(xb))
}
