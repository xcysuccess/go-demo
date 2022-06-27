package main //1

import (
	"errors"
	"fmt"
	"reflect"
	"time"
	"unsafe"
) //2
//import "unsafe"

func main() { //3
	//var a, b int = 1, 2
	//fmt.Println("a,b", a, b)
	fmt.Println("Hello World") //4
	fmt.Println()
	//testNumber()
	//testConst()
	//testEnum()
	//testArray()
	//testPointer()
	//testStruct()
	//testSlice()
	//testRange()
	//testMap()
	//testFiboacci()
	//testInterface()
	//testError()
	//testThread()
	//testUnSafe()
	//testUnsafeOffset()
	//testChannel()
	//testPanic()
	test
}

func testNumber() {
	var i int
	var f float64
	var b bool
	var s string
	//%q字符串活切换的无解译字节
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var d = true
	fmt.Println("d", d)

	var intVal int
	//intVal := 1
	fValue := 1.5
	fmt.Println("intVal", intVal)
	fmt.Println("fValue", fValue)

	var c1, d1 int = 1, 2
	var e1, f1 = 123, "hello"
	fmt.Println("c1, d1, e1, f1", c1, d1, e1, f1)

	fmt.Println("reflect")
	fmt.Println("fvalue author:", reflect.TypeOf(fValue))
	fmt.Println("fvalue author:", reflect.ValueOf(fValue))

}

func testConst() {
	const LENGTH int = 10
	const WIDTH int = 5

	var area int
	const a, b, c = 1, false, "str"
	var f float64
	f = 54.0
	area = LENGTH * WIDTH
	fmt.Println("面积为 :", area)
	fmt.Println("f:", f)
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)
}

func testEnum() {
	const (
		Unknown = iota
		Female
		Male
	)
	fmt.Println("Male:", Male)

	var a int = 100
	var b int = 200
	var result = max(a, b)
	a1, b1 := swap("x1", "y1")
	fmt.Println("result:", result)
	fmt.Println("a1,b1 = ", a1, b1)
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func swap(x, y string) (string, string) {
	return y, x
}

func testArray() {
	var array [10]int
	var i, j, k int
	for i = 0; i < 10; i++ {
		array[i] = i + 1
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = [%d]\n", j, array[j])
	}

	balance := []float32{1000, 7.0, 5.0, 2.0, 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("num:%d value:%f \n", k, balance[k])
	}
}

func testPointer() {
	a := 10
	var ip *int
	ip = &a
	fmt.Printf("a变量的地址是:%x \n", &a)
	fmt.Printf("ip变量存储的地址是:%x \n", ip)
	fmt.Println("ip指针的值是", *ip)
}

func testStruct() {
	type Books struct {
		title   string
		author  string
		subject string
		book_id int
	}
	var book Books = Books{title: "hello", author: "tom", subject: "主题", book_id: 30513207}
	var bookPtr = &book
	fmt.Println("book=>", book)
	fmt.Printf("Book1 title=>%s bookPtrSubject:%s \n", book.title, bookPtr.subject)
}

func testSlice() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var numbers1 = make([]int, 10, 10)
	fmt.Println("numbers:", numbers)
	fmt.Println("numbers(1-3)", numbers[1:4])

	fmt.Println("numbers1:", numbers1)
	fmt.Println("numbers1(0-9)", numbers1[0:9])

	fmt.Println("lennumbers,capnumbers=>", len(numbers), cap(numbers))

	numbers = append(numbers, 20, 21, 22)
	fmt.Println("new numbers", numbers)

}

func testRange() {
	var pow = []int{1, 2, 4, 6, 87, 8, 9, 23, 44, 333}
	for i, v := range pow {
		fmt.Printf("%d = %d\n", i, v)
	}
	map1 := make(map[int]float32)
	map1[0] = 1.0
	map1[1] = 2.0
	map1[2] = 3.0
	map1[3] = 4.0
	for key, value := range map1 {
		fmt.Printf("key:%d,value:%f\n", key, value)
	}
}

func testMap() {
	var mapCountry map[string]string
	mapCountry = make(map[string]string)

	mapCountry["China"] = "中国"
	mapCountry["France"] = "法国"
	for key := range mapCountry {
		fmt.Printf("key:%s value:%s \n", key, mapCountry[key])
	}

	delete(mapCountry, "France")
	fmt.Println("删除后的元素")
	for key := range mapCountry {
		fmt.Printf("key:%s value:%s \n", key, mapCountry[key])
	}
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
func testFiboacci() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}

type Phone interface {
	call()
}

type XiaomiPhone struct {
}

func (xiaomiPhone XiaomiPhone) call() {
	fmt.Println("小米手机")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("苹果手机")
}

func testInterface() {
	var phone Phone
	phone = new(XiaomiPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return 0, nil
}
func testError() {
	result, error := Sqrt(-1)
	fmt.Println(result, error)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func testThread() {
	go say("world")
	say("hello")
}

func testUnSafe() {
	num := 5
	numPointer := &num

	flnum := (*float32)(unsafe.Pointer(numPointer))

	fmt.Println("打印flnum的指针地址:")
	fmt.Println(flnum)
	fmt.Printf("%p", flnum)

	fmt.Println("打印flnum的指针的值:")
	fmt.Println(*flnum)
}

type Num struct {
	i string
	j int64
}

func testUnsafeOffset() {
	n := Num{i: "EDDYCJY", j: 1}
	nPointer := unsafe.Pointer(&n)

	niPointer := (*string)(unsafe.Pointer(nPointer))
	*niPointer = "煎鱼"

	njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))
	*njPointer = 2

	fmt.Printf("n.i: %s, n.j: %d", n.i, n.j)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, i := range s {
		sum += i
	}
	c <- sum
}

func testChannel() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func testPanic() {
	defer println("testPanic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("unknown error")
}


