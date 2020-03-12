package main

import "fmt"

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
	slice();
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