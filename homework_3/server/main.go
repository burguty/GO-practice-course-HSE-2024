package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"homework_3/models"
	"homework_3/proto"
	"net"
)

type Server struct {
	database *models.AccountDatabase
	proto.UnimplementedBankServer
}

func (s *Server) CreateAccount(ctx context.Context,
	req *proto.CreateAccountRequest) (*proto.CreateAccountReply, error) {
	fmt.Printf("Command: %s \t Received: %s\n", "CreateAccount", req.Name)
	err := s.database.CreateAccount(req.Name)
	if err != nil {
		return nil, err
	}
	return &proto.CreateAccountReply{Message: fmt.Sprintf("Account with name '%s' created", req.Name)}, nil
}

func (s *Server) DeleteAccount(ctx context.Context,
	req *proto.DeleteAccountRequest) (*proto.DeleteAccountReply, error) {
	fmt.Printf("Command: %s \t Received: %s\n", "DeleteAccount", req.Name)
	err := s.database.DeleteAccount(req.Name)
	if err != nil {
		return nil, err
	}
	return &proto.DeleteAccountReply{Message: fmt.Sprintf("Account with name '%s' deleted", req.Name)}, nil
}

func (s *Server) UpdateAmount(ctx context.Context,
	req *proto.UpdateAmountRequest) (*proto.UpdateAmountReply, error) {
	fmt.Printf("Command: %s \t Received: %s, %d\n", "UpdateAmount", req.Name, req.Amount)
	err := s.database.UpdateAmount(req.Name, req.Amount)
	if err != nil {
		return nil, err
	}
	return &proto.UpdateAmountReply{Message: fmt.Sprintf("Transfered %d to account '%s'",
		req.Amount, req.Name)}, nil
}

func (s *Server) UpdateName(ctx context.Context, req *proto.UpdateNameRequest) (*proto.UpdateNameReply, error) {
	fmt.Printf("Command: %s \t Received: %s, %s\n", "UpdateName", req.Name, req.NewName)
	err := s.database.UpdateName(req.Name, req.NewName)
	if err != nil {
		return nil, err
	}
	return &proto.UpdateNameReply{Message: fmt.Sprintf("Name of account '%s' updated to '%s'",
		req.Name, req.NewName)}, nil
}

func (s *Server) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountReply, error) {
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
	proto.RegisterBankServer(ser, &Server{database: models.NewAccountDatabase()})
	if err := ser.Serve(lis); err != nil {
		panic(err)
	}
}
