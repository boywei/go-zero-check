package test

import "github.com/boywei/go-zero-check/internal/model"

var (
	Model1 = &model.Uppaal{
		SystemDeclaration: model.SystemDeclaration{
			Models: []string{"model1", "model2"},
		},
		Automatons: []model.Automaton{
			{
				Name:       "a",
				Parameters: []model.Parameter{"automaton_parameter1", "automaton_parameter2"},
			},
		},
		Declaration: model.Declaration{
			Functions: []model.Function{
				{
					Name:       "f",
					Parameters: []model.Parameter{"x", "y"},
				},
			},
		},
	}
)
