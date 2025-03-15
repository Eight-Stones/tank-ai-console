package view

import (
	"context"
	"time"
)

func (v *View) drawJob(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(v.opt.redrawTimeout)
		for {
			select {
			case <-ticker.C:
				v.app.ForceDraw()
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (v *View) createViewJob(ctx context.Context) {
	go func() {
		id := ""
		ticker := time.NewTicker(time.Millisecond * 1000)
		for {
			select {
			case getID := <-v.jobs.start:
				id = getID
			case <-v.jobs.end:
				id = ""
			case <-ticker.C:
				if id != "" {
					v.readGame(id)
					v.readView(id)
				}
			case <-ctx.Done():
				return
			}
		}

	}()
}
