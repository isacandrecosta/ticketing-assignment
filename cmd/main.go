package main

import (
	"fmt"
	"net/http"
	"ticketing-assignment/internal/controllers"
	"ticketing-assignment/internal/mockdata"
	"ticketing-assignment/internal/models"
	"ticketing-assignment/internal/services"
)

func main() {
	rs := &models.ReservationSystem{}
	reservationService := &services.ReservationService{ReservationSystem: rs}

	booking := mockdata.CreateAmsterdamToBerlinBooking()

	err := reservationService.MakeReservation(booking)
	if err != nil {
		fmt.Println("Error making reservation:", err)
	} else {
		fmt.Println("Reservation made successfully!")
	}

	bookingController := controllers.NewBookingController(reservationService)

	http.HandleFunc("/create-booking", bookingController.CreateBooking)
	http.ListenAndServe(":8080", nil)
}
