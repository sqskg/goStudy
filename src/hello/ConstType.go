package main

import (
	"fmt"
	"math"
)

// go 的常量类型用 const 标注
func main() {

	var a int = 50
	//const a int = 50
	a = 60
	fmt.Println(a)

	var b = math.Sqrt(4)
	//const b1 = math.Sqrt(8);

	fmt.Println(b)

	const hello string = "hello world"

	fmt.Println(hello)

	const trueConst = true
	type myBool bool
	var defaultBool = trueConst       // 允许
	var customBool myBool = trueConst // 允许
	fmt.Println(defaultBool)
	fmt.Println(customBool)
	//defaultBool = customBool // 不允许

	const s = 5
	var intVar int = s
	var int32Var int32 = s
	var float64Var float64 = s
	var complex64Var complex64 = s
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
	{
		var i = 5
		var f = 5.6
		var c = 5 + 6i
		fmt.Printf("i's type %T, f's type %T, c's type %T", i, f, c)
	}
}
