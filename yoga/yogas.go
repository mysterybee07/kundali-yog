package yoga

var Yogas = []Yoga{
	{
		Name: "Gajakesari Yoga",
		Conditions: []Condition{
			PlanetExists(Moon),
			PlanetExists(Jupiter),
			IsKendraFrom(Jupiter, Moon),
		},
		WeakIf: []Condition{
			IsCombust(Moon),
			IsCombust(Jupiter),
		},
	},
	{
		Name: "Pancha Mahapurush Yoga",
		Conditions: []Condition{
			AnyCondition(
				IsPanchaMahapurushPlanet(Mars),
				IsPanchaMahapurushPlanet(Mercury),
				IsPanchaMahapurushPlanet(Jupiter),
				IsPanchaMahapurushPlanet(Venus),
				IsPanchaMahapurushPlanet(Saturn),
			),
		},
		WeakIf: []Condition{
			AllConditions(IsPanchaMahapurushPlanet(Mars), IsCombust(Mars)),
			AllConditions(IsPanchaMahapurushPlanet(Mercury), IsCombust(Mercury)),
			AllConditions(IsPanchaMahapurushPlanet(Jupiter), IsCombust(Jupiter)),
			AllConditions(IsPanchaMahapurushPlanet(Venus), IsCombust(Venus)),
			AllConditions(IsPanchaMahapurushPlanet(Saturn), IsCombust(Saturn)),
		},
	},
	{
		Name:  "Ruchaka Yoga",
		Group: "Pancha Mahapurush Yoga",
		Conditions: []Condition{
			IsPanchaMahapurushPlanet(Mars),
		},
		WeakIf: []Condition{
			IsCombust(Mars),
		},
	},
	{
		Name:  "Bhadra Yoga",
		Group: "Pancha Mahapurush Yoga",
		Conditions: []Condition{
			IsPanchaMahapurushPlanet(Mercury),
		},
		WeakIf: []Condition{
			IsCombust(Mercury),
		},
	},
	{
		Name:  "Hamsa Yoga",
		Group: "Pancha Mahapurush Yoga",
		Conditions: []Condition{
			IsPanchaMahapurushPlanet(Jupiter),
		},
		WeakIf: []Condition{
			IsCombust(Jupiter),
		},
	},
	{
		Name:  "Malavya Yoga",
		Group: "Pancha Mahapurush Yoga",
		Conditions: []Condition{
			IsPanchaMahapurushPlanet(Venus),
		},
		WeakIf: []Condition{
			IsCombust(Venus),
		},
	},
	{
		Name:  "Sasa Yoga",
		Group: "Pancha Mahapurush Yoga",
		Conditions: []Condition{
			IsPanchaMahapurushPlanet(Saturn),
		},
		WeakIf: []Condition{
			IsCombust(Saturn),
		},
	},
	{
		Name: "Budhaditya Yoga",
		Conditions: []Condition{
			PlanetExists(Sun),
			PlanetExists(Mercury),
			AreConjunct(ConjunctionConfig{Type: HouseBased}, Sun, Mercury),
		},
		WeakIf: []Condition{
			IsCombust(Mercury),
		},
	},
	{
		Name: "Chandra Mangala Yoga",
		Conditions: []Condition{
			PlanetExists(Moon),
			PlanetExists(Mars),
			AreConjunct(ConjunctionConfig{Type: HouseBased}, Moon, Mars),
		},
	},
	{
		Name: "Guru Mangala Yoga",
		Conditions: []Condition{
			PlanetExists(Jupiter),
			PlanetExists(Mars),
			AreConjunct(ConjunctionConfig{Type: HouseBased}, Jupiter, Mars),
		},
	},
	{
		Name: "Sunapha Yoga",
		Conditions: []Condition{
			PlanetExists(Moon),
			AnyPlanetInHouseFrom(Moon, 2, Mars, Mercury, Jupiter, Venus, Saturn, Sun),
		},
	},
	{
		Name: "Anapha Yoga",
		Conditions: []Condition{
			PlanetExists(Moon),
			AnyPlanetInHouseFrom(Moon, 12, Mars, Mercury, Jupiter, Venus, Saturn, Sun),
		},
	},
	{
		Name: "Durudhara Yoga",
		Conditions: []Condition{
			PlanetExists(Moon),
			AnyPlanetInHouseFrom(Moon, 2, Mars, Mercury, Jupiter, Venus, Saturn, Sun),
			AnyPlanetInHouseFrom(Moon, 12, Mars, Mercury, Jupiter, Venus, Saturn, Sun),
		},
	},
	{
		Name: "Kemadruma Yoga",
		Conditions: []Condition{
			PlanetExists(Moon),
			NoPlanetInHouseFrom(Moon, 2, Mars, Mercury, Jupiter, Venus, Saturn, Sun),
			NoPlanetInHouseFrom(Moon, 12, Mars, Mercury, Jupiter, Venus, Saturn, Sun),
		},
	},
	{
		Name: "Amala Yoga",
		Conditions: []Condition{
			AnyPlanetInHouseFrom(Asc, 10, Mercury, Jupiter, Venus),
		},
	},
	{
		Name: "Saraswati Yoga",
		Conditions: []Condition{
			PlanetExists(Jupiter),
			PlanetExists(Venus),
			PlanetExists(Mercury),
			AnyCondition(IsOwnSign(Jupiter), IsExalted(Jupiter)),
			AnyCondition(IsOwnSign(Venus), IsExalted(Venus)),
			AnyCondition(IsOwnSign(Mercury), IsExalted(Mercury)),
		},
	},
	{
		Name: "Adhi Yoga",
		Conditions: []Condition{
			PlanetExists(Moon),
			AnyCondition(
				AnyPlanetInHouseFrom(Moon, 6, Mercury, Jupiter, Venus),
				AnyPlanetInHouseFrom(Moon, 7, Mercury, Jupiter, Venus),
				AnyPlanetInHouseFrom(Moon, 8, Mercury, Jupiter, Venus),
			),
		},
	},
	{
		Name: "Raj Yoga",
		Conditions: []Condition{
			PlanetExists(Asc),
			IsRajYoga(),
		},
	},
}
