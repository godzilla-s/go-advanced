package reflect

import (
	"go-advanced/run"
)

func init() {
	run.Register("refelct", Run)
}

func Run() {
	// test()
	testChan()
}
