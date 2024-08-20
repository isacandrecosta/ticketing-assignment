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
