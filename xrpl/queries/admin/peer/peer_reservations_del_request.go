package peer

type ReservationDelRequest struct {
	PublicKey string `json:"public_key"`
}

func (*ReservationDelRequest) Method() string {
	return "peer_reservations_del"
}
