package grpc

import (
	"go-advanced/grpc/inf"
	"log"
	"net"
	"runtime"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	listen_port = ":51011"
)

type Data struct{}

func (r *Data) GetData(ctx context.Context, request *inf.DataReq) (response *inf.DataRsp, err error) {
	response = &inf.DataRsp{
		Name: strconv.Itoa(int(request.Id)) + ":test",
	}
	return response, err
}

func Server() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	lis, err := net.Listen("tcp", listen_port)
	if err != nil {
		log.Fatalf("fail to listen: %#v", listen_port)
	}

	s := grpc.NewServer()
	inf.RegisterDataServer(s, &Data{})
	s.Serve(lis)

	log.Println("grpc server start:", listen_port)
}
