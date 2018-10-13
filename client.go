package main

import (
	"flag"
	"fmt"
	"log"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/manybooks/deleuze/pb"
)

var (
	serverAddr = flag.String("server_addr", "120.0.0.1:10000", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr)
	if err != nil {
		log.Fatalf("fail to dail: %v", err)
	}
	defer conn.Close()
	client := pb.NewOracleClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	q := pb.Question{Question: "what time is it now?", UserId: 100}
	ans, err := client.Reveal(ctx, &q)
	if err != nil {
		log.Fatalf("fail to get answer")
	}
	fmt.Println(ans.Answer)
}