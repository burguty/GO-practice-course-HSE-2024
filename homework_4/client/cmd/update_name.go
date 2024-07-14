package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"homework_4/proto"
	"time"
)

var updateNameCmd = &cobra.Command{
	Use:   "update-name [name] [new name]",
	Short: "Change name of particular account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		newName := args[1]
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

		rep, err := cli.UpdateName(ctx, &proto.UpdateNameRequest{Name: name, NewName: newName})
		if err != nil {
			panic(err)
		}
		fmt.Println(rep.Message)
	},
}

func init() {
	rootCmd.AddCommand(updateNameCmd)
}
