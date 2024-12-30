package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
* Features of channels
- Hold data
- Are thread safe
- Listen for data
*/

var MaxChickenPrice float32 = 5
var MaxTofuPrice float32 = 3

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}

	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice > MaxChickenPrice {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice > MaxTofuPrice {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:

		fmt.Printf("\nText Sent: Found deal on chicken at : %v ", website)
	case website := <-tofuChannel:
		fmt.Printf("\nEmail Sent: Found a deal on tofu at : %v ", website)
	}

}
