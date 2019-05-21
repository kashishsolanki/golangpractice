package main

import "fmt"

func main() {
	fmt.Println("Basic programme is here.")
	// fmt.Println(abc("Text"))
	fmt.Println(multipleReturn())

	// req, res = multipleReturn()

	fmt.Println("Data Types-------------------")
	var x = 5
	fmt.Println(x)

	var y int = 6
	z := "String Text"
	var p bool = true
	fmt.Println(y, z, p)

	fmt.Println("If/Else Example-------------------")
	if (y%2 == 0) {
		fmt.Println("Even", y%2)
	} else {
		fmt.Println("Odd")
	}

	fmt.Println("For loop Example-------------------1")
	i := 1
	for i <= 3 {
		fmt.Println("i value :: ", i)
		i += 1
	}

	fmt.Println("For loop Example-------------------2")
	for j := 0; j <= 3; j++ {
		fmt.Println("j value :: ", j)
	}

	fmt.Println("Arrays Example-------------------")
	var a [5]int
	fmt.Println("a : ", a)

	b := [5]int{1,2,3,4,5}
	fmt.Println("b : ", b)

	fmt.Println("Struct------------------")
	fmt.Println("Person : ", person{"Kashish", 24})
	fmt.Println("Person : ", person{name: "Kashish"})
	fmt.Println("Person : ", person{name:"Kashish", age:24})

	r := rect{weight: 10, height: 5}
	fmt.Println("r : ", r);

	fmt.Println("Area :: ", r.area())
	fmt.Println("Perim :: ", r.perim())

	rp := &r
	fmt.Println("Area :: ", rp.area())
	fmt.Println("Perim :: ", rp.perim())
}

type person struct {
	name string
	age int
}

type rect struct {
	weight, height int
}

func (r *rect) area() int {
	return r.weight * r.height
}

func (r rect) perim() int {
	return 2*r.weight + 2*r.height
}

func abc(str string) (string) {
	return str
}

func multipleReturn() (int, int) {
	return 6, 7
}