package yoga

var planetDignities = map[Planet]PlanetDignity{
	Sun: {
		OwnSigns:       []int{5},
		ExaltationSign: 1,
		Debilitation:   7,
	},
	Moon: {
		OwnSigns:       []int{4},
		ExaltationSign: 2,
		Debilitation:   8,
	},
	Mars: {
		OwnSigns:       []int{1, 8},
		ExaltationSign: 10,
		Debilitation:   4,
	},
	Mercury: {
		OwnSigns:       []int{3, 6},
		ExaltationSign: 6,
		Debilitation:   12,
	},
	Jupiter: {
		OwnSigns:       []int{9, 12},
		ExaltationSign: 4,
		Debilitation:   10,
	},
	Venus: {
		OwnSigns:       []int{2, 7},
		ExaltationSign: 12,
		Debilitation:   6,
	},
	Saturn: {
		OwnSigns:       []int{10, 11},
		ExaltationSign: 7,
		Debilitation:   1,
	},
}

var signLords = map[int]Planet{
	1:  Mars,
	2:  Venus,
	3:  Mercury,
	4:  Moon,
	5:  Sun,
	6:  Mercury,
	7:  Venus,
	8:  Mars,
	9:  Jupiter,
	10: Saturn,
	11: Saturn,
	12: Jupiter,
}

func DefaultYogas() []Yoga {
	return []Yoga{
		{
			Name:  "Ruchaka Yoga",
			Group: "Pancha Mahapurush",
			Conditions: []Condition{
				IsPanchaMahapurushPlanet(Mars),
			},
			WeakIf: []Condition{
				IsCombust(Mars),
			},
		},
		{
			Name:  "Bhadra Yoga",
			Group: "Pancha Mahapurush",
			Conditions: []Condition{
				IsPanchaMahapurushPlanet(Mercury),
			},
			WeakIf: []Condition{
				IsCombust(Mercury),
			},
		},
		{
			Name:  "Hamsa Yoga",
			Group: "Pancha Mahapurush",
			Conditions: []Condition{
				IsPanchaMahapurushPlanet(Jupiter),
			},
		},
		{
			Name:  "Malavya Yoga",
			Group: "Pancha Mahapurush",
			Conditions: []Condition{
				IsPanchaMahapurushPlanet(Venus),
			},
			WeakIf: []Condition{
				IsCombust(Venus),
			},
		},
		{
			Name:  "Sasa Yoga",
			Group: "Pancha Mahapurush",
			Conditions: []Condition{
				IsPanchaMahapurushPlanet(Saturn),
			},
		},
		{
			Name:  "Raja Yoga",
			Group: "Raja Yoga",
			Conditions: []Condition{
				IsRajYoga(),
			},
		},
	}
}
