package test

import (
	"fmt"
	"go-advanced/grpc/inf"
	"log"
	"net"
	"runtime"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	listen_port = ":51011"
)

type Data struct{}

// 实现服务
func (r *Data) GetData(ctx context.Context, request *inf.DataReq) (response *inf.DataRsp, err error) {
	response = &inf.DataRsp{
		Name:  strconv.Itoa(int(request.Id)) + ":hello my new world",
		Times: time.Now().Unix(),
	}
	return response, err
}

func (r *Data) GetString(ctx context.Context, request *inf.DataReq) (response *inf.Data2Rsp, err error) {
	response = &inf.Data2Rsp{
		Resp: fmt.Sprintf("right now : %d", time.Now().UnixNano()),
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
