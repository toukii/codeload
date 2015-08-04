package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	var read string
	fmt.Scanf("%s", &read)
	fmt.Println(read)
	_url := githubBranch(read)
	fmt.Printf("Downloading: %s\n", _url)
	// _url := "https://codeload.github.com/everfore/bconv/zip/master"
	b := get(_url)
	zipHttp("master.zip", b)
	unZip("master.zip")
}

func githubBranch(branch string) string {
	s1 := strings.Split(branch, ":")
	if len(s1) == 1 {
		s1 = append(s1, "master")
	}
	s2 := strings.Split(s1[0], "/")
	if len(s2) != 2 {
		panic("for example: everfore/bconv:master")
	}
	return fmt.Sprintf("https://codeload.github.com/%s/zip/%s", s1[0], s1[1])
}

func checkErr(err error) bool {
	if nil != err {
		fmt.Println(err)
		return true
	}
	return false
}

func get(_url string) []byte {
	resp, err := http.Get(_url)
	if checkErr(err) {
		return nil
	}
	b, err := ioutil.ReadAll(resp.Body)
	if checkErr(err) {
		return nil
	}
	return b
}

func zipHttp(filename string, b []byte) error {
	f, err := os.OpenFile(filename, os.O_CREATE, 0644)
	defer f.Close()
	if checkErr(err) {
		return err
	}
	_, err = f.Write(b)
	return err
}

func unZip(filename string) {
	r, err := zip.OpenReader(filename)
	defer r.Close()
	if checkErr(err) {
		return
	}
	for _, it := range r.File {
		name := it.Name
		n := strings.Index(name, "/")
		dstName := name[n+1:]
		if len(dstName) <= 0 {
			continue
		}
		if it.FileInfo().IsDir() {
			err := os.Mkdir(dstName, 0064)
			if checkErr(err) {
				continue
			}
			continue
		}
		dstF, err := os.Create(dstName)
		defer dstF.Close()
		if checkErr(err) {
			continue
		}
		src := new(bytes.Buffer)
		dst := new(bytes.Buffer)
		rc, err := it.Open()
		defer rc.Close()
		if checkErr(err) {
			continue
		}
		src.ReadFrom(rc)
		_, err = io.Copy(dst, src)
		if checkErr(err) {
			continue
		}
		dst.WriteTo(dstF)
	}
}
