package main

import (
	"fmt"
	"github.com/everfore/codeload/code"
	"os"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	var input string
	tips := "[everfore/]codeload[:master]  > $"
	fmt.Print(tips)
	fmt.Scanf("%s", &input)
	if strings.Contains(input, "/") {
		codeuri := code.GithubCodeURI(input)
		os.Chdir(codeuri.GithuUserPath())
		codeuri.Download()
		codeuri.Unzip()
	} else {
		var repo code.CodeURI
		repo.Set(input)
		repo.Download()
		repo.Unzip()
	}

	fmt.Printf("cost time:%v\n", time.Now().Sub(start))
}
