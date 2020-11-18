package main

import (
	"github.com/canc3s/gnc/internal/runner"
)

func main() {
	options := runner.ParseOptions()

	runner.Process(options)
}