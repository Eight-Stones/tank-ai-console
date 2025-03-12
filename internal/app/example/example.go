package example

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go-micro-service-template/internal/app/example/config"
	"go-micro-service-template/internal/controller/grpc"
	"go-micro-service-template/internal/controller/rest"
	restexample "go-micro-service-template/internal/controller/rest/handler/example"
	"go-micro-service-template/internal/controller/rest/handler/probe"
	gwexample "go-micro-service-template/internal/gateway/storage/postgres/example"
	logicexample "go-micro-service-template/internal/usecase/example"
	trs "go-micro-service-template/pkg/database/transactioner"
	er "go-micro-service-template/pkg/error"
	"go-micro-service-template/pkg/micro/loggerm"
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
	// -------------------logger-------------------
	// --------------------------------------------

	l, err := loggerm.New(
		loggerm.WithLevelString(cfg.Observability.Logger.Level),
		loggerm.WithInitialFields(cfg.Observability.Logger.Keys),
	)
	if err != nil {
		return er.Wrap(err, "failed to initialize logger")
	}

	// --------------------------------------------
	// -------------------db poll------------------
	// --------------------------------------------

	pool := trs.New(
		trs.WithLogger(l),
		trs.WithHost(cfg.Storage.Book.Host),
		trs.WithPort(cfg.Storage.Book.Port),
		trs.WithDBName(cfg.Storage.Book.DBName),
		trs.WithUsername(cfg.Storage.Book.Username),
		trs.WithPassword(cfg.Storage.Book.Password),
		trs.WithMaxOpenConns(int32(cfg.Storage.Book.MaxOpenConns)),
		trs.WithSSLMode(cfg.Storage.Book.SSLMode),
	)

	if err = pool.Connect(ctx); err != nil {
		return er.Wrap(err, "failed to initialize database pool")
	}

	// --------------------------------------------
	// -------------------gateway------------------
	// --------------------------------------------

	exampleGateway := gwexample.New()

	// --------------------------------------------
	// -------------------logic--------------------
	// --------------------------------------------

	exampleLogic := logicexample.New(
		logicexample.WithTxProvider(pool),
		logicexample.WithExampleGW(exampleGateway),
	)

	// --------------------------------------------
	// -----------------handler--------------------
	// --------------------------------------------

	probeHandler := probe.New()
	exampleHandler := restexample.New(exampleLogic)

	// --------------------------------------------
	// ----------------api server------------------
	// --------------------------------------------

	rc := rest.New(
		rest.WithName("example-rest"),
		rest.WithHost(cfg.Controller.ExampleRest.Host),
		rest.WithPort(cfg.Controller.ExampleRest.Port),
		rest.WithLogger(loggerm.Sugar(l)),
		rest.WithHandler(probeHandler),
		rest.WithHandler(exampleHandler),
	)

	// --------------------------------------------
	// ----------------grpc server-----------------
	// --------------------------------------------

	gs, err := grpc.New(
		grpc.WithName("example-grpc"),
		grpc.WithHost(cfg.Controller.ExampleGrpc.Host),
		grpc.WithPort(cfg.Controller.ExampleGrpc.Port),
		grpc.WithLogger(loggerm.Sugar(l)),
	)
	if err != nil {
		return er.Wrap(err, "failed to initialize grpc server")
	}

	// --------------------------------------------
	// -------------start http server--------------
	// --------------------------------------------

	rc.Start()

	// --------------------------------------------
	// -------------start grpc server--------------
	// --------------------------------------------

	gs.Start()

	<-ctx.Done()

	return nil
}
