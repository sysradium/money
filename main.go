package main

import (
	"fmt"
	"time"

	"github.com/coreos/rkt/tests/testutils/logger"
	"github.com/davecgh/go-spew/spew"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
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

	q := &money.Quote{Value: &money.Quote_Name{Name: "Foo"}}
	now := time.Now()
	q.Timestamp = &now
	/*
		m := &jsonpb.Marshaler{}
		out := &bytes.Buffer{}
		m.Marshal(out, q)

		q2 := &money.Quote{}
		if err := jsonpb.Unmarshal(out, q2); err != nil {
			log.Fatal(err)
		}
	*/
	out, _ := proto.Marshal(q)
	q2 := &money.Quote{}
	proto.Unmarshal(out, q2)

	switch i := q2.Value.(type) {
	case *money.Quote_Name:
		fmt.Println(i.Name)
	}
	spew.Dump(q2)

	m := &jsonpb.Marshaler{}
	fmt.Println(m.MarshalToString(q2))
	return
	if err := service.Run(); err != nil {
		logger.Fatalf("%#v", errors.Wrapf(err, "failed to start %s", "money"))
	}

}
