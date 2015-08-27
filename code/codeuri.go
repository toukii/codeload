package code

import (
	"fmt"

	"bytes"
	"github.com/everfore/exc"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"net/http"
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
	fmt.Printf("default user:%s\n", base)
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
	if !strings.Contains(ur, "/") {
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

func (c *CodeURI) GithuUserPath() string {
	path := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", c.User)
	_, err := os.Stat(path)
	if err != nil {
		err = os.MkdirAll(path, 0777)
		if checkerr(err) {
			return path
		}
	}
	return path
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
	// curled := c.curl()
	curled := c.download()
	if curled {
		log.Println("download success!")
	} else {
		log.Println("download failed!")
	}
}

func (c *CodeURI) download() bool {
	uri := c.URI()
	resp, err := http.Get(uri)
	fmt.Printf("downloading...  %s\n", uri)
	if checkerr(err) {
		return false
	}
	f, err := os.OpenFile(fmt.Sprintf("%s.zip", c.Branch), os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if checkerr(err) {
		return false
	}
	_, err = io.Copy(f, resp.Body)
	if checkerr(err) {
		return false
	}
	return true
}

func (c *CodeURI) curl() bool {
	uri := c.URI()
	curl_command := fmt.Sprintf("curl %s\n", uri)
	b, err := exc.NewCMD(curl_command).Do()
	if !exc.Checkerr(err) {
		f, err := os.OpenFile(fmt.Sprintf("%s.zip", c.Branch), os.O_WRONLY|os.O_CREATE, 0644)
		defer f.Close()
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
	cmd := exc.NewCMD(unzip_command)
	_, err := cmd.Debug().Do()
	if !exc.Checkerr(err) {
		rename_command := fmt.Sprintf("mv %s %s", c.UnzipName(), c.Repo)
		_, err = os.Stat(c.Repo)
		if nil == err {
			err = os.RemoveAll(c.Repo)
			checkerr(err)
		}
		_, renamed := cmd.Reset(rename_command).Do()
		if !exc.Checkerr(renamed) {
			zipfile := fmt.Sprintf("%s.zip", c.Branch)
			removed := os.Remove(zipfile)
			if !checkerr(removed) {
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
