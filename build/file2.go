// +build darwin

package build

import "fmt"

func Add(a, b int) int {
	fmt.Println("file-002 a+b")
	return a + b
}
