package route

import "errors"

// Route represents a request of new delivery request
type Route struct {
	ID        string
	ClientID  string
	Positions []Position
}

// Position is a type which contains the lat and long
type Position struct {
	Lat  float64
	Long float64
}

// LoadPositions loads from a .txt file all positions (lat and long) to the Position attribute of the struct
func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("route id not informed")
	}
}