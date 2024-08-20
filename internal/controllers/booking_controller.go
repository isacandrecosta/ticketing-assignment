package controllers

import (
	"encoding/json"
	"net/http"
	"ticketing-assignment/internal/services"
)

type BookingController struct {
	ReservationService *services.ReservationService
}

func NewBookingController(rs *services.ReservationService) *BookingController {
	return &BookingController{ReservationService: rs}
}

func (bc *BookingController) CreateBooking(w http.ResponseWriter, r *http.Request) {
	response := bc.ReservationService.HandleCreateBooking(r)
	w.WriteHeader(response.GetStatusCode())
	json.NewEncoder(w).Encode(response.GetBody())
}
