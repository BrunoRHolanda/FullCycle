package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/BrunoRHolanda/FullCycle/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Cold not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	// AddUser(client)

	//AddUserVerbose(client)

	// AddUsers(client)

	AddUserStreamBoth(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Camila",
		Email: "mcamilabf@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Cold not make to gRPC Request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Camila",
		Email: "mcamilabf@gmail.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Cold not make to gRPC Request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Cold not receive the msg: %v", err)
		}

		fmt.Println("Status:", stream.Status)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "B1",
			Name:  "Bruno 1",
			Email: "bruno1@email.com",
		},
		&pb.User{
			Id:    "B2",
			Name:  "Bruno 2",
			Email: "bruno2@email.com",
		},
		&pb.User{
			Id:    "B3",
			Name:  "Bruno 3",
			Email: "bruno3@email.com",
		},
		&pb.User{
			Id:    "B3",
			Name:  "Bruno 3",
			Email: "bruno3@email.com",
		},
		&pb.User{
			Id:    "B3",
			Name:  "Bruno 3",
			Email: "bruno3@email.com",
		},
		&pb.User{
			Id:    "B3",
			Name:  "Bruno 3",
			Email: "bruno3@email.com",
		},
		&pb.User{
			Id:    "B3",
			Name:  "Bruno 3",
			Email: "bruno3@email.com",
		},
		&pb.User{
			Id:    "B3",
			Name:  "Bruno 3",
			Email: "bruno3@email.com",
		},
		&pb.User{
			Id:    "B3",
			Name:  "Bruno 3",
			Email: "bruno3@email.com",
		},
		&pb.User{
			Id:    "B3",
			Name:  "Bruno 3",
			Email: "bruno3@email.com",
		},
		&pb.User{
			Id:    "B3",
			Name:  "Bruno 3",
			Email: "bruno3@email.com",
		},
		&pb.User{
			Id:    "B3",
			Name:  "Bruno 3",
			Email: "bruno3@email.com",
		},
		&pb.User{
			Id:    "B3",
			Name:  "Bruno 3",
			Email: "bruno3@email.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatal("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		// time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatal("Error receiving response: %v", err)
	}

	fmt.Println(res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {
	stream, err := client.AddUserStreamBoth(context.Background())

	if err != nil {
		log.Fatal("Error creating request: %v", err)
	}

	reqs := []*pb.User{
		&pb.User{
			Id:    "B1",
			Name:  "Bruno 1",
			Email: "bruno1@email.com",
		},
		&pb.User{
			Id:    "B2",
			Name:  "Bruno 2",
			Email: "bruno2@email.com",
		},
		&pb.User{
			Id:    "B3",
			Name:  "Bruno 3",
			Email: "bruno3@email.com",
		},
		&pb.User{
			Id:    "B3",
			Name:  "Bruno 3",
			Email: "bruno3@email.com",
		},
		&pb.User{
			Id:    "B4",
			Name:  "Bruno 4",
			Email: "bruno4@email.com",
		},
		&pb.User{
			Id:    "B5",
			Name:  "Bruno 5",
			Email: "bruno5@email.com",
		},
		&pb.User{
			Id:    "B6",
			Name:  "Bruno 6",
			Email: "bruno6@email.com",
		},
		&pb.User{
			Id:    "B7",
			Name:  "Bruno 7",
			Email: "bruno7@email.com",
		},
		&pb.User{
			Id:    "B8",
			Name:  "Bruno 8",
			Email: "bruno8@email.com",
		},
		&pb.User{
			Id:    "B9",
			Name:  "Bruno 9",
			Email: "bruno9@email.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending User: ", req.GetName())
			stream.Send(req)
			time.Sleep(time.Second * 1)
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
				log.Fatalf("Error receiving data: %v", err)
				break
			}

			fmt.Printf("Recebendo user %v com status %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
