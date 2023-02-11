package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)


var (
	passwordLength int
	noNumbers bool
	noSymbols bool
)

const (
	upperAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerAlpha = "abcdefghijklmnopqrstuvwxyz"
	digits = "0123456789"
	symbols = "`~!@#$%^&*()_-+={[}]|\\:;\"'<,>.?/"
)

var PasswordCmd = &cobra.Command{
	Use: "password",
	Short: "Create a simple password",
	Long: "Create a simple password",
	Run: createPassword,
}

func init() {
	flags := PasswordCmd.Flags()

	flags.IntVarP(&passwordLength, "length", "l", 8, "Length of the password")
	flags.BoolVarP(&noNumbers, "no-numbers", "N", false, "Do not include numbers in password")
	flags.BoolVarP(&noSymbols, "no-symbols", "S", false, "Do not include symbols in password")
}


func createPassword(cmd *cobra.Command, args []string)  {
	options := upperAlpha + lowerAlpha
	result := make([]byte, passwordLength)

	if !noNumbers {
		options = options + digits
	}

	if !noSymbols {
		options = options + symbols
	}

	rand.Seed(time.Now().Unix())
	for passwordLength > 0 {
		val := rand.Intn(len(options))
		result[passwordLength-1] = options[val]

		passwordLength -= 1
	}

	fmt.Printf("%s",result)
}