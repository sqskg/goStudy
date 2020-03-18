package hello

import "fmt"

func main()  {
	
	//fmt.Println(calculateBill(10,2))
	//var area,perimetr = rectProps(10,20)
	//fmt.Println("area = ",area,";perimeter = ",perimetr)
	//area1, _ := rectProps(10.8, 5.6) // 返回值周长被丢弃
	//fmt.Printf("Area %f ", area1)
	changingFunc(89,1,2,3,4,89);
}

func calculateBill(price int,no int) int {
	return price * no
}

func rectProps(length,width float64) (float64,float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area,perimeter
}

func changingFunc(num int,ele ...int){
	fmt.Printf("the type of ele is %T\n",ele);
	var found = false;
	for _,v := range ele{
		if(v == num){
			found = true;
			fmt.Println("found")
			break;
		}
	}
	if(!found){
		fmt.Println("not found")
	}

}