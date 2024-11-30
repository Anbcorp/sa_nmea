package nmea

import (
	"fmt"
)

// HDT - Heading True - Actual vessel heading in degrees true
type HDT struct {
	nmea_data
	heading float64
}

func (hdt *HDT) FromBoat(b NMEABoat) {
	hdt.id = "II"
	hdt.stype = "HDT"
	hdt.heading = b.Hdg
}

func (hdt *HDT) Values() string {
	// $IIHDT,265,T*3D
	return fmt.Sprintf("%d,T", int(hdt.heading))
}
