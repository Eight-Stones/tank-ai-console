package convert

import (
	"go-micro-service-template/entity"
	"go-micro-service-template/internal/controller/rest/dto"
)

func DTOToExample(in *dto.Example) *entity.Example {
	return &entity.Example{
		ID:   in.ID,
		Code: in.Code,
		Name: in.Name,
		Meta: in.Meta,
	}
}

func ExampleToDTO(in *entity.Example) *dto.Example {
	return &dto.Example{
		ID:   in.ID,
		Code: in.Code,
		Name: in.Name,
		Meta: in.Meta,
	}
}

func ExamplesToDTOs(in []entity.Example) []*dto.Example {
	out := make([]*dto.Example, 0, len(in))
	for _, example := range in {
		out = append(out, ExampleToDTO(&example))
	}
	return out
}
