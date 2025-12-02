package model

import "time"

type Driver struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	FirstName string    `json:"firstName" bson:"firstName"`
	LastName  string    `json:"lastName" bson:"lastName"`
	Plate     string    `json:"plate" bson:"plate"`
	TaxiType  string    `json:"taxiType" bson:"taxiType"`
	CarBrand  string    `json:"carBrand" bson:"carBrand"`
	CarModel  string    `json:"carModel" bson:"carModel"`
	Lat       float64   `json:"lat" bson:"lat"`
	Lon       float64   `json:"lon" bson:"lon"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
