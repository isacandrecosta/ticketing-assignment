package models

type Ticket struct {
	ID          string
	Seat        Seat
	Service     Service
	Origin      Station
	Destination Station
}
