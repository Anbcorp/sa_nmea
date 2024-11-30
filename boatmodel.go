package nmea

import "time"

type NMEABoat struct {
	Timestamp time.Time
	Latitude  float64
	Longitude float64
	Cog       float64
	Sog       float64
	Spd       float64
	Hdg       float64
	Awa       float64
	Aws       float64
	Twa       float64
	Tws       float64
}
