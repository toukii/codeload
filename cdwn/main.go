package main

import (
	"fmt"
	"github.com/everfore/codeload/unzip"
	"github.com/shaalx/goutils"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	var input, user, repo, branch, input_1, target string
	tips := "[user/]repo[:branch]  > $"
	fmt.Print(tips)
	fmt.Scanf("%s", &input)

	start := time.Now()
	if strings.Contains(input, "/") {
		inputs := strings.Split(input, "/")
		user = inputs[0]
		input_1 = inputs[1]
	} else {
		pwd, _ := os.Getwd()
		user = filepath.Base(pwd)
		input_1 = input
	}

	if strings.Contains(input_1, ":") {
		input_1s := strings.Split(input_1, ":")
		repo = input_1s[0]
		branch = input_1s[1]
	} else {
		repo = input_1
		branch = "master"
	}
	fmt.Printf("%s/%s:%s\n", user, repo, branch)
	codeload_uri := fmt.Sprintf("https://codeload.github.com/%s/%s/zip/%s", user, repo, branch)
	resp, err := http.Get(codeload_uri)
	if goutils.CheckErr(err) {
		panic(fmt.Sprintf("GET:%s ERROR:%v\n", codeload_uri, err))
	}
	if resp == nil {
		panic("nil")
	}

	GOPATH := os.Getenv("GOPATH")
	target = filepath.Join(GOPATH, "src", "github.com", user, repo)
	unzip.UnzipReader(resp.Body, target)
	fmt.Printf("cost time:%v\n", time.Now().Sub(start))
}
