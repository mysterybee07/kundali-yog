package kundali

import "github.com/mysterybee07/kundali-yog/yoga"

type (
	ConjunctionType   = yoga.ConjunctionType
	Condition         = yoga.Condition
	PlanetDignity     = yoga.PlanetDignity
	PlanetData        = yoga.PlanetData
	Nakshatra         = yoga.Nakshatra
	Planet            = yoga.Planet
	Placement         = yoga.Placement
	Chart             = yoga.Chart
	ChartPayload      = yoga.ChartPayload
	YogaStrength      = yoga.YogaStrength
	Yoga              = yoga.Yoga
	YogaResult        = yoga.YogaResult
	ConjunctionConfig = yoga.ConjunctionConfig
)

const (
	HouseBased    = yoga.HouseBased
	DegreeBased   = yoga.DegreeBased
	StrengthBased = yoga.StrengthBased

	Sun     = yoga.Sun
	Moon    = yoga.Moon
	Mars    = yoga.Mars
	Mercury = yoga.Mercury
	Jupiter = yoga.Jupiter
	Venus   = yoga.Venus
	Saturn  = yoga.Saturn
	Rahu    = yoga.Rahu
	Ketu    = yoga.Ketu
	Asc     = yoga.Asc

	YogaStrengthStrong = yoga.YogaStrengthStrong
	YogaStrengthWeak   = yoga.YogaStrengthWeak
)

func EvaluateAll(chart Chart, yogas []Yoga) []YogaResult {
	return yoga.EvaluateAll(chart, yogas)
}

func DefaultYogas() []Yoga {
	return yoga.DefaultYogas()
}

func GetYogas(chart Chart) []YogaResult {
	return yoga.EvaluateAll(chart, yoga.DefaultYogas())
}

func NewChart(placements ...Placement) (Chart, error) {
	return yoga.NewChart(placements...)
}

func NewChartFromPlanetData(data []PlanetData) (Chart, error) {
	return yoga.NewChartFromPlanetData(data)
}

func ParseChartJSON(data []byte) (Chart, error) {
	return yoga.ParseChartJSON(data)
}

func ParsePlanet(value string) (Planet, error) {
	return yoga.ParsePlanet(value)
}

func GetPlanetHouse(chart Chart, planet Planet) (int, bool) {
	return yoga.GetPlanetHouse(chart, planet)
}

func GetPlanetSign(chart Chart, planet Planet) (int, bool) {
	return yoga.GetPlanetSign(chart, planet)
}

func GetAscSign(chart Chart) (int, bool) {
	return yoga.GetAscSign(chart)
}

func GetHouseSign(ascSign, house int) int {
	return yoga.GetHouseSign(ascSign, house)
}

func GetHouseLord(chart Chart, house int) (Planet, bool) {
	return yoga.GetHouseLord(chart, house)
}

func AnyCondition(conditions ...Condition) Condition {
	return yoga.AnyCondition(conditions...)
}

func AllConditions(conditions ...Condition) Condition {
	return yoga.AllConditions(conditions...)
}

func PlanetExists(p Planet) Condition {
	return yoga.PlanetExists(p)
}

func AreConjunct(config ConjunctionConfig, planets ...Planet) Condition {
	return yoga.AreConjunct(config, planets...)
}

func IsKendraHouseFrom(referenceHouse, targetHouse int) bool {
	return yoga.IsKendraHouseFrom(referenceHouse, targetHouse)
}

func IsKendraFrom(targetPlanet, referencePlanet Planet) Condition {
	return yoga.IsKendraFrom(targetPlanet, referencePlanet)
}

func IsInHouseFrom(targetPlanet, referencePlanet Planet, offset int) Condition {
	return yoga.IsInHouseFrom(targetPlanet, referencePlanet, offset)
}

func AnyPlanetInHouseFrom(referencePlanet Planet, offset int, planets ...Planet) Condition {
	return yoga.AnyPlanetInHouseFrom(referencePlanet, offset, planets...)
}

func NoPlanetInHouseFrom(referencePlanet Planet, offset int, planets ...Planet) Condition {
	return yoga.NoPlanetInHouseFrom(referencePlanet, offset, planets...)
}

func IsNotCombust(p Planet) Condition {
	return yoga.IsNotCombust(p)
}

func IsCombust(p Planet) Condition {
	return yoga.IsCombust(p)
}

func IsInKendraFromAsc(p Planet) Condition {
	return yoga.IsInKendraFromAsc(p)
}

func IsPanchaMahapurushPlanet(p Planet) Condition {
	return yoga.IsPanchaMahapurushPlanet(p)
}

func IsRajYoga() Condition {
	return yoga.IsRajYoga()
}

func DegreeDifference(d1, d2 float64) float64 {
	return yoga.DegreeDifference(d1, d2)
}

func IsOwnSign(p Planet) Condition {
	return yoga.IsOwnSign(p)
}

func IsExalted(p Planet) Condition {
	return yoga.IsExalted(p)
}

func IsDebilitated(p Planet) Condition {
	return yoga.IsDebilitated(p)
}
