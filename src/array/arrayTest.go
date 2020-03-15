package main

import (
	"fmt"
)

func main()  {
	//var a [2]int
	//a[0] = 12;
	//a[1] = 13;
	//fmt.Println(a)
	//b := [3]int{12,12,22}
	//fmt.Println(b)
	//
	//c := [...]int{1,2,3,4,6,9}
	//fmt.Println(c)
	//
	//d := [...]string{"USA", "China", "India", "Germany", "France"}
	//e := d // a copy of a is assigned to b
	//e[0] = "Singapore"
	//fmt.Println("d is ", d)
	//fmt.Println("e is ", e)
	//num := [...]int{5, 6, 7, 8, 8}
	//fmt.Println("before passing to function ", num)
	//changeLocal(num) //num is passed by value
	////数组 num 实际上是通过值传递给函数 changeLocal
	//fmt.Println("after passing to function ", num)
	//range1()
	//slice();
	//valueTransfer()
	//multiArray();
	//makeSilce();
	//extendsCap();
	//nilTest();
	//sliceValueTest();
	//multiSlice();
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}


func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}

func range1()  {
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a {//range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a",sum)
}

func slice()  {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] // creates a slice from a[1] to a[3]
	fmt.Println(b)
}

//当数组作为参数传递给函数时，它们是按值传递，而原始数组保持不变
func valueTransfer()  {
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)
}

func multiArray()  {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, // this comma is necessary. The compiler will complain if you omit this comma
	}
	printArray(a)
	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printArray(b)
}
func printArray(a [3][2]string)  {
	for _, v1 := range a  {
		for _, v2 := range v1{
			fmt.Printf("%s ",v2)
		}
		fmt.Println("")
	}
}
func makeSilce(){
	i := make([]int,5,5)
	fmt.Println("len is ",len(i),"cap is ", cap(i))
	fmt.Println(i)
}

//如果往一个满了切片里添加数据，它会自动扩容到原理的2倍
func extendsCap(){
	car := []string{"Ferrari","Honda", "Ford"}
	fmt.Println("car's = ",car,"car's len = ", len(car),"car's cap = ", cap(car))
	car = append(car, "toyota")
	fmt.Println("car's = ",car,"car's len = ", len(car),"car's cap = ", cap(car))
}

//zero value of a slice is nil
func nilTest()  {
	var names []string;
	if names == nil {
		fmt.Println("names = ",names,"names len = ", len(names),"names cap = ", cap(names))
		names = append(names, "John","Sebastian");
		fmt.Println("names = ",names,"names len = ", len(names),"names cap = ", cap(names))
		names = append(names, "jordon")
		fmt.Println("names = ",names,"names len = ", len(names),"names cap = ", cap(names))
		names = append(names, "jordon")
		fmt.Println("names = ",names,"names len = ", len(names),"names cap = ", cap(names))
	}
}

func sunElement(number []int)  {
	for i, _ := range number{
		number[i] -= 2
	}
}
func sliceValueTest(){
	nos := []int{9,10,10}
	fmt.Println("nos = ",nos)
	sunElement(nos)
	fmt.Println("nos = ",nos)
}

func multiSlice()  {
	pls := [][]string{{"22","23"},{"000"},{"32","33"}}
	for _,v1 := range pls{
		for _,v2 := range v1 {
			fmt.Print(" ",v2)
		}
		fmt.Println("")
	}

}

func countries() []string{
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededContires := countries[:len(countries) -2];
	contiriescp := make([]string,len(neededContires))
	copy(contiriescp,neededContires)
   return contiriescp;
}