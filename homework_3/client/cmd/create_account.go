package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"homework_3/proto"
	"time"
)

var createAccountCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new bank account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port),
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}

		defer func() {
			_ = conn.Close()
		}()

		cli := proto.NewBankClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		rep, err := cli.CreateAccount(ctx, &proto.CreateAccountRequest{Name: name})
		if err != nil {
			panic(err)
		}
		fmt.Println(rep.Message)
	},
}

func init() {
	rootCmd.AddCommand(createAccountCmd)
}
