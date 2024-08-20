package models

type ReservationSystem struct {
	Bookings []Booking
}

func (rs *ReservationSystem) IsSeatTaken(serviceID, date, seatNumber string) bool {
	for _, booking := range rs.Bookings {
		for _, passenger := range booking.Passengers {
			for _, ticket := range passenger.Tickets {
				if ticket.Service.ID == serviceID && ticket.Service.Date == date && ticket.Seat.Number == seatNumber {
					return true
				}
			}
		}
	}
	return false
}

func (rs *ReservationSystem) CountPassengersAtStation(stationName string, boarding bool) int {
	count := 0
	for _, booking := range rs.Bookings {
		for _, passenger := range booking.Passengers {
			for _, ticket := range passenger.Tickets {
				if (boarding && ticket.Origin.Name == stationName) || (!boarding && ticket.Destination.Name == stationName) {
					count++
				}
			}
		}
	}
	return count
}

func (rs *ReservationSystem) FindPassengerInSeat(serviceID, date, carriage, seatNumber string) *Passenger {
	for _, booking := range rs.Bookings {
		for _, passenger := range booking.Passengers {
			for _, ticket := range passenger.Tickets {
				if ticket.Service.ID == serviceID && ticket.Service.Date == date && ticket.Seat.Number == seatNumber && ticket.Seat.Class == carriage {
					return &passenger
				}
			}
		}
	}
	return nil
}

func (rs *ReservationSystem) AddBooking(booking Booking) {
	rs.Bookings = append(rs.Bookings, booking)
}
