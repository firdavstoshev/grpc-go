package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"sync"

	pb "github.com/firdavstoshev/grpc-go/account"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAccountServiceServer
	accounts map[string]*pb.Account
	mu       sync.Mutex
}

func (s *server) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	account, exists := s.accounts[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Account not found")
	}
	return &pb.GetAccountResponse{Account: account}, nil
}

func (s *server) UpdateAccountName(ctx context.Context, req *pb.UpdateAccountNameRequest) (*pb.UpdateAccountNameResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	account, exists := s.accounts[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Account not found")
	}
	account.Name = req.NewName
	return &pb.UpdateAccountNameResponse{Account: account}, nil
}

func (s *server) UpdateAccountBalance(ctx context.Context, req *pb.UpdateAccountBalanceRequest) (*pb.UpdateAccountBalanceResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	account, exists := s.accounts[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Account not found")
	}
	account.Balance = req.NewBalance
	return &pb.UpdateAccountBalanceResponse{Account: account}, nil
}

func (s *server) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := fmt.Sprintf("%d", len(s.accounts)+1)
	account := &pb.Account{Id: id, Name: req.Name, Balance: 0}
	s.accounts[id] = account
	return &pb.CreateAccountResponse{Account: account}, nil
}

func (s *server) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.accounts, req.Id)
	return &pb.DeleteAccountResponse{Message: "Account deleted"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAccountServiceServer(s, &server{accounts: make(map[string]*pb.Account)})

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
