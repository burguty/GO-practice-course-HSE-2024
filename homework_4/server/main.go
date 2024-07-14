package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"homework_4/models"
	"homework_4/proto"
	"net"
)

type server struct {
	database *models.AccountDatabase
	proto.UnimplementedBankServer
}

func (s *server) CreateAccount(ctx context.Context,
	req *proto.CreateAccountRequest) (*proto.CreateAccountReply, error) {
	fmt.Printf("Command: %s \t Received: %s\n", "CreateAccount", req.Name)
	err := s.database.CreateAccount(req.Name)
	if err != nil {
		return nil, err
	}
	return &proto.CreateAccountReply{Message: fmt.Sprintf("Account with name '%s' created", req.Name)}, nil
}

func (s *server) DeleteAccount(ctx context.Context,
	req *proto.DeleteAccountRequest) (*proto.DeleteAccountReply, error) {
	fmt.Printf("Command: %s \t Received: %s\n", "DeleteAccount", req.Name)
	err := s.database.DeleteAccount(req.Name)
	if err != nil {
		return nil, err
	}
	return &proto.DeleteAccountReply{Message: fmt.Sprintf("Account with name '%s' deleted", req.Name)}, nil
}

func (s *server) UpdateAmount(ctx context.Context,
	req *proto.UpdateAmountRequest) (*proto.UpdateAmountReply, error) {
	fmt.Printf("Command: %s \t Received: %s, %d\n", "UpdateAmount", req.Name, req.Amount)
	err := s.database.UpdateAmount(req.Name, req.Amount)
	if err != nil {
		return nil, err
	}
	return &proto.UpdateAmountReply{Message: fmt.Sprintf("Transfered %d to account '%s'",
		req.Amount, req.Name)}, nil
}

func (s *server) UpdateName(ctx context.Context, req *proto.UpdateNameRequest) (*proto.UpdateNameReply, error) {
	fmt.Printf("Command: %s \t Received: %s, %s\n", "UpdateName", req.Name, req.NewName)
	err := s.database.UpdateName(req.Name, req.NewName)
	if err != nil {
		return nil, err
	}
	return &proto.UpdateNameReply{Message: fmt.Sprintf("Name of account '%s' updated to '%s'",
		req.Name, req.NewName)}, nil
}

func (s *server) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountReply, error) {
	fmt.Printf("Command: %s \t Received: %s\n", "GetAccount", req.Name)
	account, err := s.database.GetAccount(req.Name)
	if err != nil {
		return nil, err
	}
	return &proto.GetAccountReply{Message: fmt.Sprintf("Name: %s\tAmount: %d",
		account.Name, account.Amount)}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4567))
	if err != nil {
		panic(err)
	}

	ser := grpc.NewServer()

	connectionString := "host=0.0.0.0 port=5432 dbname=postgres user=postgres password=12345"
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}

	defer func() {
		_ = db.Close()
	}()

	proto.RegisterBankServer(ser, &server{database: models.ConnectAccountDatabase(db)})
	if err = ser.Serve(lis); err != nil {
		panic(err)
	}
}
