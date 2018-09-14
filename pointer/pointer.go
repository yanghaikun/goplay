package pointer

import (
	"fmt"
	"github.com/yanghaikun/goplay/container"
)

var test container.IntHeap
type MyPoint struct {
	X int
	Y int
}

func printFuncValue(p MyPoint){
	p.X = 1
	p.Y = 1
	fmt.Printf(" -> %v", p)
}

func printFuncPointer(pp *MyPoint){
	pp.X = 1 // 实际上应该写做 (*pp).X，Golang 给了语法糖，减少了麻烦，但是也导致了 * 的不一致
	pp.Y = 1
	fmt.Printf(" -> %v", pp)
}

func printChangeLocation(pp *MyPoint){
	fmt.Printf("\n location of pp is %p", &pp)
	pp = &MyPoint{}
	pp.X = 1 // 实际上应该写做 (*pp).X，Golang 给了语法糖，减少了麻烦，但是也导致了 * 的不一致
	pp.Y = 1
	fmt.Printf("\n location of pp is %p", &pp)
	fmt.Printf(" -> %v", pp)
}



func (p MyPoint) printMethodValue(){
	p.X += 1
	p.Y += 1
	fmt.Printf(" -> %v", p)
}

// 建议使用指针作为方法（method：printMethodPointer）的接收者（receiver：*MyPoint），一是可以修改接收者的值，二是可以避免大对象的复制
func (pp *MyPoint) printMethodPointer(){
	pp.X += 1
	pp.Y += 1
	fmt.Printf(" -> %v", pp)
}

func testChangeMapValue(value map[string]int) {
	value["0"] = value["0"] + 1
	value["1"] = value["1"] + 1
	fmt.Printf(" -> %v", value)
}

func testChangeMapLocation(value map[string]int) {
	fmt.Printf("\n location of value is %p", &value)
	value = make(map[string]int, 1)
	value["0"] = 999
	fmt.Printf("\n location of value is %p", &value)
	fmt.Printf(" -> %v", value)
}

func Pointer(){
	p := MyPoint{0, 0}
	pp := &MyPoint{0, 0}

	fmt.Printf("\n value to func(value): %v", p)
	printFuncValue(p)
	fmt.Printf(" --> %v", p)
	// Output: value to func(value): {0 0} -> {1 1} --> {0 0}

	//printFuncValue(pp) // cannot use pp (type *MyPoint) as type MyPoint in argument to printFuncValue

	//printFuncPointer(p) // cannot use p (type MyPoint) as type *MyPoint in argument to printFuncPointer

	fmt.Printf("\n pointer to func(pointer): %v", pp)
	printFuncPointer(pp)
	fmt.Printf(" --> %v", pp)
	// Output: pointer to func(pointer): &{0 0} -> &{1 1} --> &{1 1}

	fmt.Printf("\n value to method(value): %v", p)
	p.printMethodValue()
	fmt.Printf(" --> %v", p)
	// Output: value to method(value): {0 0} -> {1 1} --> {0 0}

	fmt.Printf("\n value to method(pointer): %v", p)
	p.printMethodPointer()
	fmt.Printf(" --> %v", p)
	// Output: value to method(pointer): {0 0} -> &{1 1} --> {1 1}

	fmt.Printf("\n pointer to method(value): %v", pp)
	pp.printMethodValue()
	fmt.Printf(" --> %v", pp)
	// Output: pointer to method(value): &{1 1} -> {2 2} --> &{1 1}

	fmt.Printf("\n pointer to method(pointer): %v", pp)
	pp.printMethodPointer()
	fmt.Printf(" --> %v", pp)
	// Output: pointer to method(pointer): &{1 1} -> &{2 2} --> &{2 2}

	fmt.Printf("\n test change Location: %v, location is %p", pp, pp)
	printChangeLocation(pp)
	fmt.Printf(" --> %v, location is %p", pp, pp)

	value := map[string]int{"0" : 0, "1" : 1}
	fmt.Printf("\n testChangeMapValue: %v", value)
    testChangeMapValue(value)
	fmt.Printf("testChangeMapValue: %v", value)

	fmt.Printf("\n testChangeMapLocation: %v, location is %p", value, &value)
	testChangeMapLocation(value)
	fmt.Printf("\n testChangeMapLocation: %v, location is %p", value, &value)

	value = make(map[string]int, 1)
	fmt.Printf("\nlocation of value is %p", &value)
	value["0"] = 0
	a := value["0"]
	fmt.Printf("\nlocation of value 0 is %p", &a)
	value["1"] = 0
	value["2"] = 0
	value["3"] = 0
	fmt.Printf("\nlocation of value is %p", &value)
	b := value["0"]
	c := value["0"]
	fmt.Printf("\nlocation of value 0 is %p", &b)
	fmt.Printf("\nlocation of value 0 is %p", &c)
	fmt.Printf("\nlocation of value is %p", &value)


	myValue := make(map[string]*MyPoint, 1)
	myValue["0"] = &MyPoint{0, 0}
	fmt.Printf("\n myValue is %v", *myValue["0"])
	va := myValue["0"]
	va.X = 999
	va.Y = 999
	myValue["1"] = &MyPoint{1, 1}
	myValue["2"] = &MyPoint{2, 2}
	myValue["3"] = &MyPoint{3, 3}
	fmt.Printf("\n myValue is %v", *myValue["0"])

	var tmp MyPoint
	p1 := &tmp
	fmt.Printf("\nlocation of tmp is %p", p1)
	tmp = MyPoint{1, 1}
	fmt.Printf("\nlocation of tmp is %p", &tmp)
}
