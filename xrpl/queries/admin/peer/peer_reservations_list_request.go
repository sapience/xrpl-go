package peer

type ReservationsListRequest struct {
}

func (*ReservationsListRequest) Method() string {
	return "peer_reservations_list"
}
