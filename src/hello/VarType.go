package main

import "fmt"
import "unsafe"

func main()  {
	//boolTest();
	//intTest();
	//floatTest()
	//complexText()
	//strintTest()
	typeConvert()
}



func boolTest()  {
	//bool 类型表示一个布尔值，值为 true 或者 false。
	a := true;
	b := false;
	fmt.Println("a = ",a,"b =",b);
	c := a && b;
	fmt.Println("c = ",c);
	d := a || b;
	fmt.Println("d = ",d);
}
func intTest()  {
	var a int = 1;
	var b = 189;
	fmt.Println("a = ",a ,"b = ",b)
	fmt.Printf("type of a is %T,size of a is %d",a,unsafe.Sizeof(a))
	fmt.Printf("\n type of b is %T,size of b is %d",b,unsafe.Sizeof(b))
}

func floatTest()  {
	a ,b := 5.67,8.97
	fmt.Printf("type of a %T b%T \n",a,b)
	sum := a + b;
	diff := b - a
	fmt.Println("sum = ",sum,"b - a = ",diff)

	no1 ,no2 := 56,89
	fmt.Println("sum = ",no1+no2,"diff = ",no1-no2)

	var1,var2 := 0.03,0.97;
	var3 := 1.0;
	diff = var3 - var2;
	fmt.Println("var1 + var2 = ", var1 + var2,"var3 - var2 = ",diff)
	var4 := 1.11111
	fmt.Println(var4)
	fmt.Println(10.0/3)
	var5,var6 := 0.3,0.1
	fmt.Println(var5 - var6)
}

func complexText()  {
	c1 := complex(5,7)
	c2 := 8 + 27i
	cadd := c1 + c2;
	product := c1 * c2;
	fmt.Println("cadd = ",cadd)
	fmt.Println("product = ",product)
}

func strintTest()  {
	fist := "shi"
	last := "qingshan"
	fmt.Println(fist + last)
}

func typeConvert()  {
	j := 100
	i := 10.1
	fmt.Println(float64(j) + i)
	
}