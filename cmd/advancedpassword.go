package cmd

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	advPasswordLength int
	minUppercase      int
	minLowerCase      int
	minDigits         int
	minSymbols        int
)

var AdvPasswordCmd = &cobra.Command{
	Use: "adv-password",
	Short: "Create a password with more customization",
	Long: "Create a password with more customization",
	Run: createAdvancedPassword,
}

func init() {
	flags := AdvPasswordCmd.Flags()

	flags.IntVarP(&advPasswordLength, "length", "l", 8, "Length of the password")
	flags.IntVarP(&minUppercase, "min-upper", "U", 1, "Minimum number of uppercase letters")
	flags.IntVarP(&minLowerCase, "min-lower", "L", 1, "Minimum number of lowercase letters")
	flags.IntVarP(&minDigits, "min-digits", "D", 1, "Minimum number of digits")
	flags.IntVarP(&minSymbols, "min-symbols", "S", 1, "Minimum number of Symbols")
}

func createAdvancedPassword(cmd *cobra.Command, args []string) {
	options := upperAlpha + lowerAlpha + digits + symbols

	if advPasswordLength < (minUppercase + minLowerCase + minDigits + minSymbols) {
		fmt.Println(errors.New("Invalid Input"))
		os.Exit(1)
	} 
	
	result := make([]byte, advPasswordLength + 1)
	index := 0
	rand.Seed(time.Now().Unix())

	for minUppercase > 0 {
		val := rand.Intn(26)
		result[index] = upperAlpha[val]

		index += 1
		minUppercase -= 1
	}

	for minLowerCase > 0 {
		val := rand.Intn(26)
		result[index] = lowerAlpha[val]

		index += 1
		minLowerCase -= 1
	}

	for minDigits > 0 {
		val := rand.Intn(11)
		result[index] = digits[val]

		index += 1
		minDigits -= 1
	}

	for minSymbols > 0 {
		val := rand.Intn(len(symbols))
		result[index] = symbols[val]

		index += 1
		minSymbols -= 1
	}

	for index < advPasswordLength {
		val := rand.Intn(len(options))
		result[index] = options[val]

		index += 1
	}

	for i := range result {
		j := rand.Intn(i + 1)
		result[i], result[j] = result[j], result[i]
	}

	fmt.Printf("%s", result)
}