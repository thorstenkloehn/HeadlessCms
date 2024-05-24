package model

type Restaurant struct {
	ID int64
	Name string
	Kategorie []string
	Location string //WKT representation of the location
}