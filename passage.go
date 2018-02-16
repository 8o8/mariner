package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Passage describes part of a Voyage. A Voyage may consist of one of more Passages.
// A Passage contains a set of navigation Waypoints.
type Passage struct {
	OID         bson.ObjectId `json:"_id"`
	Description string        `json:"name"`
	Waypoints   []Waypoint    `json:"waypoints"`
}

type Waypoint struct {
	OID bson.ObjectId `json:"_id"`

	// time.Time is a struct so omitempty doesn't work. Use *time.Time
	// as omitempty works for nil pointer value
	ETA *time.Time    `json:"eta,omitempty"`
	ATA *time.Time    `json:"ata,omitempty"`

	// Latitude and Longitude stored as float64 and converted to other formats as needed.
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// NewPassage returns a pointer to a Passage initialised with an OID value.
func NewPassage() *Passage {
	return &Passage{
		OID: bson.NewObjectId(),
	}
}

// NewWaypoint returns an initialised Waypoint.
func NewWaypoint(lat, lon float64) Waypoint {
	return Waypoint{
		OID: bson.NewObjectId(),
		Lat: lat,
		Lon: lon,
	}
}

// AddWaypoint adds a Waypoint to a Passage, at position specified by i.
// If n < 0, or out of range, the Waypoint is appended, otherwise it is
// inserted at the position specified.
func (p *Passage) AddWaypoint(w Waypoint, i int) {

	// Append if i out of range
	if i < 0 || i > len(p.Waypoints) {
		p.Waypoints = append(p.Waypoints, w)
		return
	}

	// Insert at position i
	newWaypoints := []Waypoint{}
	for c := 0; c < len(p.Waypoints); c++ {
		if c == i {
			newWaypoints = append(newWaypoints, w)
		}
		newWaypoints = append(newWaypoints, p.Waypoints[c])
	}
	p.Waypoints = newWaypoints
}
