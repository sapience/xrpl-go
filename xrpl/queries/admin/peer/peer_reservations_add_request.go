package peer

type ReservationAddRequest struct {
	PublicKey   string `json:"public_key"`
	Description string `json:"description,omitempty"`
}

func (*ReservationAddRequest) Method() string {
	return "peer_reservations_add"
}
