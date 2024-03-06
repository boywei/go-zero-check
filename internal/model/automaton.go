package model

type Automaton struct {
	Name        string       `json:"name"`
	Parameters  []Parameter  `json:"parameters"`
	Locations   []Location   `json:"locations"`
	Transitions []Transition `json:"transitions"`
	Init        int          `json:"init"`

	Declaration  Declaration `json:"declaration"`
	InitLocation Location    //需要解析
}

type Parameter string

// 第一层次: 自动机的主要元素

type Location struct {
	State
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Invariant string `json:"invariant"`
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
	Select        string `json:"select"`
	Guard         string `json:"guard"`
	Sync          string `json:"sync"`
	Update        string `json:"update"`
	Source        State  // 需要解析
	Destination   State  // 需要解析
}

// TODO: 第二层次: 自动机各元素上的标签

type Probability struct {
}
