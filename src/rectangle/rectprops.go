package rectangle

import "math"
import "fmt"

func init()  {
	fmt.Println("init rectprops pacage")
}
/**
任何以大写字母开头的变量或者函数都是被导出的名字，其他包只能访问被导出的函数或变量
 */
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}

/**
首字母小写不能被其他包调用
 */
func area() int {
	return 1;
}
