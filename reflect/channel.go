package reflect

import (
	"fmt"
	"reflect"
)

func chanref(v interface{}) {
	typ := reflect.TypeOf(v)
	//val := reflect.ValueOf(v)

	// type: reflect.Chan
	if typ.Kind() != reflect.Chan {
		fmt.Println("not chan type")
		return
	}
	fmt.Println("type:", typ.Kind())
	//
	fmt.Println("chan dir:", typ.ChanDir())

	// chan type
	fmt.Println("chan type:", typ.Elem())
}

type event struct {
	ch  reflect.SelectCase
	typ reflect.Type
}

func (e *event) sub(c interface{}) {
	cVal := reflect.ValueOf(c)
	cTyp := cVal.Type()

	if cTyp.Kind() != reflect.Chan {
		fmt.Println("not chan type")
		return
	}

	// 定义一个事件
	cas := reflect.SelectCase{Dir: reflect.SelectSend, Chan: cVal}
	e.ch = cas
	e.typ = cTyp.Elem() // 数据类型
}

func (e *event) send(data interface{}) {
	rval := reflect.ValueOf(data)
	rtyp := rval.Type()

	if rtyp.Kind() != e.typ.Kind() {
		fmt.Println("type not match:", rtyp.Kind(), e.typ.Kind())
		return
	}

	if e.ch.Chan.TrySend(rval) {
		fmt.Println("send ok")
	} else {
		fmt.Println("send fail")
	}
}

func testChan() {
	c := make(chan int)
	// chanref(c)
	e := new(event)
	e.sub(c)
	go e.send(23)
	fmt.Println("recv: ", <-c)
}
