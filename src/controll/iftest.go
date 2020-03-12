package controll

import "fmt"

func main() {
	//forTest();
	//switchTest(8);
	switchTest2()
}

func ifTest() {
	num := 11
	if num%2 == 0 {
		fmt.Println("num is even")
	} else {
		fmt.Println("num is odd")
	}

}
func forTest() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func number() int {
	num := 15 * 5
	return num
}

//每执行完一个 case 后，会从 switch 语句中跳出来，不再做后续 case 的判断和执行。使用 fallthrough 语句可以在已经执行完成的 case 之后，把控制权转移到下一个 case 的执行代码中
func switchTest2() {
	switch num := number(); { // num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
}
func switchTest(finger int) {
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default: // 默认情况
		fmt.Println("incorrect finger number")
	}
}
