package main

import (
	"fmt"
	"os"

	"github.com/nokamoto/grpc-tryout/internal/protogen"
)

func main() {
	if err := protogen.NewOption().Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
