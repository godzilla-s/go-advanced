// +build linux

package build

import "fmt"

func Add(a, b int) int {
	fmt.Println("file -- 001 a + b")
	return a + b
}
