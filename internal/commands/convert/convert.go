package convert

import (
	"fmt"
	"vsimonari/gocconv/config"
	cmd "vsimonari/gocconv/internal/commands"
	"vsimonari/gocconv/internal/core"
	"vsimonari/gocconv/internal/storage"

	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert command",
	Run:   convertCurrenciesCmd,
}

var convertFromToken, convertToToken string
var amount float64

func init() {
	convertCmd.Flags().StringVarP(&convertFromToken, "from", "f", "USD", "Currency token to convert from")
	convertCmd.Flags().StringVarP(&convertToToken, "to", "t", "EUR", "Currency token to convert to")
	convertCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "Amount to convert")

	cmd.RootCmd.AddCommand(convertCmd)
}

var configuration *config.Configuration = config.Read()
var RatesStoragePath string = configuration.RatesStoragePath

func convertCurrenciesCmd(cmd *cobra.Command, args []string) {
	file := storage.OpenRatesFile(RatesStoragePath)

	currenciesRate := file.GetRate(convertFromToken, convertToToken)

	if currenciesRate == nil {
		rateInfo := core.GetRateInfo(core.CurrencyRate{From: convertFromToken, To: convertToToken})

		currenciesRate = rateInfo.GetRate()

		file.AddRate(*currenciesRate)
	}

	forward := (currenciesRate.From == convertFromToken) && (currenciesRate.To == convertToToken)

	result := 1.0

	if forward {
		result = amount * float64(currenciesRate.Rate)
	} else {
		result = amount * float64(currenciesRate.ReverseRate().Rate)
	}

	fmt.Printf("[+] Conversion from [%s] to [%s]: %.4f", currenciesRate.From, currenciesRate.To, result)
}
