package model

type Automaton struct {
	Name        string
	Parameters  []Parameter
	Locations   []Location
	Transitions []Transition
	Init        Location

	Declaration Declaration
}

// 第一层次: 自动机的主要元素

type Location struct {
	State
	Id        int
	Name      string
	Invariant Invariant
}

type Nail struct {
	State
	Id int
}

type State interface {
}

type Transition struct {
	Id            int
	SourceId      int
	DestinationId int
	Select        Select
	Guard         Guard
	Sync          Sync
	Update        Update
	Source        State // 需要解析
	Destination   State // 需要解析
}

// TODO: 第二层次: 自动机各元素上的标签

type Select struct {
}

type Guard struct {
}

type Sync interface {
}

type Update struct {
}

type Invariant struct {
}

type Probability struct {
}
