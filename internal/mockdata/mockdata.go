package mockdata

import (
	"ticketing-assignment/internal/models"
)

const (
	BookingID        = "Booking-id"
	ServiceID        = "5160"
	ServiceDate      = "2021-04-01"
	FirstClass       = "First-class"
	SecondClass      = "Second-class"
	CarriageA        = "A"
	CarriageH        = "H"
	CarriageN        = "N"
	CarriageT        = "T"
	SeatA11          = "A11"
	SeatA12          = "A12"
	SeatH1           = "H1"
	SeatN5           = "N5"
	SeatA1           = "A1"
	SeatT7           = "T7"
	SeatB1           = "B1"
	SeatB2           = "B2"
	StationParis     = "Paris"
	StationLondon    = "London"
	StationAmsterdam = "Amsterdam"
	StationBerlin    = "Berlin"
	StationBrussels  = "Brussels"
	StationCologne   = "Cologne"
	StationCalais    = "Calais"
	StationRotterdam = "Rotterdam"
)

func CreateParisToAmsterdamTwoFirstClassBooking() models.Booking {
	return models.Booking{
		ID: BookingID,
		Passengers: []models.Passenger{
			{
				Name: "John Doe",
				Tickets: []models.Ticket{
					{
						Seat: models.Seat{Number: SeatA11, Class: FirstClass},
						Service: models.Service{
							ID:   ServiceID,
							Date: ServiceDate,
							Route: models.Route{
								ID: "Paris-Amsterdam",
								Stops: []models.Stop{
									{Station: models.Station{Name: StationParis}, Distance: 0},
									{Station: models.Station{Name: StationBrussels}, Distance: 300},
									{Station: models.Station{Name: StationRotterdam}, Distance: 500},
									{Station: models.Station{Name: StationAmsterdam}, Distance: 700},
								},
							},
						},
						Origin:      models.Station{Name: StationParis},
						Destination: models.Station{Name: StationAmsterdam},
					},
				},
			},
			{
				Name: "Jane Doe",
				Tickets: []models.Ticket{
					{
						Seat: models.Seat{Number: SeatA12, Class: FirstClass},
						Service: models.Service{
							ID:   ServiceID,
							Date: ServiceDate,
							Route: models.Route{
								ID: "Paris-Amsterdam",
								Stops: []models.Stop{
									{Station: models.Station{Name: StationParis}, Distance: 0},
									{Station: models.Station{Name: StationBrussels}, Distance: 300},
									{Station: models.Station{Name: StationRotterdam}, Distance: 500},
									{Station: models.Station{Name: StationAmsterdam}, Distance: 700},
								},
							},
						},
						Origin:      models.Station{Name: StationParis},
						Destination: models.Station{Name: StationAmsterdam},
					},
				},
			},
		},
	}
}

func CreateLondonToAmsterdamOneFirstClassOneSecondClassBooking1() models.Booking {
	return models.Booking{
		ID: BookingID,
		Passengers: []models.Passenger{
			{
				Name: "Alice",
				Tickets: []models.Ticket{
					{
						Seat: models.Seat{Number: SeatH1, Class: SecondClass},
						Service: models.Service{
							ID:   ServiceID,
							Date: ServiceDate,
							Route: models.Route{
								ID: "London-Amsterdam",
								Stops: []models.Stop{
									{Station: models.Station{Name: StationLondon}, Distance: 0},
									{Station: models.Station{Name: StationCalais}, Distance: 100},
									{Station: models.Station{Name: StationParis}, Distance: 300},
									{Station: models.Station{Name: StationBrussels}, Distance: 500},
									{Station: models.Station{Name: StationAmsterdam}, Distance: 700},
								},
							},
						},
						Origin:      models.Station{Name: StationLondon},
						Destination: models.Station{Name: StationAmsterdam},
					},
				},
			},
			{
				Name: "Bob",
				Tickets: []models.Ticket{
					{
						Seat: models.Seat{Number: SeatA1, Class: FirstClass},
						Service: models.Service{
							ID:   ServiceID,
							Date: ServiceDate,
							Route: models.Route{
								ID: "London-Amsterdam",
								Stops: []models.Stop{
									{Station: models.Station{Name: StationLondon}, Distance: 0},
									{Station: models.Station{Name: StationCalais}, Distance: 100},
									{Station: models.Station{Name: StationParis}, Distance: 300},
									{Station: models.Station{Name: StationBrussels}, Distance: 500},
									{Station: models.Station{Name: StationAmsterdam}, Distance: 700},
								},
							},
						},
						Origin:      models.Station{Name: StationLondon},
						Destination: models.Station{Name: StationAmsterdam},
					},
				},
			},
		},
	}
}

func CreateAmsterdamToBerlinBooking() models.Booking {
	return models.Booking{
		ID: BookingID,
		Passengers: []models.Passenger{
			{
				Name: "Charlie",
				Tickets: []models.Ticket{
					{
						Seat: models.Seat{Number: SeatB1, Class: FirstClass},
						Service: models.Service{
							ID:   ServiceID,
							Date: ServiceDate,
							Route: models.Route{
								ID: "Amsterdam-Berlin",
								Stops: []models.Stop{
									{Station: models.Station{Name: StationAmsterdam}, Distance: 0},
									{Station: models.Station{Name: StationBrussels}, Distance: 200},
									{Station: models.Station{Name: StationCologne}, Distance: 400},
									{Station: models.Station{Name: StationBerlin}, Distance: 600},
								},
							},
						},
						Origin:      models.Station{Name: StationAmsterdam},
						Destination: models.Station{Name: StationBerlin},
					},
				},
			},
			{
				Name: "Dave",
				Tickets: []models.Ticket{
					{
						Seat: models.Seat{Number: SeatB2, Class: FirstClass},
						Service: models.Service{
							ID:   ServiceID,
							Date: ServiceDate,
							Route: models.Route{
								ID: "Amsterdam-Berlin",
								Stops: []models.Stop{
									{Station: models.Station{Name: StationAmsterdam}, Distance: 0},
									{Station: models.Station{Name: StationBrussels}, Distance: 200},
									{Station: models.Station{Name: StationCologne}, Distance: 400},
									{Station: models.Station{Name: StationBerlin}, Distance: 600},
								},
							},
						},
						Origin:      models.Station{Name: StationAmsterdam},
						Destination: models.Station{Name: StationBerlin},
					},
				},
			},
		},
	}
}
