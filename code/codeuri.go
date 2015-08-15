package code

import (
	"fmt"

	"bytes"
	"github.com/everfore/codeload/execc"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type CodeURI struct {
	User   string
	Repo   string
	Branch string
}

func NewCodeURI(user, repo, branch string) *CodeURI {
	return &CodeURI{
		User:   user,
		Repo:   repo,
		Branch: branch,
	}
}

func defaultUser() string {
	dir, _ := os.Getwd()
	base := filepath.Base(dir)
	fmt.Println(base)
	return base
}

func (c *CodeURI) Set(value string) error {
	urb := strings.Split(value, ":")
	ur := urb[0]
	if len(urb) <= 1 {
		c.Branch = "master"
	} else {
		c.Branch = urb[1]
	}

	urs := strings.Split(ur, "/")
	if len(urs) <= 1 {
		c.User = defaultUser()
		c.Repo = ur
	} else {
		c.User = urs[0]
		c.Repo = urs[1]
	}
	return nil
}

func (c *CodeURI) String() string {
	return fmt.Sprintf("%s/%s:%s", c.User, c.Repo, c.Branch)
}

func GithubCodeURI(branch string) *CodeURI {
	s1 := strings.Split(branch, ":")
	if len(s1) == 1 {
		s1 = append(s1, "master")
	}
	s2 := strings.Split(s1[0], "/")
	if len(s2) != 2 {
		panic("for example: everfore/bconv:master")
	}
	return NewCodeURI(s2[0], s2[1], s1[1])
}

func (c *CodeURI) URI() string {
	return fmt.Sprintf("https://codeload.github.com/%s/%s/zip/%s", c.User, c.Repo, c.Branch)
}

func (c *CodeURI) UnzipName() string {
	return fmt.Sprintf("%s-%s", c.Repo, c.Branch)
}

func (c *CodeURI) Download() {
	curled := c.curl()
	if curled {
		log.Println("curl success!")
	} else {
		log.Println("curl failed!")
	}
}

func (c *CodeURI) curl() bool {
	uri := c.URI()
	curl_command := fmt.Sprintf("curl %s", uri)
	b, ok := execc.ExecuteCmdHere(curl_command)
	if ok {
		f, err := os.OpenFile(fmt.Sprintf("%s.zip", c.Branch), os.O_WRONLY|os.O_CREATE, 0644)
		if checkerr(err) {
			return false
		}
		src := bytes.NewReader(b)
		_, err = io.Copy(f, src)
		if !checkerr(err) {
			return true
		}
	}
	return false
}

func (c *CodeURI) Unzip() bool {
	unzip_command := fmt.Sprintf("unzip %s.zip", c.Branch)
	_, ok := execc.ExecuteCmdHere(unzip_command)
	if ok {
		rename_command := fmt.Sprintf("mv %s %s", c.UnzipName(), c.Repo)
		_, renamed := execc.ExecuteCmdHere(rename_command)
		if renamed {
			del_command := fmt.Sprintf("rm -rf %s.zip", c.Branch)
			_, deled := execc.ExecuteCmdHere(del_command)
			if !deled {
				return false
			}
			return true
		}
	}
	return false
}

func checkerr(err error) bool {
	if nil != err {
		log.Println(err)
		return true
	}
	return false
}
