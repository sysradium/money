package handler

import (
	"context"
	"time"

	"github.com/sysradium/money/money"
	pb "github.com/sysradium/money/proto"
)

type Handler struct {
}

func (h *Handler) UpdateQuote(ctx context.Context, in *pb.Quote, out *pb.Quote) error {
	price := in.Price.Div(money.NewFromFloat(3))
	out.Price = price
	now := time.Now()
	out.Timestamp = &now
	return nil
}
