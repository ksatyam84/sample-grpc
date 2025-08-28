package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/ksatyam84/sample-grpc/user"
	"google.golang.org/grpc"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer
	mu    sync.Mutex
	users map[string]*pb.User
}

func newServer() *userServiceServer {
	return &userServiceServer{
		users: make(map[string]*pb.User),
	}
}

func (s *userServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user := req.GetUser()
	_, ok := s.users[user.Id]
	if ok {
		return nil, fmt.Errorf("user with id %s already exists", user.Id)
	}
	s.users[user.Id] = user
	log.Printf("User created: %v", user)

	return &pb.CreateUserResponse{
		User: user,
	}, nil
}

func (c *userServiceServer) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	user := in.GetUser()
	_, ok := c.users[user.Id]
	if !ok {
		return nil, fmt.Errorf("user with id %s not found", user.Id)
	}
	c.users[user.Id] = user
	log.Printf("User updated: %v", user)
	return &pb.UpdateUserResponse{
		User: user,
	}, nil

}

func (c *userServiceServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	userId := in.GetId()
	user, ok := c.users[userId]
	if !ok {
		return nil, fmt.Errorf("user with id %s not found", userId)
	}
	log.Printf("User retrieved: %v", user)
	return &pb.GetUserResponse{
		User: user,
	}, nil
}

func (c *userServiceServer) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	userId := in.GetId()
	_, ok := c.users[userId]
	if !ok {
		return nil, fmt.Errorf("user with id %s not found", userId)
	}
	delete(c.users, userId)
	log.Printf("User deleted: %s", userId)
	return &pb.DeleteUserResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, newServer())
	log.Println("gRPC server listening on port 5001")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
