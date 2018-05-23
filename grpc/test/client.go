package test

import (
	"fmt"
	"go-advanced/grpc/inf"
	"log"
	"math/rand"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	dial_tcp = "localhost:51011"
)

func Exec() {
	conn, err := grpc.Dial(dial_tcp, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failm to dail:%v", err)
	}
	defer conn.Close()

	client := inf.NewDataClient(conn)

	var req inf.DataReq
	r := rand.Intn(1000)
	req.Id = int32(r)

	resp, err := client.GetData(context.Background(), &req)
	if err != nil {
		log.Fatalf("response err: %v", err)
	}

	fmt.Println(resp.Name, resp.Times)

	resp2, err := client.GetString(context.Background(), &req)
	if err != nil {
		log.Fatalf("response err: %v", err)
	}

	fmt.Println(resp2.Resp)
}

func Client() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		Exec()

	}
}
