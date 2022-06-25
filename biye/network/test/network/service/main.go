package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "network/api/sbi/service/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:8001", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSbiClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// in := &pb.ListSbiReq{PageNum: 1, PageSize: 10}
	// r, err := c.ListSbi(ctx, in)
	// if err != nil {
	// 	log.Fatalf("could not ListSbi: %v", err)
	// }

	// fmt.Println("count:", r.Count)
	// for _, o := range r.Sbis {
	// 	fmt.Printf("Id:%d,Title:%s,Artist:%s,CreateAt:%s\r\n", o.Id, o.Title, o.Artist, o.CreateAt)
	// }
	start := time.Now()
	in := &pb.GetSbiByIdReq{Id: 1}
	o, err := c.GetSbiById(ctx, in)
	if err != nil {
		log.Fatalf("could not GetSbiByIdReq: %v", err)
	}
	fmt.Printf("Id:%d,Title:%s,Artist:%s,CreateAt:%s\r\n", o.Sbi.Id, o.Sbi.Title, o.Sbi.Artist, o.Sbi.CreateAt)
	elapse := time.Since(start)
	fmt.Println("elapsed time is :", elapse)
}
