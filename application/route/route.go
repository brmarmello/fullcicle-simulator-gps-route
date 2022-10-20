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

// PartialRoutePosition is the actual response which the system will return
type PartialRoutePosition struct {
	ID       string
	ClientID string
	Position []float64
	Finished bool
}

// LoadPositions loads from a .txt file all positions (lat and long) to the Position attribute of the struct
func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("route id not informed")
	}
	f, err := os.Open("destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}
	// PS.: In "go" the "defer" resource is used to run everything that is needed, leaving "f.Close() to run last!
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return nil
		}
		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return nil
		}
		r.Positions = append(r.Positions, Position{
			Lat:  lat,
			Long: long,
		})
	}
	return nil

}

func (r *Route) ExportJsonPositions() ([]string, error) {
	
}