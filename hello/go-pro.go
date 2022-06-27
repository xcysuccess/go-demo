package main

import "fmt"

func main()  {
	fmt.Println("====testSlice====")
	testSlice()

	fmt.Println("====testString2Interface====")
	testString2Interface()

	fmt.Println("====testInterface2String====")
	testInterface2String()

	fmt.Println("====testInterface1====")
	testInterface1()

	fmt.Println("====testInterface2====")
	testInterface2()

	fmt.Println("====testInterface3====")
	testInterface3()
}



func testSlice()  {
	var ids = []int{1, 2, 3}
	printSlice(ids)
}

func printSlice(vals []int) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func testString2Interface()  {
	var itf interface{} = 1
	i, ok := itf.(string)
	println("值:", i, "; 断言结果", ok)

	j, ok := itf.(int)
	println("值:", j, "; 断言结果", ok)
	//---------------
	names := []string{"stanley", "david", "oscar"}
	vals := make([]interface{},len(names))
	for i, v:= range names{
		vals[i] = v
	}
	printAll(vals)
}

func printAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func testInterface2String()  {
	a := "tomxiang"
	v, ok := interface{}(a).(string)	// 将 a 转换为接口类型
	if ok {
		fmt.Printf("v type is %T\n", v)
	}
	fmt.Println(a)
}

func testInterface1()  {
	var ids1 = []interface{}{1, 2.2, 3.4}
	printInterface1(ids1)
}

func printInterface1(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func testInterface2()  {
	var ids2 = []interface{}{1, 2.2, 3.4}
	printInterface2(ids2)
}
func printInterface2(vals ...interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func testInterface3()  {
	var ids3 = []interface{}{1, 2.2, 3.4}
	printInterface3(ids3...)
}
func printInterface3(vals ...interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}