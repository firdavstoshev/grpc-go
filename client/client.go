package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/firdavstoshev/grpc-go/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		if err := conn.Close(); err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}(conn)

	c := pb.NewAccountServiceClient(conn)

	action := flag.String("action", "get", "Action to perform: get, create, update-name, update-balance, delete")
	id := flag.String("id", "", "Account ID")
	name := flag.String("name", "", "Account name")
	balance := flag.Float64("balance", 0, "Account balance")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch *action {
	case "get":
		resp, err := c.GetAccount(ctx, &pb.GetAccountRequest{Id: *id})
		if err != nil {
			log.Fatalf("could not get account: %v", err)
		}
		fmt.Printf("Account: %v\n", resp.Account)
	case "create":
		resp, err := c.CreateAccount(ctx, &pb.CreateAccountRequest{Name: *name})
		if err != nil {
			log.Fatalf("could not create account: %v", err)
		}
		fmt.Printf("Created Account: %v\n", resp.Account)
	case "update-name":
		resp, err := c.UpdateAccountName(ctx, &pb.UpdateAccountNameRequest{Id: *id, NewName: *name})
		if err != nil {
			log.Fatalf("could not update account name: %v", err)
		}
		fmt.Printf("Updated Account: %v\n", resp.Account)
	case "update-balance":
		resp, err := c.UpdateAccountBalance(ctx, &pb.UpdateAccountBalanceRequest{Id: *id, NewBalance: *balance})
		if err != nil {
			log.Fatalf("could not update account balance: %v", err)
		}
		fmt.Printf("Updated Account: %v\n", resp.Account)
	case "delete":
		resp, err := c.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: *id})
		if err != nil {
			log.Fatalf("could not delete account: %v", err)
		}
		fmt.Printf("Delete Account: %s\n", resp.Message)
	default:
		fmt.Println("Unknown action")
	}
}
