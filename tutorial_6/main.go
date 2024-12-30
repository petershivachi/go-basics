package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it")
	} else {
		fmt.Println("You need to fuel ASAP")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15}
	canMakeIt(myEngine, 10)
	fmt.Printf("Miles Per Gallon : %d Gallons : %d\n", myEngine.mpg, myEngine.gallons)
	fmt.Printf("Total miles left in tank %d ", myEngine.milesLeft())

}
