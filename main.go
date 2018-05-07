package main

import (
	"fmt"
	"github.com/x1ah/buzz/src"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(`
		Usage:
			> buzz [keyword] [detail]

		- keyword: Keyword of this buzzword
		- detail: Detail of the buzzword`)
		os.Exit(0)
	}

	buzzword := buzz.BuzzWord{
		Keyword: os.Args[1],
		Detail:  os.Args[2],
	}

	_, err := buzz.AppendBuzzword(buzzword)
	if err != nil {
		panic(err)
	}

	fmt.Println("✔︎")
}
