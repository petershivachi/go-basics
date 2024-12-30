package main

import "fmt"

func main() {
	/* Arrays
	- Fixed length,
	- Stores elements of the same type
	- Indexable
	- Contiguous in memory
	*/

	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[0:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	//Array initialization
	var intArr1 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr1)

	//initialize using the colon shorthand
	intArr2 := [3]int32{1, 2, 3}
	fmt.Println(intArr2)

	//initialize using the ... syntax
	intArr3 := [...]int32{1, 2, 3}
	fmt.Println(intArr3)

	/*
		* SLICES
		- Slices wrap arrays to give more general, powerful and convenient interface to sequence of data
		- Similar to arrays but don't include the length value
	*/
	var intSlice []int32 = []int32{1, 2, 3}
	fmt.Printf("The length of the slice is %d with a capacity of %d \n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length of the slice is %d with a capacity of %d \n", len(intSlice), cap(intSlice))
	var intSlice1 []int32 = []int32{1, 2}
	intSlice = append(intSlice, intSlice1...)
	fmt.Println(intSlice)
	var intSlice2 = make([]int32, 3, 4)
	fmt.Println(intSlice2)

	/*
		* MAPS
		- Maps store key value pairs
	*/
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 28, "Jane": 30, "Stephen": 25, "Smith": 18}
	var age, ok = myMap2["Stephen"]

	if ok {
		fmt.Printf("Age is %d", age)
	} else {
		fmt.Println("Invalid name")
	}

	for key, value := range myMap2 {
		fmt.Printf("key:%s value:%d \n", key, value)
	}

}
