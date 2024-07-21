package rate

import (
	"fmt"
	"vsimonari/gocconv/internal/core"
	"vsimonari/gocconv/internal/storage"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add rate to storage",
	Run:   addExchangeRateCmd,
}

var addFromToken, addToToken string
var rate float32

func init() {
	addCmd.Flags().StringVarP(&addFromToken, "from", "f", "EUR", "Currency token to convert from")
	addCmd.Flags().StringVarP(&addToToken, "to", "t", "USD", "Currency token to convert to")
	addCmd.Flags().Float32VarP(&rate, "rate", "r", 1, "Conversion rate")

	rateCmd.AddCommand(addCmd)
}

func addExchangeRateCmd(cmd *cobra.Command, args []string) {
	if addFromToken == addToToken {
		fmt.Printf("[+] Equal tokens provided. Operation will not be performed")
		return
	}

	file := storage.OpenRatesFile(RatesStoragePath)

	c := core.CurrencyRate{
		From: addFromToken,
		To:   addToToken,
		Rate: rate,
	}

	file.AddRate(c)

	fmt.Printf("[+] Rate added. Now file contains [%v] rates\n", file.Stored)
}
