package yoga

import "math"

func PlanetExists(p Planet) Condition {
	return func(chart Chart) bool {
		_, ok := chart.Planets[p]
		return ok
	}
}

func AreConjunct(config ConjunctionConfig, planets ...Planet) Condition {
	return func(chart Chart) bool {

		switch config.Type {

		case HouseBased:
			return houseBased(chart, planets)

		case DegreeBased:
			return degreeBased(chart, planets, config.Orb)

		case StrengthBased:
			return strengthBased(chart, planets, config.Orb)

		default:
			return false
		}
	}
}

func IsKendraHouseFrom(referenceHouse, targetHouse int) bool {
	diff := (targetHouse - referenceHouse + 12) % 12
	return diff == 0 || diff == 3 || diff == 6 || diff == 9
}

func IsKendraFrom(targetPlanet, referencePlanet Planet) Condition {
	return func(chart Chart) bool {
		referenceHouse, ok := GetPlanetHouse(chart, referencePlanet)
		if !ok {
			return false
		}

		targetHouse, ok := GetPlanetHouse(chart, targetPlanet)
		if !ok {
			return false
		}

		return IsKendraHouseFrom(referenceHouse, targetHouse)
	}
}

func IsInHouseFrom(targetPlanet, referencePlanet Planet, offset int) Condition {
	return func(chart Chart) bool {
		referenceHouse, ok := GetPlanetHouse(chart, referencePlanet)
		if !ok {
			return false
		}

		targetHouse, ok := GetPlanetHouse(chart, targetPlanet)
		if !ok {
			return false
		}

		expectedHouse := ((referenceHouse + offset - 2) % 12) + 1
		return targetHouse == expectedHouse
	}
}

func AnyPlanetInHouseFrom(referencePlanet Planet, offset int, planets ...Planet) Condition {
	return func(chart Chart) bool {
		for _, planet := range planets {
			if IsInHouseFrom(planet, referencePlanet, offset)(chart) {
				return true
			}
		}
		return false
	}
}

func NoPlanetInHouseFrom(referencePlanet Planet, offset int, planets ...Planet) Condition {
	return func(chart Chart) bool {
		for _, planet := range planets {
			if IsInHouseFrom(planet, referencePlanet, offset)(chart) {
				return false
			}
		}
		return true
	}
}

func IsNotCombust(p Planet) Condition {
	return func(chart Chart) bool {
		pl, ok := chart.Planets[p]
		return ok && !pl.Combust
	}
}

func IsCombust(p Planet) Condition {
	return func(chart Chart) bool {
		pl, ok := chart.Planets[p]
		return ok && pl.Combust
	}
}

func IsInKendraFromAsc(p Planet) Condition {
	return func(chart Chart) bool {
		placement, ok := chart.Planets[p]
		if !ok {
			return false
		}

		return IsKendraHouseFrom(1, placement.House)
	}
}

func IsPanchaMahapurushPlanet(p Planet) Condition {
	return func(chart Chart) bool {
		if !PlanetExists(p)(chart) {
			return false
		}

		if !IsInKendraFromAsc(p)(chart) {
			return false
		}

		return AnyCondition(IsOwnSign(p), IsExalted(p))(chart)
	}
}

func IsRajYoga() Condition {
	return func(chart Chart) bool {
		kendraLords := make([]Planet, 0, 4)
		for _, house := range []int{1, 4, 7, 10} {
			lord, ok := GetHouseLord(chart, house)
			if ok {
				kendraLords = append(kendraLords, lord)
			}
		}

		trikonaLords := make([]Planet, 0, 3)
		for _, house := range []int{1, 5, 9} {
			lord, ok := GetHouseLord(chart, house)
			if ok {
				trikonaLords = append(trikonaLords, lord)
			}
		}

		for _, kendraLord := range kendraLords {
			for _, trikonaLord := range trikonaLords {
				if kendraLord == trikonaLord {
					return true
				}

				if AreConjunct(ConjunctionConfig{Type: HouseBased}, kendraLord, trikonaLord)(chart) {
					return true
				}
			}
		}

		return false
	}
}

func DegreeDifference(d1, d2 float64) float64 {
	diff := math.Abs(d1 - d2)
	if diff > 15 {
		diff = 30 - diff
	}
	return diff
}

func degreeBased(chart Chart, planets []Planet, orb float64) bool {
	if len(planets) < 2 {
		return false
	}

	for i := 0; i < len(planets); i++ {
		p1 := chart.Planets[planets[i]]

		for j := i + 1; j < len(planets); j++ {
			p2 := chart.Planets[planets[j]]

			if p1.House != p2.House {
				return false
			}

			if DegreeDifference(p1.Degree, p2.Degree) > orb {
				return false
			}
		}
	}
	return true
}

func strengthBased(chart Chart, planets []Planet, orb float64) bool {
	if len(planets) < 2 {
		return false
	}

	for i := 0; i < len(planets); i++ {
		p1 := chart.Planets[planets[i]]

		if !isStrong(p1) {
			return false
		}

		for j := i + 1; j < len(planets); j++ {
			p2 := chart.Planets[planets[j]]

			if !isStrong(p2) {
				return false
			}

			if p1.House != p2.House {
				return false
			}

			if DegreeDifference(p1.Degree, p2.Degree) > orb {
				return false
			}
		}
	}

	return true
}

func houseBased(chart Chart, planets []Planet) bool {
	if len(planets) < 2 {
		return false
	}

	base := chart.Planets[planets[0]]

	for _, p := range planets[1:] {
		if chart.Planets[p].House != base.House {
			return false
		}
	}
	return true
}
