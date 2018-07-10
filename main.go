package main

import (
	"github.com/coreos/rkt/tests/testutils/logger"
	microgrpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/transport/grpc"
	"github.com/pkg/errors"
	"github.com/sysradium/money/app/handler"
	money "github.com/sysradium/money/proto"
)

func main() {
	service := microgrpc.NewService(
		micro.Name("quotes"),
		micro.Version("1.0.0"),
	)

	service.Init()

	money.RegisterQuotesHandler(
		service.Server(),
		&handler.Handler{},
	)

	if err := service.Run(); err != nil {
		logger.Fatalf("%#v", errors.Wrapf(err, "failed to start %s", "money"))
	}

}
