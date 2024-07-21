package rate

import (
	"fmt"
	"vsimonari/gocconv/internal/storage"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete rate from storage",
	Run:   deleteRateCmd,
}

var deleteFromToken, deleteToToken string

func init() {
	deleteCmd.Flags().StringVarP(&deleteFromToken, "from", "f", "EUR", "Currency token to convert from")
	deleteCmd.Flags().StringVarP(&deleteToToken, "to", "t", "USD", "Currency token to convert to")

	rateCmd.AddCommand(deleteCmd)
}

func deleteRateCmd(cmd *cobra.Command, args []string) {
	if addFromToken == addToToken {
		fmt.Printf("[!] Equal tokens provided. Operation will not be performed")
		return
	}

	path := ".storage/rates.json"

	file := storage.OpenRatesFile(path)

	file.DeleteRate(deleteFromToken, deleteToToken)

	fmt.Printf("[+] Rate deleted. Now file contains [%v] rates\n", file.Stored)
}
