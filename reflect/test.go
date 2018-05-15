package reflect

import (
	"go-advanced/run"
)

func init() {
	run.Register("reflect", Run)
}

func Run() {
	// test()
	//testChan()

	testFunc()

	testInterface()
}
