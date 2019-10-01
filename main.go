package main

import "fmt"

func main() {
	colors := map[string]string{
		"orange":        "#FFA500",
		"white":         "#FFFFFF",
		"red":           "#dsvbkjsd",
		"blue":          "#sdvnlksnd",
		"yellow":        "#ndsvjbsd",
		"green":         "#wdfywuf",
		"gray":          "#sbdise",
		"purple":        "#sfhowue",
		"peacock-green": "#nsslklwei",
	}
	fmt.Println("Before correction")
	fmt.Println(colors)

	//correctRGBValues(colors)
	fmt.Println("After Correction")

	fmt.Println(colors)
}

func main3() {
	myslice := []string{
		"A", "B", "C", "D", "E",
	}
	fmt.Println("Before update.")
	fmt.Println(myslice)

	updateSlice(myslice)

	fmt.Println("After update")
	fmt.Println(myslice)
}
func main2() {
	s := Student{firstName: "Sandip",
		std:       15,
		aggregate: 72.04,
		contactInfo: contactInfo{
			phone: "8753752385372",
			city:  "Pune",
			email: "sandip@calsoftinc.com",
		},
	}
	fmt.Println("\nBefore update\n")
	s.print()

	s.updateFirstName("Sandeep")

	fmt.Println("\nAfter update\n")
	s.print()
}

func main1() {

	fmt.Println(addition(100, 2, 4, 5, 7, 8))
	//card := "Ace of Diamonds"

	//cards := newDeck()
	/*	cards := newDeckByReadingFile("cards_set")

		cards.print()

	/*	fmt.Println(cards.combineCards())
		cards.writeIntoFile("cards_set")

		/*	d1, d2 := cards.cut(10)

		fmt.Println("Deck 1:")
		d1.print()
		fmt.Println("=============================================")
		fmt.Println("Deck 2:")
		d2.print()
	*/}

func addition(a int, b ...int) int {
	var sum int
	sum = sum + a
	for _, v := range b {
		sum = sum + v
	}
	return sum
}

func updateSlice(s []string) {
	s[1] = "Good Bye."
}
