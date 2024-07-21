package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gocconv",
	Short: "A brief description of your application",
}

func Execute() {
	err := RootCmd.Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
