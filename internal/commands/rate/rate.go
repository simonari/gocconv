package rate

import (
	"vsimonari/gocconv/config"
	cmd "vsimonari/gocconv/internal/commands"

	"github.com/spf13/cobra"
)

var rateCmd = &cobra.Command{
	Use:   "rate",
	Short: "Rate subcommand to invoke operations on rates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	cmd.RootCmd.AddCommand(rateCmd)
}

var configuration *config.Configuration = config.Read()
var RatesStoragePath string = configuration.RatesStoragePath
