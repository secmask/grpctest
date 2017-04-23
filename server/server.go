package main

import (
	"log"
	"net"

	pb "github.com/secmask/grpctest"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net/http"
	_ "net/http/pprof"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Helloerver is used to implement helloworld.GreeterServer.erver is used to implement helloworld.GreeterServer.erver is used to implement helloworld.GreeterServer.erver is used to implement helloworld.GreeterServer.erver is used to implement helloworld.GreeterServer.erver is used to implement helloworld.GreeterServer.erver is used to implement helloworld.GreeterServer.erver is used to implement helloworld.GreeterServer.erver is used to implement helloworld.GreeterServer.erver is used to implement helloworld.GreeterServer.erver is used to implement helloworld.GreeterServer.erver is used to implement helloworld.GreeterServer.erver is used to implement helloworld.GreeterServer.erver is used to implement helloworld.GreeterServer.erver is used to implement helloworld.GreeterServer.erver is used to implement helloworld.GreeterServer.erver is used to implement helloworld.GreeterServer. "}, nil
}

func (s *server) Push(push pb.Greeter_PushServer) error {
	msg := pb.Message{}
	for {
		err := push.RecvMsg(&msg)
		if err != nil {
			log.Println(err)
			break
		}
		//log.Printf("%s-%s\n", msg.Channel, msg.Data)
	}
	return nil
}

func main() {
	grpc.EnableTracing = false
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go http.ListenAndServe(":8085", nil)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	s.Stop()
}
