package nmea

import (
	"fmt"
	"log"
	"math"
	"strings"
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
	m := mins * 60
	mm, dec := math.Modf(m)
	//strpos := fmt.Sprintf("%04d.%04d", int(dd*100+mm), int(dec*10000))
	strdec := strings.Split(fmt.Sprintf("%.4f", dec), ".")[1]
	strpos := fmt.Sprintf("%04d.%s", int(dd*100+mm), strdec)
	log.Printf("Converting coord %.6f to %f %f(%f %f) -> %s", pos, dd, m, mm, dec, strpos)
	return strpos
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
