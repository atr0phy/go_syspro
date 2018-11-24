package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()
	file, err := zipWriter.Create("file.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(file, strings.NewReader("圧縮データ"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
