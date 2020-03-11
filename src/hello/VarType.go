package main

import "fmt"

func main()  {
	//bool 类型表示一个布尔值，值为 true 或者 false。
	a := true;
	b := false;
	fmt.Println("a = ",a,"b =",b);
	c := a && b;
	fmt.Println("c = ",c);
	d := a || b;
	fmt.Println("d = ",d);
}
