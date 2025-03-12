# golang-service-template
target on developing of simple golang web service

## ability roadmap
### layers
- [x] app builder layer
- [x] rest controller layer
- [x] usecase layer
- [x] gateway repo with pg example
### infrastructure
- [x] working docker-compose example app
- [x] docker example for example app
- [x] docker example for migrator app
### database
- [x] working node of local pg
- [x] special object for building query aims
- [x] simple transaction manager implementations
- [x] example applying migrations by goose cmd 
- [x] example applying migrations by migrator app
### generator
- [x] using go-buf for generating proto api
### observability 
- [x] working grafana
- [x] golang system dashboard
- [ ] service info dashboard
- [ ] jaeger trace collector
- [ ] loki logs collector
- [ ] alert manager 
### structure and package
- [x] struct error package
- [x] struct logger package (with go-micro adapt)
- [x] fasthttp server package (with go-micro adapt, not mine package)
- [x] config package by toml format (with go-micro adapt)
### scripts
- [x] sh scripts with dependencies
### support
- [x] Makefile
- [x] Build commands
- [x] `Golangci-lint` utility
- [x] `Golang Arch-lint` utility
- [x] `Golang Arch-lint` basic setup
- [x] `Golang Arch-lint` graph setup 
- [ ] `Golang Arch-lint` vendor fine tuning