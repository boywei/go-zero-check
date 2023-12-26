package model

type Uppaal struct {
	Declaration       Declaration       `json:"declaration"`
	Automatons        []Automaton       `json:"automatons"`
	SystemDeclaration SystemDeclaration `json:"system_declaration"`
}
