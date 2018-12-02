package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("ユーザーID: %d", os.Getuid())
	fmt.Printf("グループID: %d", os.Getgid())
	groups, _ := os.Getgroups()
	fmt.Printf("サブグループID: %v", groups)
}
