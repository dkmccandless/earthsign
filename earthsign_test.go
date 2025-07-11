package earthsign

import (
	"testing"
	"time"
)

func TestAt(t *testing.T) {
	for _, tt := range []struct {
		lat, lon float64
		t        time.Time
		name     string
	}{
		{0, 0, time.Date(1875, 3, 21, 0, 0, 0, 0, time.UTC), "Pisces"},
		{0, 0, time.Date(1875, 3, 21, 6, 0, 0, 0, time.UTC), "Orion"},
		{0, 0, time.Date(1875, 3, 21, 12, 0, 0, 0, time.UTC), "Virgo"},
		{0, 0, time.Date(1875, 3, 21, 18, 0, 0, 0, time.UTC), "Ophiuchus"},
		{0, 0, time.Date(1875, 3, 28, 18, 0, 0, 0, time.UTC), "Serpens"},
		{10, 0, time.Date(1875, 3, 28, 18, 0, 0, 0, time.UTC), "Scutum"},
		{-40, 0, time.Date(1875, 3, 28, 18, 0, 0, 0, time.UTC), "Lyra"},
		{-30, 30, time.Date(1875, 3, 21, 0, 0, 0, 0, time.UTC), "Triangulum"},
		{30, -30, time.Date(1875, 3, 21, 0, 0, 0, 0, time.UTC), "Piscis Austrinus"},
		{20, 0, time.Date(1875, 3, 21, 0, 0, 0, 0, time.UTC), "Cetus"},
		{20, 0, time.Date(1875, 6, 21, 0, 0, 0, 0, time.UTC), "Lepus"},
		{20, 0, time.Date(1875, 9, 21, 0, 0, 0, 0, time.UTC), "Corvus"},
		{20, 0, time.Date(1875, 12, 21, 0, 0, 0, 0, time.UTC), "Sagittarius"},
		{54.5, 0, time.Date(1725, 3, 21, 15, 10, 0, 0, time.UTC), "Circinus"},
		{54.5, 0, time.Date(1875, 3, 21, 15, 10, 0, 0, time.UTC), "Norma"},
		{54.5, 0, time.Date(2025, 3, 21, 15, 10, 0, 0, time.UTC), "Lupus"},
		{-90, 0, time.Date(2025, 6, 21, 2, 42, 11, 0, time.UTC), "Ursa Minor"},
		{90, 0, time.Date(2025, 6, 21, 2, 42, 11, 0, time.UTC), "Octans"},
	} {
		name, err := At(tt.lat, tt.lon, tt.t)
		if name != tt.name || err != nil {
			t.Errorf("At(%v, %v, %v): got %v, %v; want %v, %v",
				tt.lat, tt.lon, tt.t, name, err, tt.name, nil,
			)
		}
	}
}
