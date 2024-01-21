package health

import "go_stucture/core"

type healthHandler struct {
}

func NewHealthHandler() *healthHandler {
	return &healthHandler{}
}

func (h *healthHandler) Handle(ctx core.Context) {
	ctx.StatusOK()
}
