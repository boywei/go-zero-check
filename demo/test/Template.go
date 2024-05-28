package test

var id int

var (
	Template = &Automaton{
		Name: "Template",
		Parameters: "id int",
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
					return true
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					
				},
			},
		},
		Init: 1,
	}

)

// Place local declarations here.