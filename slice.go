package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var currentIndex int = 0
	var length int = 3
	var slice = make([]int, length)
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter an integer: ")
		text, _ := reader.ReadString('\n')
		saneInput := strings.Trim(text, "\n")
		//if input is X exit rigt away
		if saneInput == "X" {
			fmt.Println("Thank you. Have a good one!")
			os.Exit(3)
		}
		//try to convert to integer
		i, err := strconv.Atoi(saneInput)
		if err != nil {
			fmt.Println("invalid input, please enter an integer. To exit please press enter X")
			continue
		}
		//good integer
		if currentIndex < length {
			slice[currentIndex] = i
			if currentIndex == 1 {
				if slice[0] > slice[1] {
					var tmp int = slice[0]
					slice[0] = slice[1]
					slice[1] = tmp
				}
			} else if currentIndex == 2 {
				iSort(slice)
			}
			currentIndex += 1
		} else {
			slice = append(slice, i)
			iSort(slice)
			currentIndex += 1
		}
		iPrint(currentIndex, slice)
	}
}

func iPrint(currentIndex int, slice []int) {
	if currentIndex < len(slice) {
		fmt.Println(slice[0:currentIndex])
	} else {
		fmt.Println(slice)
	}
}

func iSort(slice []int) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}
