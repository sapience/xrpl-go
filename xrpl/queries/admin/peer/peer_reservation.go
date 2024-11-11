package peer

type Reservation struct {
	Node        string `json:"node"`
	Description string `json:"description,omitempty"`
}
