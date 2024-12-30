package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is : %v\n", *p)
	fmt.Printf("The value of i is : %v\n", i)
	p = &i
	*p = 10
	fmt.Printf("The value p points to is : %v\n", *p)
	fmt.Printf("The value of i is : %v\n", i)

	var thing1 = [5]int32{1, 2, 3, 4, 5}
	fmt.Printf("The memory location of thing1 array is : %p\n", &thing1)
	//var result = square(&thing1)

}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("The memory location of thing2 array is : %p\n", &thing2)

	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return *thing2
}
