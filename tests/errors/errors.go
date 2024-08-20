package errors

const (
	ErrPassengerNotFound         = "Expected passenger %s in seat %s, got %v"
	ErrPassengersBoarding        = "Expected %d passengers boarding at %s, got %d"
	ErrPassengersLeaving         = "Expected %d passengers leaving at %s, got %d"
	ErrPassengersBetweenStations = "Expected %d passengers between %s and %s, got %d"
	ErrFailedBooking             = "Failed to make booking: %v"
	ErrDuplicateBooking          = "Expected error for duplicate booking, got nil"
	ErrNumberOfStops             = "Expected %d stops, got %d"
	ErrTotalDistance             = "Expected total distance %d, got %d"
	ErrSeatAlreadyReserved       = "seat already reserved"
	ErrInvalidRequestPayload     = "Invalid request payload"
)
