package services

import (
	"encoding/json"
	"errors"
	"net/http"
	intErrors "ticketing-assignment/internal/errors"
	"ticketing-assignment/internal/models"
	"ticketing-assignment/internal/views"
)

type ReservationService struct {
	ReservationSystem *models.ReservationSystem
}

func (rs *ReservationService) MakeReservation(booking models.Booking) error {
	for _, passenger := range booking.Passengers {
		for _, ticket := range passenger.Tickets {
			if rs.ReservationSystem.IsSeatTaken(ticket.Service.ID, ticket.Service.Date, ticket.Seat.Number) {
				return errors.New(intErrors.ErrSeatAlreadyReserved)
			}
		}
	}
	rs.ReservationSystem.Bookings = append(rs.ReservationSystem.Bookings, booking)
	return nil
}

func (rs *ReservationService) HandleCreateBooking(r *http.Request) views.Response {
	var booking models.Booking
	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		return views.Response{
			StatusCode: http.StatusBadRequest,
			Body:       intErrors.ErrInvalidRequestPayload,
		}
	}

	err = rs.MakeReservation(booking)
	if err != nil {
		return views.Response{
			StatusCode: http.StatusConflict,
			Body:       err.Error(),
		}
	}

	return views.Response{
		StatusCode: http.StatusCreated,
		Body:       booking,
	}
}

func (rs *ReservationService) IsSeatReserved(serviceID, date, seatNumber string) bool {
	for _, booking := range rs.ReservationSystem.Bookings {
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

func (rs *ReservationService) CountPassengersAtStation(stationName string, boarding bool) int {
	count := 0
	for _, booking := range rs.ReservationSystem.Bookings {
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

func (rs *ReservationService) FindPassengerInSeat(serviceID, date, carriage, seatNumber string) *models.Passenger {
	for _, booking := range rs.ReservationSystem.Bookings {
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

func (rs *ReservationService) AddBooking(booking models.Booking) {
	rs.ReservationSystem.Bookings = append(rs.ReservationSystem.Bookings, booking)
}
