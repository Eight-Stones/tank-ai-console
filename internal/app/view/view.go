package view

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-micro-service-template/internal/app/view/config"
	"go-micro-service-template/internal/gateway/client/manager"
	"go-micro-service-template/internal/usecase/view"
	er "go-micro-service-template/pkg/error"
)

func Run(configPath string) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// --------------------------------------------
	// -------------------config-------------------
	// --------------------------------------------

	cfg, err := config.New(configPath)
	if err != nil {
		return er.Wrap(err, "failed to initialize config")
	}

	// --------------------------------------------
	// ------------create tank manager-------------
	// --------------------------------------------

	tm := manager.New(
		manager.WithHost(cfg.Gateway.TankClient.Host),
		manager.WithPort(cfg.Gateway.TankClient.Port),
		manager.WithTimeout(time.Second*60),
	)

	// --------------------------------------------
	// ----------------create view-----------------
	// --------------------------------------------

	vw := view.New(
		view.WithReDrawTimeout(time.Millisecond*100),
		view.WithManager(tm),
	)
	defer vw.Stop()

	go func() {
		if err = vw.Run(ctx); err != nil {
			panic(err)
		}
	}()

	<-ctx.Done()

	return nil
}
