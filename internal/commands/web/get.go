package web

import (
	"fmt"
	"vsimonari/gocconv/internal/core"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Web subcommand to invoke data operations in web",
	Run:   getRateFromWebCmd,
}

var getFromToken, getToToken string

func init() {
	getCmd.Flags().StringVarP(&getFromToken, "from", "f", "EUR", "Token of currency to convert from")
	getCmd.Flags().StringVarP(&getToToken, "to", "t", "USD", "Token of currency to convert to")

	webCmd.AddCommand(getCmd)
}

func getRateFromWebCmd(cmd *cobra.Command, args []string) {
	info := core.GetRateInfo(getFromToken, getToToken)

	fmt.Printf("F: %s\n", info.From)
	fmt.Printf("T: %s\n", info.To)
	fmt.Printf("R: %.2f\n", info.Rate)
}
