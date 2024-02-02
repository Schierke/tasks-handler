package main

import (
	"github.com/Schierke/tasks-handler/cmd"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	cmd.Execute()
}
