package model127_0_0_1

var (
	Car = &Automaton{
		Name: "Car",
		Parameters: "",
		Locations: []Location{
			{
				Id: 3,
				Name: "Keep",
				Invariant: func() bool {
					return true
				},
			},
			{
				Id: 4,
				Name: "SpeedDown",
				Invariant: func() bool {
					return true
				},
			},
			{
				Id: 6,
				Name: "SpeedUp",
				Invariant: func() bool {
					return true
				},
			},
		},
		Transitions: []Transition{
			{
				Id: 5,
				SourceId: 6,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return v <= 150 && v >= 50
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					
				},
			},
			{
				Id: 6,
				SourceId: 3,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return v >= 100
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					
				},
			},
			{
				Id: 11,
				SourceId: 6,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return true
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					v >= 150
				},
			},
			{
				Id: 12,
				SourceId: 3,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return v <= 100
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					
				},
			},
			{
				Id: 13,
				SourceId: 6,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return true
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					speed_up()
				},
			},
			{
				Id: 14,
				SourceId: 4,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return true
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					speed_down()
				},
			},
			{
				Id: 15,
				SourceId: 3,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return true
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					keep()
				},
			},
			{
				Id: 16,
				SourceId: 4,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return v >= 50 && v <= 150
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					
				},
			},
			{
				Id: 17,
				SourceId: 4,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return true
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					v <= 100
				},
			},
		},
		Init: 1,
	}

)

var v = 100

func speed_up() {
    v = v + 10
}

func speed_down() {
    v = v - 10
}

func keep() {
    
}