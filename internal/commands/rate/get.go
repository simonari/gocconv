package rate

import (
	"fmt"
	"vsimonari/gocconv/internal/storage"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get rate from storage",
	Run:   getExchangeRateCmd,
}

var getFromToken, getToToken string

func init() {
	getCmd.Flags().StringVarP(&getFromToken, "from", "f", "EUR", "Currency token to convert from")
	getCmd.Flags().StringVarP(&getToToken, "to", "t", "USD", "Currency token to convert to")

	rateCmd.AddCommand(getCmd)
}

func getExchangeRateCmd(cmd *cobra.Command, args []string) {
	file := storage.OpenRatesFile(RatesStoragePath)

	rate := file.GetRate(getFromToken, getToToken)

	if rate == nil {
		fmt.Printf("[+] Rate from %v to %v was not found!", getFromToken, getToToken)
		return
	}

	fmt.Printf("[+] %.2f", rate.Rate)
}
