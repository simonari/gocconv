package main

import (
	cmd "vsimonari/gocconv/internal/commands"

	_ "vsimonari/gocconv/internal/commands/convert"
	_ "vsimonari/gocconv/internal/commands/rate"
)

func main() {
	cmd.Execute()
}
