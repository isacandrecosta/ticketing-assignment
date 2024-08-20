package models

type Service struct {
	ID        string
	Carriages []Carriage
	Route     Route
	Date      string
}
