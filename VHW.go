package nmea

import (
	"fmt"
)

// VHW is composed by a nmea_sentence, plus heading and stw
type VHW struct {
	nmea_data
	heading float64 // True and Mag heading are the same
	stw     float64 // TODO: which unit do we use internally ?
}

func (vhw *VHW) FromBoat(b NMEABoat) {
	vhw.id = "II"
	vhw.stype = "VHW"
	vhw.heading = b.Hdg
	vhw.stw = b.Spd
}

func (vhw *VHW) Values() string {
	// $IIVHW,279,T,279,M,8.2,N,15,K*75
	return fmt.Sprintf("%d,T,%d,M,%.1f,N,%.1f,K", int(vhw.heading), int(vhw.heading), vhw.stw, Knt2Kmh(vhw.stw))
}
