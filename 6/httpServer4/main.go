package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

var contents = []string{
	"テスト文字列１テスト文字列１テスト文字列１テスト文字列１テスト文字列１テスト文字列１",
	"テスト文字列2テスト文字列2テスト文字列2テスト文字列2テスト文字列2テスト文字列2テスト文字列2",
	"テスト文字列3テスト文字列3テスト文字列3テスト文字列3テスト文字列3テスト文字列3テスト文字列3テスト文字列3",
	"テスト文字列4テスト文字列4テスト文字列4テスト文字列4テスト文字列4テスト文字列4テスト文字列4テスト文字列4テスト文字列4",
	"テスト文字列5テスト文字列5テスト文字列5テスト文字列5テスト文字列5テスト文字列5テスト文字列5テスト文字列5",
	"テスト文字列6テスト文字列6テスト文字列6テスト文字列6テスト文字列6テスト文字列6テスト文字列6",
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()
	for {
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))
		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}
		fmt.Fprintf(conn, "0\r\n\r\n")
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}
}
