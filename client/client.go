package main

import (
	pb "github.com/secmask/grpctest"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"sync/atomic"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	var cc int64 = 0
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	for i := 0; i < 32; i++ {
		go func() {
			c := pb.NewGreeterClient(conn)
			if err != nil {
				log.Panicln(err)
			}
			for {
				start := time.Now()
				_, err = c.SayHello(context.Background(), &pb.HelloRequest{Name: defaultName})
				//err = stream.Send(&pb.Message{Channel: "hello", Data: "world"})
				now := time.Now()
				if now.Sub(start).Seconds() > 0.5 {
					log.Printf("slow %f\n", now.Sub(start).Seconds())
				}
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
				atomic.AddInt64(&cc, 1)
			}
		}()
	}
	for range time.Tick(time.Second) {
		log.Printf("rate %d\n", atomic.SwapInt64(&cc, 0))
	}
}
