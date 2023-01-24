package fund

import (
	"context"
	"fmt"
	"kwil/cmd/kwil-cli/common"
	grpc_client "kwil/kwil/client/grpc-client"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func getAccountCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "account",
		Short: "Gets account balance, spent, and nonce information",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return common.DialGrpc(cmd.Context(), viper.GetViper(), func(ctx context.Context, cc *grpc.ClientConn) error {
				client, err := grpc_client.NewClient(cc, viper.GetViper())
				if err != nil {
					return fmt.Errorf("error creating client: %w", err)
				}

				// check if account is set
				account, err := cmd.Flags().GetString("account")
				if err != nil {
					return fmt.Errorf("error getting account flag: %w", err)
				}

				if account == "" {
					account = client.Config.Address
				}

				acc, err := client.Accounts.GetAccount(ctx, account)
				if err != nil {
					return fmt.Errorf("error getting account: %w", err)
				}

				fmt.Println("Address: ", acc.Address)
				fmt.Println("Balance: ", acc.Balance)
				fmt.Println("Spent:   ", acc.Spent)
				fmt.Println("Nonce:   ", acc.Nonce)

				return nil
			},
			)
		},
	}

	cmd.Flags().StringP("account", "a", "", "Account address to get information for")

	return cmd
}
