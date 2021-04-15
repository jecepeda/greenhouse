package main

import (
	"github.com/jecepeda/greenhouse/server/cmd"
	_ "github.com/lib/pq"
)

func main() {
	cmd.Execute()
}
