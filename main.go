package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/everfore/codeload/code"
	// "github.com/everfore/codeload/unzip"
	"flag"
)

var repo code.CodeURI

func init() {
	flag.Var(&repo, "c", "usage")
}
func main() {
	flag.Parse()

	start := time.Now()

	if len(repo.String()) <= 2 {
		var input string
		tips := "everfore/codeload:master  > "
		fmt.Print(tips)
		fmt.Scanf("%s", &input)
		if strings.Contains(input, "/") {
			codeuri := code.GithubCodeURI(input)
			codeuri.Download()
			codeuri.Unzip()
		} else {
			codeuri := code.GithubCodeURI("/" + input)
			codeuri.Unzip()
		}
	} else {
		repo.Download()
		repo.Unzip()
	}

	fmt.Printf("cost time:%v\n", time.Now().Sub(start))
}
