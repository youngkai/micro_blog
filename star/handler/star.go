package handler

import (
	star "blog/star/proto/star"
	"context"
	log "github.com/micro/go-micro/v2/logger"
)

type Star struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Star) Call(ctx context.Context, req *star.Request, rsp *star.Response) error {
	log.Info("Received Star.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
