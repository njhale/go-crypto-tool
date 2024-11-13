package main

import (
	"fmt"
	"os"

	"github.com/otto8-ai/go-hash-tool/commands"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: gptscript-go-tool <command>")
		os.Exit(1)
	}

	var (
		err error
		res string
	)
	switch cmd := os.Args[1]; cmd {
	case "hash":
		res, err = commands.Hash(os.Getenv("DATA"), os.Getenv("ALGO"))
	default:
		err = fmt.Errorf("Unsupported command: %s", cmd)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if res != "" {
		fmt.Println(res)
	}
}
