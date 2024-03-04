package test

type Automaton struct {
	Name        string       `json:"name"`
	Parameters  []string     `json:"parameters"`
	Locations   []Location   `json:"locations"`
	Transitions []Transition `json:"transitions"`
	Init        int          `json:"init"`

	Declaration  string   `json:"declaration"`
	InitLocation Location //需要解析
}

type Parameter string

type Declaration string

// 第一层次: 自动机的主要元素

type Location struct {
	State
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Invariant Invariant `json:"invariant"`
}

type Nail struct {
	State
	Id int `json:"id"`
}

type State interface {
}

type Transition struct {
	Id            int    `json:"id"`
	SourceId      int    `json:"source_id"`
	DestinationId int    `json:"destination_id"`
	Select        Select `json:"select"`
	Guard         Guard  `json:"guard"`
	Sync          Sync   `json:"sync"`
	Update        Update `json:"update"`
	Source        State  // 需要解析
	Destination   State  // 需要解析
}

// TODO: 第二层次: 自动机各元素上的标签

type Select string

type Guard string

type Sync string

type Update string

type Invariant string

type Probability struct {
}
