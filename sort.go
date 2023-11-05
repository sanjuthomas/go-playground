package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter 10 numbers (eg: 5, 3, 1) : ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		log.Fatalf("failed to read from stdin")
	}
	cleanInput := strings.TrimSuffix(input, "\n")
	numbers := strings.Split(cleanInput, ",")
	var data = make([]int, 0, 10)
	for _, number := range numbers {
		i, _ := strconv.Atoi(strings.Trim(number, " "))
		data = append(data, i)
	}

	if len(data) > 10 {
		fmt.Printf("You entered %d number, please enter only 10 number\n", len(data))
		os.Exit(3)
	}

	for _, number := range BubbleSort(data) {
		fmt.Printf("%d ", number)
	}
	fmt.Println("\n")
}

func BubbleSort(data []int) []int {
	for i := len(data); i > 0; i-- {
		for j := 1; j < i; j++ {
			if data[j-1] > data[j] {
				Swap(data, j)
			}
		}
	}
	return data
}

func Swap(slice []int, j int) {
	tmp := slice[j]
	slice[j] = slice[j-1]
	slice[j-1] = tmp
}
