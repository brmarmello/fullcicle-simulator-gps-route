package route

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
