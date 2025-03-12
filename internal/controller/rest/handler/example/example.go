package example

import (
	"context"

	"go-micro-service-template/entity"
	"go-micro-service-template/internal/controller/rest/convert"
	"go-micro-service-template/internal/controller/rest/dto"
	er "go-micro-service-template/pkg/error"
)

type Exampler interface {
	Get(ctx context.Context, id int64) (*entity.Example, error)
	Read(ctx context.Context) ([]entity.Example, error)
	Create(ctx context.Context, in *entity.Example) (int64, error)
	Update(ctx context.Context, in *entity.Example) error
	Delete(ctx context.Context, id int64) error
}

// Example describe object with probes.
type Example struct {
	exampler Exampler
}

// New build new rest controller for examplers.
func New(exampler Exampler) *Example {
	return &Example{
		exampler: exampler,
	}
}

// GetExampleParamId return entity.Example by ID.
func (p *Example) GetExampleParamId(
	ctx context.Context,
	req *dto.GetExampleRequest,
	resp *dto.GetExampleResponse,
) error {
	exp, err := p.exampler.Get(ctx, req.ID)
	if err != nil {
		return er.Wrap(err, "p.exampler.Get")
	}

	resp.Payload = convert.ExampleToDTO(exp)

	return nil
}

// PostExamples return slice of entity.Example.
func (p *Example) PostExamples(ctx context.Context, _ *struct{}, resp *dto.ReadExampleResponse) error {
	exps, err := p.exampler.Read(ctx)
	if err != nil {
		return er.Wrap(err, "p.exampler.Read")
	}

	resp.Payload = convert.ExamplesToDTOs(exps)

	return nil
}

// PostExample create entity.Example.
func (p *Example) PostExample(ctx context.Context, req *dto.CreateExampleRequest, resp *dto.CreateExampleResponse) error {
	id, err := p.exampler.Create(ctx, convert.DTOToExample(req.Payload))
	if err != nil {
		return er.Wrap(err, "p.exampler.Create")
	}

	resp.Payload = &dto.Example{ID: id}

	return nil
}

// PutExampleParamId update entity.Example by ID.
func (p *Example) PutExampleParamId(ctx context.Context, req *dto.UpdateExampleRequest, _ *struct{}) error {
	exp := convert.DTOToExample(req.Payload)
	exp.ID = req.ID

	if err := p.exampler.Update(ctx, exp); err != nil {
		return er.Wrap(err, "p.exampler.Update")
	}

	return nil
}

// DeleteExampleParamId delete entity.Example by ID.
func (p *Example) DeleteExampleParamId(ctx context.Context, req *dto.DeleteExampleRequest, _ *struct{}) error {
	if err := p.exampler.Delete(ctx, req.ID); err != nil {
		return er.Wrap(err, "p.exampler.Delete")
	}

	return nil
}
