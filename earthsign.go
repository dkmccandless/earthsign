// Package earthsign identifies the constellation directly below a point
// on Earth's surface at a particular time; that is, the constellation that
// (the center of) Earth is "in" at that location and at that moment.
// This information is presumably relevant for astrological purposes.
package earthsign

import (
	"time"

	"github.com/dkmccandless/constellation"
	"github.com/soniakeys/meeus/v3/base"
	"github.com/soniakeys/meeus/v3/coord"
	"github.com/soniakeys/meeus/v3/globe"
	"github.com/soniakeys/meeus/v3/precess"
	"github.com/soniakeys/meeus/v3/sidereal"
	"github.com/soniakeys/unit"
)

// unixEpochJDE is the Julian date of the Unix epoch.
const unixEpochJDE = 2440587.5

// b1875 is the Julian year value of the Besselian year 1875.0 epoch.
var b1875 = base.JDEToJulianYear(base.BesselianYearToJDE(1875.0))

// nadir is the celestial position directly below an observer.
var nadir = coord.Horizontal{Alt: unit.AngleFromDeg(-90)}

// At returns the constellation at the nadir of an observer
// at the given latitude and longitude, in degrees, at time t.
// North and east are positive.
func At(lat, lon float64, t time.Time) (string, error) {
	jde := julianDate(t)

	// Type Coord's longitude is measured positively westward from the Greenwich meridian.
	co := globe.Coord{Lat: unit.AngleFromDeg(lat), Lon: -unit.AngleFromDeg(lon)}

	eq := new(coord.Equatorial).HzToEq(&nadir, co, sidereal.Mean(jde))
	precess.NewPrecessor(base.JDEToJulianYear(jde), b1875).Precess(eq, eq)

	return constellation.At(eq.RA.Hour(), eq.Dec.Deg())
}

func julianDate(t time.Time) float64 {
	return unixEpochJDE + float64(t.Sub(time.Unix(0, 0)))/float64(24*time.Hour)
}
