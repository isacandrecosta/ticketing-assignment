package tests

import (
	"testing"
	"ticketing-assignment/internal/mockdata"
	"ticketing-assignment/internal/models"
	"ticketing-assignment/internal/services"
	"ticketing-assignment/tests/errors"
)

func TestPassengersBoardingAtLondon(t *testing.T) {
	rs := &models.ReservationSystem{}
	reservationService := &services.ReservationService{ReservationSystem: rs}

	reservationService.MakeReservation(mockdata.CreateLondonToAmsterdamOneFirstClassOneSecondClassBooking1())

	passengersBoardingLondon := rs.CountPassengersAtStation("London", true)
	if passengersBoardingLondon != 2 {
		t.Errorf(errors.ErrPassengersBoarding, 2, "London", passengersBoardingLondon)
	}
}

func TestPassengersLeavingAtParis(t *testing.T) {
	rs := &models.ReservationSystem{}
	reservationService := &services.ReservationService{ReservationSystem: rs}

	reservationService.MakeReservation(mockdata.CreateLondonToAmsterdamOneFirstClassOneSecondClassBooking1())

	passengersLeavingParis := rs.CountPassengersAtStation("Paris", false)
	if passengersLeavingParis != 0 {
		t.Errorf(errors.ErrPassengersLeaving, 0, "Paris", passengersLeavingParis)
	}
}

func TestPassengersBetweenCalaisAndParis(t *testing.T) {
	rs := &models.ReservationSystem{}
	reservationService := &services.ReservationService{ReservationSystem: rs}

	reservationService.MakeReservation(mockdata.CreateLondonToAmsterdamOneFirstClassOneSecondClassBooking1())

	passengersBetweenCalaisParis := rs.CountPassengersAtStation("Calais", true) - rs.CountPassengersAtStation("Paris", false)
	if passengersBetweenCalaisParis != 0 {
		t.Errorf(errors.ErrPassengersBetweenStations, 0, "Calais", "Paris", passengersBetweenCalaisParis)
	}
}

func TestPassengerInSeatA11(t *testing.T) {
	rs := &models.ReservationSystem{}
	reservationService := &services.ReservationService{ReservationSystem: rs}

	booking := mockdata.CreateParisToAmsterdamTwoFirstClassBooking()
	reservationService.MakeReservation(booking)

	t.Logf("Booking added: %+v", booking)

	passengerInA11 := rs.FindPassengerInSeat(mockdata.ServiceID, mockdata.ServiceDate, "First-class", mockdata.SeatA11)
	if passengerInA11 == nil || passengerInA11.Name != "John Doe" {
		t.Errorf(errors.ErrPassengerNotFound, "John Doe", mockdata.SeatA11, passengerInA11)
	} else {
		t.Logf("Found passenger in seat A11: %s", passengerInA11.Name)
	}
}
