package nmea

import (
	"fmt"
	"time"
)

// GLL -  Geographic Position - Latitude/Longitude
type GLL struct {
	nmea_data
	latitude  float64
	lat_sign  byte
	longitude float64
	lon_sign  byte
	utc       time.Time
}

func (gll *GLL) FromBoat(b NMEABoat) {
	gll.id = "GP"
	gll.stype = "GLL"

	// Latitude : positive is N, negative is S
	gll.latitude = b.Latitude
	if b.Latitude < 0 {
		gll.lat_sign = 'S'
	} else {
		gll.lat_sign = 'N'
	}

	// Longitude : positive is E, negative is W
	gll.longitude = b.Longitude
	if b.Longitude < 0 {
		gll.lon_sign = 'W'
	} else {
		gll.lon_sign = 'E'
	}

	gll.utc = b.Timestamp
}

func (gll *GLL) Values() string {
	// $GPGLL,4630.5092,N,0934.7077,W,222701.519,A*19
	return fmt.Sprintf("%s,%c,%s,%c,%s,A", NMEALatLon(gll.latitude), gll.lat_sign, NMEALatLon(gll.longitude), gll.lon_sign, NMEATimestamp(gll.utc))
}
