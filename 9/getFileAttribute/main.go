package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]", os.Args[0])
		os.Exit(1)
	}
	info, err := os.Stat(os.Args[1])
	if err == os.ErrNotExist {
		fmt.Printf("file not found: %s\n", os.Args[1])
	} else if err != nil {
		panic(err)
	}
	fmt.Println("FileInfo")
	fmt.Printf("\tファイル名: %v\n", info.Name())
	fmt.Printf("\tサイズ: %v\n", info.Size())
	fmt.Printf("\t変更日時: %v\n", info.ModTime())
	fmt.Println("Mode()")
	fmt.Printf("\tディレクトリ? %v\n", info.Mode().IsDir())
	fmt.Printf("\t読み書き可能な通常ファイル? %v\n", info.Mode().IsRegular())
	fmt.Printf("\tUnixのファイルアクセス権限ビット %o\n", info.Mode().Perm())
	fmt.Printf("\tモードのテキスト表現 %v\n", info.Mode().String())
}
