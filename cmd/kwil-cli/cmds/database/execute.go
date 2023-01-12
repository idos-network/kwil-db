package database

import (
	"context"
	"fmt"
	"kwil/cmd/kwil-cli/util"
	"kwil/cmd/kwil-cli/util/display"
	"kwil/kwil/client"
	"kwil/x/transactions"
	"kwil/x/types/databases"
	"kwil/x/types/execution"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func executeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute",
		Short: "Execute a query",
		Long: `Execute executes a query against the specified database.  The query name is
		specified as the first argument, and the query a arguments are specified after.
		In order to specify an argument, you first need to specify the argument name.
		You then specify the argument type.

		For example, if I have a query name "create_user" that takes two arguments: name and age.
		I would specify the query as follows:

		create_user name satoshi age 32

		You specify the database to execute this against with the --database-name flag, and
		the owner with the --database-owner flag.

		You can also specify the database by passing the database id with the --database-id flag.

		For example:

		create_user name satoshi age 32 --database-name mydb --database-owner 0xAfFDC06cF34aFD7D5801A13d48C92AD39609901D

		OR

		create_user name satoshi age 32 --database-id x1234`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return util.ConnectKwil(cmd.Context(), viper.GetViper(), func(ctx context.Context, cc *grpc.ClientConn) error {
				client, err := client.NewClient(cc, viper.GetViper())
				if err != nil {
					return fmt.Errorf("failed to create client: %w", err)
				}

				// check that args is odd and has at least 3 elements
				if len(args) < 3 || len(args)%2 == 0 {
					return fmt.Errorf("invalid number of arguments")
				}

				// we will check if the user specified the database id or the database name and owner
				var executables []*execution.Executable

				// get the database id
				dbId, err := cmd.Flags().GetString("db_id")
				if err != nil || dbId == "" {
					// if we get an error, it means the user did not specify the database id
					// get the database name and owner
					dbName, err := cmd.Flags().GetString("db_name")
					if err != nil {
						return fmt.Errorf("either database id or database name and owner must be specified: %w", err)
					}

					dbOwner, err := cmd.Flags().GetString("db_owner")
					if err != nil {
						return fmt.Errorf("either database id or database name and owner must be specified: %w", err)
					}

					// create the dbid.  we will need this for the execution body
					dbId = databases.GenerateSchemaName(dbOwner, dbName)
				}

				executables, err = client.Txs.GetExecutablesById(ctx, dbId)
				if err != nil {
					return fmt.Errorf("failed to get executables: %w", err)
				}

				fmt.Println(len(executables))

				// get the query from the executables
				var query *execution.Executable
				for _, executable := range executables {
					if strings.EqualFold(executable.Name, args[0]) {
						query = executable
						break
					}
				}
				if query == nil {
					return fmt.Errorf("query %s not found", args[0])
				}

				// check that each input is provided
				userIns := make([]*execution.UserInput, 0)
				for _, input := range query.UserInputs {
					found := false
					for i := 1; i < len(args); i += 2 {
						if args[i] == input.Name {
							found = true
							userIns = append(userIns, &execution.UserInput{
								Name:  input.Name,
								Value: args[i+1],
							})
							break
						}
					}
					if !found {
						return fmt.Errorf("input %s not provided", input.Name)
					}
				}

				// create the execution body
				body := &execution.ExecutionBody{
					Database: dbId,
					Query:    query.Name,
					Caller:   client.Config.Address,
					Inputs:   userIns,
				}

				// buildtx
				tx, err := client.BuildTransaction(ctx, transactions.EXECUTE_QUERY, body, client.Config.PrivateKey)
				if err != nil {
					return fmt.Errorf("failed to build transaction: %w", err)
				}

				// broadcast
				res, err := client.Txs.Broadcast(ctx, tx)
				if err != nil {
					return fmt.Errorf("failed to broadcast transaction: %w", err)
				}

				// print the response
				display.PrintTxResponse(res)

				return nil
			})
		},
	}

	cmd.Flags().StringP("db_id", "i", "", "the database id")
	cmd.Flags().StringP("db_name", "n", "", "the database name")
	cmd.Flags().StringP("db_owner", "o", "", "the database owner")
	return cmd
}
