package nmea

import (
	"fmt"
	"math"
	"time"
)

// Format a time.Time to NMEA GLL time
// HHmmss.ss
func NMEATimestamp(timestamp time.Time) string {
	// Golang way of formatting date is probably the worst I've ever seen, ugh...
	return timestamp.Format("150405.000")
}

// Format a time.Time to NMEA RMC date
// DDMMYY
func NMEADate(timestamp time.Time) string {
	return timestamp.Format("020106")
}

// Format a position in DDDMM.mmmmm format
func NMEALatLon(pos float64) string {
	dd, mins := math.Modf(math.Abs(pos))
	mm := mins * 60
	return fmt.Sprintf("%02d%02.04f", int(dd), mm)
}

// Convert an relative wind direction (TWA,AWA) (+/- 0->180) to an absolute wind angle (0->359)
func Wind2Angle(dir float64) float64 {
	if dir > 0 {
		return dir
	} else {
		return 360 + dir
	}
}

// knots to km/h
func Knt2Kmh(s float64) float64 {
	return s * 1.852
}
