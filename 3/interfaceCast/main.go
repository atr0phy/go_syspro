package main

import (
	"io"
	"io/ioutil"
	"strings"
)

func main() {
	var reader io.Reader = strings.NewReader("テストデータ")
	var readCloser io.ReadCloser = ioutil.NopCloser(reader)
	readCloser.Close()
}
