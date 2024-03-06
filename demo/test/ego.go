package test

var (
	ego = &Automaton{
		Name: "ego",
		Parameters: []string{
			"automaton_parameter1",
			"automaton_parameter2",
		},
		Locations: []Location{
			{
				Id: 1,
				Name: "start",
				Invariant: func() bool {
					return true
				},
			},
			{
				Id: 2,
				Name: "end",
				Invariant: func() bool {
					return true
				},
			},
		},
		Transitions: []Transition{
			{
				Id: 1,
				SourceId: 1,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return clock >= 0
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					decrease()
				},
			},
			{
				Id: 2,
				SourceId: 2,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return clock < 0
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					clock += 1
				},
			},
		},
		Init: 1,
	}

)

var clock = 0; func increase() { clock += 1 }; func decrease() { clock -= 1 }