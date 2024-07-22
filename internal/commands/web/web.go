package web

import (
	"vsimonari/gocconv/config"
	cmd "vsimonari/gocconv/internal/commands"

	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Web subcommand to invoke data operations in web",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	cmd.RootCmd.AddCommand(webCmd)
}

var configuration *config.Configuration = config.Read()
var RatesStoragePath string = configuration.RatesStoragePath
