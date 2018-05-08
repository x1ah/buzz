package main

import (
	"flag"
	"fmt"
	"github.com/x1ah/buzz/src"
)

var (
	keyword string
	detail  string
	list    bool
	help    bool
)

func init() {
	flag.BoolVar(&list, "l", false, "Show local buzzwords.")
	flag.StringVar(&keyword, "k", "", "Buzzword keyword.")
	flag.StringVar(&detail, "d", "", "Buzzword detail.")
}

func main() {
	flag.Parse()

	if list {
		buzz.ShowListBuzzwords()
		return
	}

	if detail == "" {
		fmt.Println("Detail required.")
		return
	}

	if keyword == "" {
		fmt.Println("Keyword required.")
		return
	}

	buzzword := buzz.BuzzWord{
		Keyword: keyword,
		Detail:  detail,
	}
	_, err := buzz.AppendBuzzword(buzzword)
	if err != nil {
		panic(err)
	}

	fmt.Println("✔︎")
}
