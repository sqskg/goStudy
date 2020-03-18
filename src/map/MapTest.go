package main

import (
	"fmt"
)

func main() {
	//initMap();
	//getValue();
	//exist()
	//forRange()
	//removeKey()
	//length();
	//valueReferenc()
	equesTes()
}

// value reference ,change one value then effect origin value
func valueReferenc() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Original person salary", personSalary)
}

// delete key from map,use delete[map,key] method
func removeKey() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("map before deletion ", personSalary)
	delete(personSalary, "steve")
	fmt.Println("map after deletion ", personSalary)
}

//遍历 map， 用for range
func forRange() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("All items of a map")

	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d \n", key, value)
	}

}

//check key is exit or not，use value，ok
func exist() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	employee := "sqs"
	value, ok := personSalary[employee]
	if ok == true {
		fmt.Println("salary of ", employee, "is 【", value, "】")
	} else {
		fmt.Println(employee, "'s salary not found")
	}
}

// get value
func getValue() {
	personSalary := map[string]int{
		"streve": 120000,
		"jamie":  15000,
	}
	personSalary["jhone"] = 14900
	var employee = "jamie"
	fmt.Println("salary of ", employee, "is [", personSalary[employee], "]")
}

// init map，use make
func initMap() {
	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("map is null, Go to make one")
		//personSarlary is nil ,so need make method to init it ;progress will output 'map is null, Go to make one'
		personSalary = make(map[string]int)
	}
	personSalary["steve"] = 12000
	personSalary["jone"] = 15000
	personSalary["mike"] = 8000
	fmt.Println("personSalary map contents:", personSalary)
}

// get length use len[map] method
func length() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("length is", len(personSalary))
}

// map 的相等性；map之间不能使用==操作符判断，==只能用来检查map是否为nil
//判断两个 map 是否相等的方法是遍历比较两个 map 中的每个元素。
func equesTes() {
	map1 := map[string]int{
		"one": 1,
		"two": 2,
		"four":4,
	}

	map2 := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}

	var equals = true
	if len(map1) != len(map2) {
		equals = false
	} else {
		for key, value1 := range map1 {
			value, ok := map2[key]
			if ok {
				if value == value1 {
					//equals
				} else {
					equals = false
					fmt.Println(value1,"!=",value)
					break
				}
			} else {
				fmt.Println("map2 not exits ",key)
				equals = false
				break
			}
		}
	}
	fmt.Println(equals)
}
