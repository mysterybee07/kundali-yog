package yoga

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func parseIntValue(v interface{}) (int, error) {
	switch value := v.(type) {
	case int:
		return value, nil
	case int32:
		return int(value), nil
	case int64:
		return int(value), nil
	case float64:
		return int(value), nil
	case json.Number:
		i, err := value.Int64()
		if err != nil {
			return 0, err
		}
		return int(i), nil
	case string:
		i, err := strconv.Atoi(strings.TrimSpace(value))
		if err != nil {
			return 0, err
		}
		return i, nil
	default:
		return 0, fmt.Errorf("unsupported int value %T", v)
	}
}

func parseFloatValue(v interface{}) (float64, error) {
	switch value := v.(type) {
	case float64:
		return value, nil
	case float32:
		return float64(value), nil
	case int:
		return float64(value), nil
	case int32:
		return float64(value), nil
	case int64:
		return float64(value), nil
	case json.Number:
		return value.Float64()
	case string:
		f, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
		if err != nil {
			return 0, err
		}
		return f, nil
	default:
		return 0, fmt.Errorf("unsupported float value %T", v)
	}
}

func parseBoolValue(v interface{}) (bool, error) {
	switch value := v.(type) {
	case bool:
		return value, nil
	case string:
		parsed, err := strconv.ParseBool(strings.TrimSpace(value))
		if err != nil {
			return false, err
		}
		return parsed, nil
	case float64:
		return value != 0, nil
	case int:
		return value != 0, nil
	}
	return false, fmt.Errorf("unsupported bool value %T", v)
}

func parseString(v interface{}) string {
	switch value := v.(type) {
	case string:
		return value
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case int:
		return strconv.Itoa(value)
	case json.Number:
		return value.String()
	default:
		return ""
	}
}

func normalizeName(value string) string {
	value = strings.TrimSpace(strings.ToLower(value))
	replacer := strings.NewReplacer("-", "", "_", "", " ", "")
	return replacer.Replace(value)
}

func normalizePower(value string) string {
	return normalizeName(value)
}

func validatePlacement(placement Placement) error {
	if placement.Planet == "" {
		return errors.New("planet is required")
	}

	if placement.House < 1 || placement.House > 12 {
		return fmt.Errorf("invalid house %d for %s", placement.House, placement.Planet)
	}

	if placement.Sign < 1 || placement.Sign > 12 {
		return fmt.Errorf("invalid sign %d for %s", placement.Sign, placement.Planet)
	}

	if placement.Degree < 0 || placement.Degree >= 30 {
		return fmt.Errorf("invalid degree %.2f for %s", placement.Degree, placement.Planet)
	}

	return nil
}

func ParseChartJSON(data []byte) (Chart, error) {
	var list []PlanetData
	if err := json.Unmarshal(data, &list); err == nil && len(list) > 0 {
		return NewChartFromPlanetData(list)
	}

	var payload ChartPayload
	if err := json.Unmarshal(data, &payload); err != nil {
		return Chart{}, err
	}

	if len(payload.Planets) == 0 {
		return Chart{}, errors.New("chart payload does not contain any planets")
	}

	return NewChartFromPlanetData(payload.Planets)
}

func ParsePlanet(value string) (Planet, error) {
	switch normalizeName(value) {
	case "sun":
		return Sun, nil
	case "moon":
		return Moon, nil
	case "mars":
		return Mars, nil
	case "mercury":
		return Mercury, nil
	case "jupiter":
		return Jupiter, nil
	case "venus":
		return Venus, nil
	case "saturn":
		return Saturn, nil
	case "rahu":
		return Rahu, nil
	case "ketu":
		return Ketu, nil
	case "asc", "ascendant", "lagna":
		return Asc, nil
	default:
		return "", fmt.Errorf("unknown planet %q", value)
	}
}

func (p *PlanetData) UnmarshalJSON(data []byte) error {
	type rawPlanetData struct {
		PlanetName string      `json:"PlanetEngName"`
		Sign       interface{} `json:"Sign"`
		House      interface{} `json:"House"`
		Degree     interface{} `json:"Degree"`
		Retro      interface{} `json:"Retro"`
		Combust    interface{} `json:"Combust"`
		Power      string      `json:"Power"`
		Nakshatra  Nakshatra   `json:"Nakshatra"`
	}

	var raw rawPlanetData
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	sign, err := parseIntValue(raw.Sign)
	if err != nil {
		return fmt.Errorf("invalid sign for %q: %w", raw.PlanetName, err)
	}

	house, err := parseIntValue(raw.House)
	if err != nil {
		return fmt.Errorf("invalid house for %q: %w", raw.PlanetName, err)
	}

	degree, err := parseFloatValue(raw.Degree)
	if err != nil {
		return fmt.Errorf("invalid degree for %q: %w", raw.PlanetName, err)
	}

	retro := false
	if raw.Retro != nil {
		retro, err = parseBoolValue(raw.Retro)
		if err != nil {
			return fmt.Errorf("invalid retro flag for %q: %w", raw.PlanetName, err)
		}
	}

	combust := false
	if raw.Combust != nil {
		combust, err = parseBoolValue(raw.Combust)
		if err != nil {
			return fmt.Errorf("invalid combust flag for %q: %w", raw.PlanetName, err)
		}
	}

	*p = PlanetData{
		PlanetName: raw.PlanetName,
		Sign:       sign,
		House:      house,
		Degree:     degree,
		Retro:      retro,
		Combust:    combust,
		Power:      raw.Power,
		Nakshatra:  raw.Nakshatra,
	}

	return nil
}
