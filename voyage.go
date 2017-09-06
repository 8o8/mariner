package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Voyage struct {
	Vessel      `json:"vessel"`
	OID         bson.ObjectId `json:"_id"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	StartAt     time.Time     `json:"startAt"`
	EndAt       time.Time     `json:"endAt"`
	Description string        `json:"description"`
	Passages    []Passage     `json:"passages"`
}

func NewVoyage(vessel Vessel) *Voyage {
	return &Voyage{
		OID:       bson.NewObjectId(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Vessel:    vessel,
	}
}

func (v *Voyage) AddPassage(desc string) {
	p := NewPassage()
	p.Description = desc
	v.Passages = append(v.Passages, *p)
}
