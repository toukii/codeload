package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/everfore/codeload/code"
	// "github.com/everfore/codeload/unzip"
)

func main() {
	var input string
	tips := "everfore/codeload:master  > "
	fmt.Print(tips)
	fmt.Scanf("%s", &input)

	start := time.Now()

	if strings.Contains(input, "/") {
		codeuri := code.GithubCodeURI(input)
		codeuri.Download()
		codeuri.Unzip()
	} else {
		codeuri := code.GithubCodeURI("/" + input)
		codeuri.Unzip()
	}

	fmt.Printf("cost time:%v\n", time.Now().Sub(start))
}
