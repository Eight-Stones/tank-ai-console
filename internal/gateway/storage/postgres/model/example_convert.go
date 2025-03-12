package model

import (
	"go-micro-service-template/entity"
)

func ConvertExampleToModel(in *entity.Example) *Example {
	return &Example{
		ID:   in.ID,
		Code: in.Code,
		Name: in.Name,
		Meta: in.Meta,
	}
}
func ConvertModelToExample(in *Example) *entity.Example {
	return &entity.Example{
		ID:   in.ID,
		Code: in.Code,
		Name: in.Name,
		Meta: in.Meta,
	}
}

func ConvertModelsToExamples(ins []Example) []entity.Example {
	out := make([]entity.Example, 0, len(ins))
	for _, in := range ins {
		exp := ConvertModelToExample(&in)
		out = append(out, *exp)
	}
	return out
}
