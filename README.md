# kundali-yog

`kundali-yog` is a Go library for evaluating yogas from a kundali chart.

After installing the module, consumers can import the root package and call:

```go
kundali.GetYogas(chart)
```

## Install

```bash
go get github.com/mysterybee07/kundali-yog
```

## Import

```go
import kundali "github.com/mysterybee07/kundali-yog"
```

## Quick Start

```go
package main

import (
	"fmt"
	"log"

	kundali "github.com/mysterybee07/kundali-yog"
)

func main() {
	chart, err := kundali.NewChart(
		kundali.Placement{Planet: kundali.Asc, House: 1, Sign: 1, Degree: 0},
		kundali.Placement{Planet: kundali.Mars, House: 1, Sign: 1, Degree: 10},
		kundali.Placement{Planet: kundali.Mercury, House: 2, Sign: 2, Degree: 12},
		kundali.Placement{Planet: kundali.Jupiter, House: 4, Sign: 4, Degree: 8},
		kundali.Placement{Planet: kundali.Venus, House: 7, Sign: 7, Degree: 15},
		kundali.Placement{Planet: kundali.Saturn, House: 10, Sign: 10, Degree: 20},
		kundali.Placement{Planet: kundali.Sun, House: 5, Sign: 5, Degree: 5},
		kundali.Placement{Planet: kundali.Moon, House: 3, Sign: 3, Degree: 18},
		kundali.Placement{Planet: kundali.Rahu, House: 11, Sign: 11, Degree: 7},
		kundali.Placement{Planet: kundali.Ketu, House: 5, Sign: 5, Degree: 7},
	)
	if err != nil {
		log.Fatal(err)
	}

	results := kundali.GetYogas(chart)

	for _, yoga := range results {
		fmt.Printf("%s (%s) - %s\n", yoga.Name, yoga.Group, yoga.Strength)
	}
}
```

## Build a Chart From Planet Data

Use `NewChartFromPlanetData` if your input already matches the package payload structure.

```go
data := []kundali.PlanetData{
	{PlanetName: "Asc", Sign: 1, House: 1, Degree: 0},
	{PlanetName: "Mars", Sign: 1, House: 1, Degree: 10},
	{PlanetName: "Jupiter", Sign: 4, House: 4, Degree: 8},
}

chart, err := kundali.NewChartFromPlanetData(data)
if err != nil {
	log.Fatal(err)
}

results := kundali.GetYogas(chart)
```

## Build a Chart From JSON

Use `ParseChartJSON` when you receive raw JSON bytes.

```go
jsonData := []byte(`[
  {"PlanetEngName":"Asc","Sign":1,"House":1,"Degree":0},
  {"PlanetEngName":"Mars","Sign":1,"House":1,"Degree":10},
  {"PlanetEngName":"Jupiter","Sign":4,"House":4,"Degree":8}
]`)

chart, err := kundali.ParseChartJSON(jsonData)
if err != nil {
	log.Fatal(err)
}

results := kundali.GetYogas(chart)
```

The parser also supports this wrapped payload shape:

```json
{
  "planets": [
    {"PlanetEngName":"Asc","Sign":1,"House":1,"Degree":0},
    {"PlanetEngName":"Mars","Sign":1,"House":1,"Degree":10}
  ]
}
```

## Main API

- `kundali.GetYogas(chart)` evaluates the default yoga set and returns `[]kundali.YogaResult`.
- `kundali.DefaultYogas()` returns the default yoga definitions.
- `kundali.EvaluateAll(chart, yogas)` evaluates a custom yoga list.
- `kundali.NewChart(...)` builds a validated chart from placements.
- `kundali.NewChartFromPlanetData(data)` builds a chart from `PlanetData`.
- `kundali.ParseChartJSON(data)` parses chart JSON into a `Chart`.

## Result Type

Each result has:

- `Name`
- `Group`
- `Strength`

Example:

```go
[]kundali.YogaResult{
	{
		Name:     "Ruchaka Yoga",
		Group:    "Pancha Mahapurush",
		Strength: kundali.YogaStrengthStrong,
	},
}
```
