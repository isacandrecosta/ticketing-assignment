package models

type Route struct {
	ID    string
	Stops []Stop
}

type Stop struct {
	Station  Station
	Distance int
}
