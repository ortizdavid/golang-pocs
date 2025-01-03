package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"golang-pocs/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	// AddUserVerbose(client)
	// AddUser(client)
	// AddUsers(client)
	AddUserBidirectional(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Client",
		Email: "client@email.com",
	}

	response, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make the gRPC request: %v", err)
	}

	log.Printf(response.String())
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "User Test",
		Email: "user@email.com",
	}
	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make the gRPC request: %v", err)
	}

	for {
		stream, error := responseStream.Recv()
		if error == io.EOF {
			break
		}
		if error != nil {
			log.Fatalf("Could not receive the message: %v ", err)
		}
		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "1",
			Name:  "Alexandra",
			Email: "alexandra@email.com",
		}, &pb.User{
			Id:    "2",
			Name:  "Chad",
			Email: "chad@email.com",
		}, &pb.User{
			Id:    "3",
			Name:  "Alisson",
			Email: "alisson@email.com",
		}, &pb.User{
			Id:    "4",
			Name:  "John",
			Email: "john@email.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}

func AddUserBidirectional(client pb.UserServiceClient) {

	stream, err := client.AddUserBidirectional(context.Background())
	if err != nil {
		log.Fatalf("Error creating request %v", err)
	}

	reqs := []*pb.User{
		{
			Id:    "1",
			Name:  "Alexandra",
			Email: "alexandra@email.com",
		}, {
			Id:    "2",
			Name:  "Chad",
			Email: "chad@email.com",
		}, {
			Id:    "3",
			Name:  "Alisson",
			Email: "alisson@email.com",
		}, {
			Id:    "4",
			Name:  "John",
			Email: "john@email.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving data, %v", err)
				break
			}
			fmt.Printf("Receiving user %v with status %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait

}
