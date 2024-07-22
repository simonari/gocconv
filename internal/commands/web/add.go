package web

import (
	"fmt"
	"vsimonari/gocconv/internal/core"
	"vsimonari/gocconv/internal/storage"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Web subcommand to invoke data operations in web",
	Run:   addRateFromWebCmd,
}

var addFromToken, addToToken string

func init() {
	addCmd.Flags().StringVarP(&addFromToken, "from", "f", "EUR", "Token of currency to convert from")
	addCmd.Flags().StringVarP(&addToToken, "to", "t", "USD", "Token of currency to convert to")

	webCmd.AddCommand(addCmd)
}

func addRateFromWebCmd(cmd *cobra.Command, args []string) {
	c := core.CurrencyRate{From: addFromToken, To: addToToken}

	info := core.GetRateInfo(c)

	c.Rate = float32(info.Rate)

	file := storage.OpenRatesFile(RatesStoragePath)

	file.AddRate(c)

	fmt.Printf("[+] Rate added. Now file contains [%v] rates\n", file.Stored)
}
