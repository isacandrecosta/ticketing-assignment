package tests

import (
	"testing"
	"ticketing-assignment/internal/mockdata"
	"ticketing-assignment/internal/models"
	"ticketing-assignment/internal/services"
	"ticketing-assignment/tests/errors"
)

func TestSuccessfulBookingParisToAmsterdam(t *testing.T) {
	rs := &models.ReservationSystem{}
	reservationService := &services.ReservationService{ReservationSystem: rs}

	booking := mockdata.CreateParisToAmsterdamTwoFirstClassBooking()
	err := reservationService.MakeReservation(booking)
	if err != nil {
		t.Errorf(errors.ErrFailedBooking, err)
	}
}
func TestDuplicateBookingParisToAmsterdam(t *testing.T) {
	rs := &models.ReservationSystem{}
	reservationService := &services.ReservationService{ReservationSystem: rs}

	booking1 := mockdata.CreateParisToAmsterdamTwoFirstClassBooking()
	_ = reservationService.MakeReservation(booking1)

	err := reservationService.MakeReservation(booking1)
	if err == nil {
		t.Errorf(errors.ErrDuplicateBooking)
	}
}
func TestNumberOfStopsParisToAmsterdam(t *testing.T) {
	booking := mockdata.CreateParisToAmsterdamTwoFirstClassBooking()
	expectedStops := 4
	actualStops := len(booking.Passengers[0].Tickets[0].Service.Route.Stops)
	if actualStops != expectedStops {
		t.Errorf(errors.ErrNumberOfStops, expectedStops, actualStops)
	}
}

func TestTotalDistanceParisToAmsterdam(t *testing.T) {
	booking := mockdata.CreateParisToAmsterdamTwoFirstClassBooking()
	expectedDistance := 700
	actualDistance := booking.Passengers[0].Tickets[0].Service.Route.Stops[len(booking.Passengers[0].Tickets[0].Service.Route.Stops)-1].Distance
	if actualDistance != expectedDistance {
		t.Errorf(errors.ErrTotalDistance, expectedDistance, actualDistance)
	}
}

func TestSuccessfulBookingLondonToAmsterdam(t *testing.T) {
	rs := &models.ReservationSystem{}
	reservationService := &services.ReservationService{ReservationSystem: rs}

	booking := mockdata.CreateLondonToAmsterdamOneFirstClassOneSecondClassBooking1()
	err := reservationService.MakeReservation(booking)
	if err != nil {
		t.Errorf(errors.ErrFailedBooking, err)
	}
}

func TestSuccessfulBookingAmsterdamToBerlin(t *testing.T) {
	rs := &models.ReservationSystem{}
	reservationService := &services.ReservationService{ReservationSystem: rs}

	booking := mockdata.CreateAmsterdamToBerlinBooking()
	err := reservationService.MakeReservation(booking)
	if err != nil {
		t.Errorf(errors.ErrFailedBooking, err)
	}
}

func TestErrorCreatingReservation(t *testing.T) {
	rs := &models.ReservationSystem{}
	reservationService := &services.ReservationService{ReservationSystem: rs}

	booking1 := mockdata.CreateParisToAmsterdamTwoFirstClassBooking()
	err := reservationService.MakeReservation(booking1)
	if err != nil {
		t.Errorf(errors.ErrFailedBooking, err)
	}

	booking2 := mockdata.CreateParisToAmsterdamTwoFirstClassBooking()
	err = reservationService.MakeReservation(booking2)
	if err == nil || err.Error() != errors.ErrSeatAlreadyReserved {
		t.Errorf("Expected error %s, got %v", errors.ErrSeatAlreadyReserved, err)
	}
}
