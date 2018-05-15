package reflect

import (
	"fmt"
	"reflect"
)

type Result struct {
}
type CallFunc func(a int, b string, c *Result) (string, error)

// 定义一个待反射的函数
func Function(a int, b string, c *Result) (string, error) {
	return "", nil
}

func reflectFunc(c CallFunc) {
	vTyp := reflect.TypeOf(c)
	fmt.Println(vTyp.Kind()) // 类型为func

	if vTyp.Kind() != reflect.Func {
		fmt.Println("not func type")
	}

	fmt.Println("传入参数数量: ", vTyp.NumIn())
	fmt.Println("传出参数数量: ", vTyp.NumOut())

	for i := 0; i < vTyp.NumIn(); i++ {
		in := vTyp.In(i) // Type
		fmt.Println("\tin:", in.Kind(), in)
	}

	for i := 0; i < vTyp.NumOut(); i++ {
		out := vTyp.Out(i)
		fmt.Println("\tout:", out.Kind(), out)
	}
}

func testFunc() {
	reflectFunc(Function)
}
