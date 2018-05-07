package reflect

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Id   string
	bak  []byte
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Person) Set(name string, age int, id string, bak []byte) {
	p.Name = name
	p.Age = age
	p.Id = id
	p.bak = bak
}

func (p *Person) Get(id string) (string, int, []byte) {
	return "", 0, []byte("")
}

func test() {
	p := new(Person)

	pTyp := reflect.TypeOf(p)
	fmt.Println("type:", pTyp)
	pVal := reflect.ValueOf(p)

	// 取对象类型
	indType := reflect.Indirect(pVal).Type()
	fmt.Println("indirect:", indType)

	for i := 0; i < pTyp.NumMethod(); i++ {
		// 获取方法
		method := pTyp.Method(i)
		fmt.Println("method:", method.Func, method.Name, method.Type, method.Index)

		mtyp := method.Type
		// 传入参数
		fmt.Println("\tparams in num:", mtyp.NumIn())
		for j := 0; j < mtyp.NumIn(); j++ {
			in := mtyp.In(j) // Type
			fmt.Println("\t\ttype in:", in.Kind(), in)
		}

		fmt.Println("\tparams out num:", mtyp.NumOut())
		for j := 0; j < mtyp.NumOut(); j++ {
			out := mtyp.Out(j) // Type
			fmt.Println("\t\ttype out:", out.Kind(), out, out)
		}
	}
}
