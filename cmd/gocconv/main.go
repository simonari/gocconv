package main

import (
	cmd "vsimonari/gocconv/internal/commands"

	_ "vsimonari/gocconv/internal/commands/convert"
	_ "vsimonari/gocconv/internal/commands/rate"
	_ "vsimonari/gocconv/internal/commands/web"
)

func main() {
	cmd.Execute()
}
