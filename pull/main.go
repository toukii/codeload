package main

import (
	"fmt"
	"github.com/everfore/exc"
	// "github.com/shaalx/goutils"
	"flag"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	read    = false // default writeable
	install = false
)

func init() {
	flag.BoolVar(&read, "r", false, "-r [true] : git@github.com[false] or git://github.com[true]")
	flag.BoolVar(&install, "i", false, "-i [true] : go install")
}

func main() {
	flag.Parse()
	var input, user, repo, branch, input_1 /*,target*/ string
	tips := "[user/]repo[:branch]  > $"

	fmt.Print(tips)
	fmt.Scanf("%s", &input)

	start := time.Now()
	if strings.Contains(input, "/") {
		inputs := strings.Split(input, "/")
		user = inputs[0]
		input_1 = inputs[1]
	} else if len(input) <= 0 {
		exc.NewCMD("git pull").Debug().Execute()
		return
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
	codeload_uri := ""
	if !read {
		codeload_uri = fmt.Sprintf("git clone --depth 1 git@github.com:%s/%s.git", user, repo)
	} else {
		codeload_uri = fmt.Sprintf("git clone --depth 1 git://github.com/%s/%s", user, repo)
	}
	GOPATH := os.Getenv("GOPATH")
	target := filepath.Join(GOPATH, "src", "github.com", user, repo)
	os.MkdirAll(target, 0777)
	cmd := exc.NewCMD(codeload_uri).Env("GOPATH").Cd("src/github.com/").Cd(user).Wd().Debug().Execute()
	if install {
		cmd.Cd(repo).Wd().Reset("go install").Execute()
	}

	fmt.Printf("cost time:%v\n", time.Now().Sub(start))
}
