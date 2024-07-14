package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"homework_4/proto"
	"strconv"
	"time"
)

var updateAmountCmd = &cobra.Command{
	Use:   "put-amount [name] [amount]",
	Short: "Deposit required amount into the account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		amount, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Println("Invalid amount: expected int64")
			return
		}

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

		rep, err := cli.UpdateAmount(ctx, &proto.UpdateAmountRequest{Name: name, Amount: amount})
		if err != nil {
			panic(err)
		}
		fmt.Println(rep.Message)
	},
}

func init() {
	rootCmd.AddCommand(updateAmountCmd)
}
