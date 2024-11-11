package peer

type ReservationsListResponse struct {
	Reservations []*Reservation `json:"reservations"`
}
