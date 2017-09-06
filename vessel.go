package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Vessel struct {
	OID       bson.ObjectId `json:"_id"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	Name      string        `json:"name"`
}

func NewVessel(name string) Vessel {
	return Vessel{
		OID:       bson.NewObjectId(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}
}
