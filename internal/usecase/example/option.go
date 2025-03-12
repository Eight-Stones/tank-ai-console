package example

import (
	txp "go-micro-service-template/pkg/database/transactioner"
)

type options struct {
	txp     *txp.TxProvider
	gateway Gateways
}

type Option func(*options)

func WithTxProvider(txp *txp.TxProvider) Option {
	return func(o *options) {
		o.txp = txp
	}
}

func WithExampleGW(in Exampler) Option {
	return func(o *options) {
		o.gateway.ex = in
	}
}
