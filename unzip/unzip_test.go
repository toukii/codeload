package unzip

import (
	"github.com/toukii/goutils"
	"testing"

	"os"

	"net/http"
	"path/filepath"
)

func TestUnzipReader1(t *testing.T) {
	of, err := os.OpenFile("test.zip", os.O_RDONLY, 0644)
	defer of.Close()
	if goutils.CheckErr(err) {
		t.Log(err)
	}
	UnzipReader(of, "./test")
}

func TestUnzipReader2(t *testing.T) {
	resp, err := http.Get("https://codeload.github.com/toukii/tools/zip/env")
	if goutils.CheckErr(err) {
		return
	}
	gopath := os.Getenv("GOPATH")
	target := filepath.Join(gopath, "src", "github.com", "toukii", "tools")
	UnzipReader(resp.Body, target)
}
