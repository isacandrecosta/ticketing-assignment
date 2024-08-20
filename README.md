# Ticketing Assignment

## Description
This project is a ticket reservation system that allows users to create, view, and manage ticket reservations.

## Features

- **Create ticket reservations**: Users can create new reservations for different services.
- **Check if a seat is taken**: Verify if a specific seat is already reserved for a given service and date.
- **Count passengers at a station**: Count the number of passengers boarding or leaving at a specific station.
- **Find passenger in a seat**: Find the passenger occupying a specific seat in a given service and date.
- **Add new bookings**: Add new bookings to the reservation system.


## Project Structure
- `cmd/`: Contains the main entry point of the application.
- `internal/controllers/`: Contains the controllers that handle HTTP requests.
- `internal/services/`: Contains business logic and services.
- `internal/models/`: Contains data model definitions.
- `internal/views/`: Contains standardized response definitions.
- `internal/errors/`: Contains internal error definitions.
- `tests/`: Contains unit and integration tests.
