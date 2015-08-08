package main

import (
	"fmt"
	"time"

	"github.com/everfore/codeload/code"
	"github.com/everfore/codeload/unzip"
)

func main() {
	var input string
	tips := "everfore/codeload:master  > "
	fmt.Print(tips)
	fmt.Scanf("%s", &input)
	fmt.Println(input)

	start := time.Now()

	codeuri := code.GithubCodeURI(input)
	codeuri.Download()
	codeuri.Unzip()

	fmt.Printf("cost time:%v\n", time.Now().Sub(start))
}

func v1_2() {
	start := time.Now()

	unzip.Unzip("master.zip")

	fmt.Printf("cost time:%v\n", time.Now().Sub(start))
}
