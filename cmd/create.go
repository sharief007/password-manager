package cmd

import (
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use: "create",
}

func init() {
	CreateCmd.AddCommand(PasswordCmd)
	CreateCmd.AddCommand(AdvPasswordCmd)
}