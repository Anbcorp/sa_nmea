package nmea

import (
	"fmt"
)

// RMC - Recommended Minimum Navigation Information
type RMC struct {
	GLL
	Sog float64
	Cog float64
}

func (rmc *RMC) FromBoat(b NMEABoat) {
	rmc.GLL.FromBoat(b)
	rmc.stype = "RMC"
	rmc.Sog = b.Sog
	rmc.Cog = b.Cog
}

func (rmc *RMC) Values() string {
	// $GPRMC,222701.519,A,4630.5092,N,0934.7077,W,9.6,266,311023,,,*13
	return fmt.Sprintf("%s,%s,%c,%s,%c,%.1f,%d,%s,,,",
		NMEATimestamp(rmc.utc),
		NMEALatLon(rmc.latitude),
		rmc.lat_sign,
		NMEALatLon(rmc.longitude),
		rmc.lon_sign,
		rmc.Sog,
		int(rmc.Cog),
		NMEADate(rmc.utc))
}
