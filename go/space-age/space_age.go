package space

const earthYear = 31557600.0

type Planet string

var orbitalPeriods = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}


// Age calculates Earth age on various palents
func Age(ageInSeconds float64, planet Planet) float64 {
	yearsOnPlanet := earthYear * orbitalPeriods[planet]
	years := ageInSeconds / yearsOnPlanet
	return years
}
