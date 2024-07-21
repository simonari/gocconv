package rate

import (
	"fmt"
	"vsimonari/gocconv/internal/storage"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update rate in storage",
	Run:   updateRateCmd,
}

var updateFromToken, updateToToken string
var updateRate float32

func init() {
	updateCmd.Flags().StringVarP(&updateFromToken, "from", "f", "EUR", "Currency token to convert from")
	updateCmd.Flags().StringVarP(&updateToToken, "to", "t", "USD", "Currency token to convert to")
	updateCmd.Flags().Float32VarP(&updateRate, "rate", "r", 1, "Conversion rate")

	rateCmd.AddCommand(updateCmd)
}

func updateRateCmd(cmd *cobra.Command, args []string) {
	path := ".storage/rates.json"

	file := storage.OpenRatesFile(path)

	file.UpdateRate(updateFromToken, updateToToken, updateRate)

	fmt.Printf("[+] Rate updated")
}
