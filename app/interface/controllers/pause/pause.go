package pause

import (
	"context"

	"github.com/wz185/go-example/app/interface/grpc/pause"
)

// Handler for Pause service API
type Handler struct {
}

// Create handles create pause request.
func (ph *Handler) Create(ctx context.Context, req *pause.PauseRequest, res *pause.PauseResponse) error {
	res.ResposneInfo = "responded"
	return nil
}
