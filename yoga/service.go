package yoga

import (
	"fmt"
)

func (y Yoga) Evaluate(chart Chart) bool {
	for _, cond := range y.Conditions {
		if cond == nil || !cond(chart) {
			return false
		}
	}
	return true
}

func (y Yoga) Result(chart Chart) (YogaResult, bool) {
	if !y.Evaluate(chart) {
		return YogaResult{}, false
	}

	result := YogaResult{
		Name:     y.Name,
		Group:    y.Group,
		Strength: YogaStrengthStrong,
	}

	for _, cond := range y.WeakIf {
		if cond != nil && cond(chart) {
			result.Strength = YogaStrengthWeak
			break
		}
	}

	return result, true
}

func EvaluateAll(chart Chart, yogas []Yoga) []YogaResult {
	results := make([]YogaResult, 0, len(yogas))

	for _, y := range yogas {
		if result, ok := y.Result(chart); ok {
			results = append(results, result)
		}
	}

	return results
}

func NewChart(placements ...Placement) (Chart, error) {
	chart := Chart{
		Planets: make(map[Planet]Placement, len(placements)),
	}

	for _, placement := range placements {
		if err := validatePlacement(placement); err != nil {
			return Chart{}, err
		}

		if _, exists := chart.Planets[placement.Planet]; exists {
			return Chart{}, fmt.Errorf("duplicate placement for %s", placement.Planet)
		}

		chart.Planets[placement.Planet] = placement
	}

	return chart, nil
}

func NewChartFromPlanetData(data []PlanetData) (Chart, error) {
	placements := make([]Placement, 0, len(data))

	for _, item := range data {
		planet, err := ParsePlanet(item.PlanetName)
		if err != nil {
			return Chart{}, err
		}

		placements = append(placements, Placement{
			Planet:  planet,
			House:   item.House,
			Sign:    item.Sign,
			Degree:  item.Degree,
			Retro:   item.Retro,
			Combust: item.Combust,
			Power:   item.Power,
		})
	}

	return NewChart(placements...)
}

func GetPlanetHouse(chart Chart, planet Planet) (int, bool) {
	placement, ok := chart.Planets[planet]
	if !ok {
		return 0, false
	}

	return placement.House, true
}

func GetPlanetSign(chart Chart, planet Planet) (int, bool) {
	placement, ok := chart.Planets[planet]
	if !ok {
		return 0, false
	}

	return placement.Sign, true
}

func GetAscSign(chart Chart) (int, bool) {
	return GetPlanetSign(chart, Asc)
}

func GetHouseSign(ascSign, house int) int {
	return ((ascSign + house - 2) % 12) + 1
}

func GetHouseLord(chart Chart, house int) (Planet, bool) {
	ascSign, ok := GetAscSign(chart)
	if !ok {
		return "", false
	}

	sign := GetHouseSign(ascSign, house)
	lord, ok := signLords[sign]
	if !ok {
		return "", false
	}

	return lord, true
}
