package main

import (
	"fmt"

	"encoding/json"
)

func main() {

	// Vessel
	boat := NewVessel("Renaissance")

	// Create a voyage, description and a few passages...
	v := NewVoyage(boat)
	v.Description = "Jervis Bay to Eden"
	v.AddPassage("Jervis Bay to Ulladulla")
	v.AddPassage("Ulladulla to Bateman's Bay")
	v.AddPassage("Batemans Bay to Eden")

	// Passage 0 waypoints
	w1 := NewWaypoint(-34.9989, 150.7261)
	w2 := NewWaypoint(-35.1389, 150.8040)
	w3 := NewWaypoint(-35.3264, 150.4939)
	w4 := NewWaypoint(-35.3568, 150.4852)

	v.Passages[0].AddWaypoint(w1, -1)
	v.Passages[0].AddWaypoint(w2, -1)
	v.Passages[0].AddWaypoint(w3, -1)
	v.Passages[0].AddWaypoint(w4, 2)


	xb, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(xb))
}
