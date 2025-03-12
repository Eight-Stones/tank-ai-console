package probe

import (
	"context"

	"go-micro.dev/v4/server"

	probesv1 "go-micro-service-template/internal/controller/proto/probes/v1"
	er "go-micro-service-template/pkg/error"
)

var _ probesv1.ProbesServiceHandler = &Probe{}

// Probe object for probes.
type Probe struct{}

// New creates new instance of Probe.
func New() *Probe {
	return &Probe{}
}

// GetHealth health probe.
func (p *Probe) GetHealth(_ context.Context, _ *probesv1.HealthRequest, _ *probesv1.HealthResponse) error {
	return nil
}

// GetReady ready probe.
func (p *Probe) GetReady(_ context.Context, _ *probesv1.ReadyRequest, _ *probesv1.ReadyResponse) error {
	return nil
}

// Register register handlers on grpc server.
func (p *Probe) Register(server server.Server) error {
	err := probesv1.RegisterProbesServiceHandler(server, p)
	if err != nil {
		return er.Wrap(err, "register grpc handler probes error")
	}
	return nil
}
