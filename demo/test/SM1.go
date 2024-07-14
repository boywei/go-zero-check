package test

var c1, c2, c3, c4, c5 bool

var (
	SM1 = &Automaton{
		Name: "SM1",
		Parameters: "c1, c2, c3, c4, c5 bool",
		Locations: []Location{
			{
				Id: 1,
				Name: "A",
				Invariant: func() bool {
					return true
				},
			},
			{
				Id: 2,
				Name: "B",
				Invariant: func() bool {
					return true
				},
			},
			{
				Id: 3,
				Name: "C",
				Invariant: func() bool {
					return true
				},
			},
			{
				Id: 4,
				Name: "D",
				Invariant: func() bool {
					return true
				},
			},
			{
				Id: 5,
				Name: "E",
				Invariant: func() bool {
					return true
				},
			},
			{
				Id: 6,
				Name: "",
				Invariant: func() bool {
					return true
				},
			},
			{
				Id: 7,
				Name: "",
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
					return c1
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					S1()
				},
			},
			{
				Id: 2,
				SourceId: 6,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return !c2
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					S3()
				},
			},
			{
				Id: 3,
				SourceId: 6,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return c2
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					S2()
				},
			},
			{
				Id: 4,
				SourceId: 7,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return c4
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					S4()
				},
			},
			{
				Id: 5,
				SourceId: 7,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return !c4 && c5
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					S5()
				},
			},
			{
				Id: 6,
				SourceId: 7,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return !c4 && !c5
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					S6()
				},
			},
		},
		Init: 1,
	}

)

func S1(){}

func S2(){}

func S3(){}

func S4(){}

func S5(){}

func S6(){}