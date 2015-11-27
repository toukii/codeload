package unzip

import (
	"archive/zip"
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func UnzipReader(r io.Reader, target string) {
	reader := bufio.NewReader(r)
	src := new(bytes.Buffer)
	n, err := src.ReadFrom(reader)
	if checkerr(err) || n <= 0 {
		return
	}
	of, err := os.OpenFile("github.zip", os.O_CREATE|os.O_WRONLY, 0666)
	defer of.Close()
	defer func() {
		os.Remove("github.zip")
	}()
	if checkerr(err) {
		return
	}
	dst := new(bytes.Buffer)

	src.ReadFrom(reader)
	n, err = io.Copy(dst, src)
	if checkerr(err) || n <= 0 {
		return
	}
	n, err = dst.WriteTo(of)
	if checkerr(err) || n <= 0 {
		return
	}

	Unzip("github.zip", target)
}

func Unzip(filename, target string) {
	r, err := zip.OpenReader(filename)
	defer r.Close()
	if checkerr(err) {
		return
	}
	os.MkdirAll(target, 0777)
	for _, it := range r.File {
		name := it.Name
		n := strings.Index(name, "/")
		dstName := name[n+1:]
		if len(dstName) <= 0 {
			continue
		}
		dstName = filepath.Join(target, dstName)
		if it.FileInfo().IsDir() {
			err := os.MkdirAll(dstName, 0777)
			if checkerr(err) {
				continue
			}
			continue
		}
		dstF, err := os.Create(dstName)
		defer dstF.Close()
		if checkerr(err) {
			continue
		}
		src := new(bytes.Buffer)
		dst := new(bytes.Buffer)
		rc, err := it.Open()
		defer rc.Close()
		if checkerr(err) {
			continue
		}
		src.ReadFrom(rc)
		_, err = io.Copy(dst, src)
		if checkerr(err) {
			continue
		}
		dst.WriteTo(dstF)
	}
}

func checkerr(err error) bool {
	if nil != err {
		log.Println(err)
		return true
	}
	return false
}
