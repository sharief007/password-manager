package main

import (
	"context"
	"fmt"
	"os"
	"github.com/sharief007/password-manager/cmd"
	"github.com/spf13/cobra"
)

func main() {
	ctx := context.Background()

	root := &cobra.Command{
		Use: "passm",
		Version: "0.0.1",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	root.AddCommand(cmd.CreateCmd)
	if err := root.ExecuteContext(ctx); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
