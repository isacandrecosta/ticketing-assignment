package controllers

import (
	"encoding/json"
	"net/http"
	"ticketing-assignment/internal/models"
	"ticketing-assignment/internal/services"
)

type ReservationController struct {
	Service *services.ReservationService
}

func (rc *ReservationController) IsSeatReservedHandler(w http.ResponseWriter, r *http.Request) {
	serviceID := r.URL.Query().Get("serviceID")
	date := r.URL.Query().Get("date")
	seatNumber := r.URL.Query().Get("seatNumber")

	isReserved := rc.Service.IsSeatReserved(serviceID, date, seatNumber)
	json.NewEncoder(w).Encode(isReserved)
}

func (rc *ReservationController) CountPassengersAtStationHandler(w http.ResponseWriter, r *http.Request) {
	stationName := r.URL.Query().Get("stationName")
	boarding := r.URL.Query().Get("boarding") == "true"

	count := rc.Service.CountPassengersAtStation(stationName, boarding)
	json.NewEncoder(w).Encode(count)
}

func (rc *ReservationController) FindPassengerInSeatHandler(w http.ResponseWriter, r *http.Request) {
	serviceID := r.URL.Query().Get("serviceID")
	date := r.URL.Query().Get("date")
	carriage := r.URL.Query().Get("carriage")
	seatNumber := r.URL.Query().Get("seatNumber")

	passenger := rc.Service.FindPassengerInSeat(serviceID, date, carriage, seatNumber)
	json.NewEncoder(w).Encode(passenger)
}

func (rc *ReservationController) AddBookingHandler(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	json.NewDecoder(r.Body).Decode(&booking)

	rc.Service.AddBooking(booking)
	w.WriteHeader(http.StatusCreated)
}
