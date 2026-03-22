package yoga

func IsOwnSign(p Planet) Condition {
	return func(chart Chart) bool {
		placement, ok := chart.Planets[p]
		return ok && isOwnSign(p, placement.Sign, placement.Power)
	}
}

func IsExalted(p Planet) Condition {
	return func(chart Chart) bool {
		placement, ok := chart.Planets[p]
		return ok && isExalted(p, placement.Sign, placement.Power)
	}
}

func IsDebilitated(p Planet) Condition {
	return func(chart Chart) bool {
		placement, ok := chart.Planets[p]
		return ok && isDebilitated(p, placement.Sign, placement.Power)
	}
}

func isStrong(p Placement) bool {
	if p.Combust {
		return false
	}
	return true
}

func isOwnOrExalted(planet Planet, sign int, power string) bool {
	return isOwnSign(planet, sign, power) || isExalted(planet, sign, power)
}

func isOwnSign(planet Planet, sign int, power string) bool {
	if power == "OwnSign" {
		return true
	}

	dignity, ok := planetDignities[planet]
	if !ok {
		return false
	}

	for _, ownSign := range dignity.OwnSigns {
		if ownSign == sign {
			return true
		}
	}

	return false
}

func isExalted(planet Planet, sign int, power string) bool {
	if power == "Exalted" {
		return true
	}

	dignity, ok := planetDignities[planet]
	if !ok {
		return false
	}

	return dignity.ExaltationSign == sign
}

func isDebilitated(planet Planet, sign int, power string) bool {
	if power == "Debilitated" {
		return true
	}

	dignity, ok := planetDignities[planet]
	if !ok {
		return false
	}

	return dignity.Debilitation == sign
}
