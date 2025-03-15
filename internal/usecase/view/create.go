package view

import (
	"context"

	er "go-micro-service-template/pkg/error"
)

func (v *View) createGame(ctx context.Context) error {
	if _, err := v.manager.CreateGame(ctx); err != nil {
		return er.Wrap(err, "CreateGame")
	}

	v.app.SetFocus(v.controls)

	return nil
}
