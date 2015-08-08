package unzip

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

func Unzip(filename string) {
	r, err := zip.OpenReader(filename)
	defer r.Close()
	if checkerr(err) {
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
