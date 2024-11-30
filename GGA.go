package nmea

import (
	"fmt"
)

// GGA -  Geographic Position - Latitude/Longitude
type GGA struct {
	GLL
	deadreck bool
}

func (gga *GGA) FromBoat(b NMEABoat) {
	gga.GLL.FromBoat(b)
	gga.stype = "GGA"
	gga.deadreck = false
}

func (gga *GGA) GpsQual() int {
	if gga.deadreck {
		return 6
	} else {
		return 1
	}
}

func (gga *GGA) Values() string {
	// $GPGGA,222701.519,4630.5092,N,0934.7077,W,6,4,0,0,M,,,,*11
	return fmt.Sprintf("%s,%s,%c,%s,%c,%d,4,0,0,M,,,,", NMEATimestamp(gga.utc), NMEALatLon(gga.latitude), gga.lat_sign, NMEALatLon(gga.longitude), gga.lon_sign, gga.GpsQual())
}
