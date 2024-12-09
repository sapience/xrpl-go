package peer

type Reservation struct {
	Node        string `json:"node"`
	Description string `json:"description,omitempty"`
}

// ############################################################
// Add a peer reservation
// ############################################################

type ReservationAddRequest struct {
	PublicKey   string `json:"public_key"`
	Description string `json:"description,omitempty"`
}

func (*ReservationAddRequest) Method() string {
	return "peer_reservations_add"
}

type ReservationsAddResponse struct {
	Previous *Reservation `json:"previous,omitempty"`
}

// ############################################################
// Delete a peer reservation
// ############################################################

type ReservationDelRequest struct {
	PublicKey string `json:"public_key"`
}

func (*ReservationDelRequest) Method() string {
	return "peer_reservations_del"
}

type ReservationsDelResponse struct {
	Previous *Reservation `json:"previous,omitempty"`
}

// ############################################################
// List peer reservations
// ############################################################

type ReservationsListRequest struct {
}

func (*ReservationsListRequest) Method() string {
	return "peer_reservations_list"
}

type ReservationsListResponse struct {
	Reservations []*Reservation `json:"reservations"`
}
