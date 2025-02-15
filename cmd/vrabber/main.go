package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/bonefabric/vrabber/gen/client"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDownloadServiceServer
}

func (s *server) DownloadVideo(stream pb.DownloadService_DownloadVideoServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		fmt.Printf("Request received: %v\n", req)

		if err := stream.Send(&pb.DownloadStatusResponse{
			RequestId: req.RequestId,
			Status:    pb.DownloadStatus_DONE,
			Message:   "",
		}); err != nil {
			return err
		}
	}
}

func main() {
	//todo stub
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer func(lis net.Listener) {
		if err = lis.Close(); err != nil {
			log.Printf("failed to close listener: %v", err)
		}
	}(lis)

	grpcServer := grpc.NewServer()
	pb.RegisterDownloadServiceServer(grpcServer, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
