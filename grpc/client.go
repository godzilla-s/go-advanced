package grpc

import (
	"go-advanced/grpc/inf"
	"log"
	"math/rand"

	"google.golang.org/grpc"
)

const (
	dial_tcp = "localhost:51011"
)

func Client() {
	conn, err := grpc.Dial(dial_tcp, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failm to dail:%v", err)
	}
	defer conn.Close()

	client := inf.NewDataClient(conn)

	var req inf.DataReq
	r := rand.Intn(1000)
	req.Id = int32(r)

	resp, err := client.GetData(context.Backgroud(), &req)
	if err != nil {
		log.Fatalf("response err: %v", err)
	}
}
