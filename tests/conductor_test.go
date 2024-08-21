package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"ticketing-assignment/internal/controllers"
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

func TestIsSeatReservedHandler(t *testing.T) {
	reservationSystem := mockdata.CreateMockReservationSystem()
	reservationService := &services.ReservationService{ReservationSystem: reservationSystem}
	reservationController := &controllers.ReservationController{Service: reservationService}

	req, err := http.NewRequest("GET", "/isSeatReserved?serviceID=5160&date=2021-04-01&seatNumber=A11", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(reservationController.IsSeatReservedHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := true
	var actual bool
	if err := json.NewDecoder(rr.Body).Decode(&actual); err != nil {
		t.Errorf("could not decode response: %v", err)
	}
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestIsSeatNotReservedHandler(t *testing.T) {
	reservationSystem := mockdata.CreateMockReservationSystem()
	reservationService := &services.ReservationService{ReservationSystem: reservationSystem}
	reservationController := &controllers.ReservationController{Service: reservationService}

	req, err := http.NewRequest("GET", "/isSeatReserved?serviceID=5160&date=2021-04-01&seatNumber=B3", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(reservationController.IsSeatReservedHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := false
	var actual bool
	if err := json.NewDecoder(rr.Body).Decode(&actual); err != nil {
		t.Errorf("could not decode response: %v", err)
	}
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
