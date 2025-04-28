package host

// *********** Entities
// Host connected to the network
type Host struct {
	Name        string
	ASname      string `json:"ASname,omitempty"`
	PrivateHost bool
	IP          string
	Mac         string
	City        string
	Country     string
}
