package model127_0_0_1

var (
	Nuclear = &Automaton{
		Name: "Nuclear",
		Parameters: "",
		Locations: []Location{
			{
				Id: 1,
				Name: "ShutDown",
				Invariant: func() bool {
					return true
				},
			},
			{
				Id: 2,
				Name: "StartUp",
				Invariant: func() bool {
					return true
				},
			},
			{
				Id: 3,
				Name: "Cooling",
				Invariant: func() bool {
					return true
				},
			},
			{
				Id: 4,
				Name: "Emergency",
				Invariant: func() bool {
					return true
				},
			},
			{
				Id: 5,
				Name: "Critical",
				Invariant: func() bool {
					return true
				},
			},
			{
				Id: 6,
				Name: "PowerDown",
				Invariant: func() bool {
					return true
				},
			},
			{
				Id: 7,
				Name: "PowerUp",
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
			{
				Id: 2,
				SourceId: 2,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return Heat >= MIN_HEAT
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					increase_heat()
				},
			},
			{
				Id: 3,
				SourceId: 5,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return true
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					increase_heat()
				},
			},
			{
				Id: 4,
				SourceId: 7,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return Heat >= MAX_HEAT
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					decrease_heat()
				},
			},
			{
				Id: 5,
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
					decrease_heat()
				},
			},
			{
				Id: 6,
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
					
				},
			},
			{
				Id: 7,
				SourceId: 2,
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
			{
				Id: 8,
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
			{
				Id: 9,
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
					
				},
			},
			{
				Id: 10,
				SourceId: 2,
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
					
				},
			},
			{
				Id: 12,
				SourceId: 7,
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
			{
				Id: 13,
				SourceId: 5,
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
			{
				Id: 14,
				SourceId: 6,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return Heat <= MIN_HEAT
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					increase_heat()
				},
			},
			{
				Id: 15,
				SourceId: 7,
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
			{
				Id: 16,
				SourceId: 1,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return Heat <= NORMAL_HEAT
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					decrease_heat_to_normal()
				},
			},
			{
				Id: 17,
				SourceId: 2,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return Heat <= MIN_HEAT
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					increase_heat_to_min()
				},
			},
			{
				Id: 19,
				SourceId: 7,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return true
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					increase_heat()
				},
			},
			{
				Id: 20,
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
					keep_emergency()
				},
			},
			{
				Id: 21,
				SourceId: 3,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return Heat <= COOL_HEAT
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					decrease_heat_to_cool()
				},
			},
			{
				Id: 22,
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
					decrease_heat()
				},
			},
			{
				Id: 23,
				SourceId: 5,
				DestinationId: 0,
				Select: "",
				Guard: func() bool {
					return Heat <= MAX_HEAT
				},
				Sync: func() bool {
					return true
				},
				Update: func() {
					increase_heat_to_critical()
				},
			},
			{
				Id: 24,
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
					
				},
			},
		},
		Init: 1,
	}

)

// 在这里编写局部变量、函数等.
var Heat = 30

func increase_heat() {
    Heat = Heat + 10
}

func decrease_heat() {
    Heat = Heat - 10
}

func decrease_heat_to_normal() {
    if Heat >= NORMAL_HEAT {
        Heat = Heat - 10
    }
}

func increase_heat_to_min() {
    if Heat <= MIN_HEAT {
        Heat = Heat + 10
    }
}

func keep_emergency() {
    
}

func decrease_heat_to_cool() {
    if Heat >= COOL_HEAT {
        Heat = Heat - 10
    }
}

func increase_heat_to_critical() {
    if Heat <= MIN_HEAT {
        Heat = Heat + 10
    }
}