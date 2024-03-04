package test

var (
	ego = &Automaton{
		Name: "ego",		Parameters: []string{
			"automaton_parameter1",
			"automaton_parameter2",
		},
		Locations: []Location{
			{
				Id: 1,
				Name: "start",
				Invariant: "",
			},
			{
				Id: 2,
				Name: "end",
				Invariant: "",
			},
		},
		Transitions: []Transition{
			{
				Id: 1,
				SourceId: 1,
				DestinationId: 0,
				Select: "",
				Guard: "clock >= 0",
				Sync: "",
				Update: "decrease()",
			},
			{
				Id: 2,
				SourceId: 2,
				DestinationId: 0,
				Select: "",
				Guard: "clock < 0",
				Sync: "",
				Update: "clock += 1",
			},
		},
		Init: 1,
	}

)

var clock = 0; func increase() { clock += 1 }; func decrease() { clock -= 1 }