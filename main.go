package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	const (
		timeRestraint = 10
	)

	// initialize global pseudo generator
	rand.Seed(time.Now().UnixNano())

	planets := planetSetup()
	goodPath := goodGuyPath()
	startingOptions := []int{1, 2, 3, 4, 5, 6, 7}

	// set starting positions
	goodLoc := goodPath[0]
	badLoc := randMove(startingOptions)

	printStatus(0, goodLoc, badLoc)

	if isBusted(goodLoc, badLoc) {
		fmt.Println("GOT EM!")
		return
	}

	for i := 1; i <= timeRestraint; i++ {
		badMoveOptions := planets[badLoc]
		badLoc = randMove(badMoveOptions)
		goodLoc = goodPath[i]

		printStatus(i, goodLoc, badLoc)

		if isBusted(goodLoc, badLoc) {
			fmt.Println("GOT 'EM! Good job, Ranger!")
			return
		}
	}

	fmt.Println("WOMP WOMP! Bad Guys got away! D:")
	return
}

func printStatus(hour int, good int, bad int) {
	fmt.Println("----> Hour ", hour)
	fmt.Println("Good Guy: ", good)
	fmt.Println("Bad Guy: ", bad)
}

func randMove(options []int) int {
	return options[rand.Intn(len(options))]
}

func goodGuyPath() []int {
	return []int{2, 3, 4, 3, 6, 2, 3, 4, 3, 6}
}

func planetSetup() map[int][]int {
	planets := make(map[int][]int, 7)
	planets[1] = []int{2}
	planets[2] = []int{1, 3}
	planets[3] = []int{2, 4, 6}
	planets[4] = []int{3, 5}
	planets[5] = []int{4}
	planets[6] = []int{3, 7}
	planets[7] = []int{6}
	return planets
}

func isBusted(good int, bad int) bool {
	return good == bad
}
