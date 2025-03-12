package probe

import (
	"context"

	"go-micro-service-template/internal/controller/rest/dto"
)

// Probes describe object with probes.
type Probes struct{}

func New() *Probes {
	return &Probes{}
}

// GetHealth health probe.
func (p *Probes) GetHealth(_ context.Context, _ *struct{}, rsp *dto.HealthResponse) error {
	rsp.Status = "ok"
	return nil
}

// GetReady ready probe.
func (p *Probes) GetReady(_ context.Context, _ *struct{}, rsp *dto.ReadyResponse) error {
	rsp.Status = "ok"
	return nil
}
