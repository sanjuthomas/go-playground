package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type NameTye struct {
	fname string
	lname string
}

func main() {

	slice := make([]NameTye, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the abosolute file name: ")
	name, _ := reader.ReadString('\n')
	fileName := strings.Trim(name, "\n")

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		saneLine := strings.Trim(line, "\n")
		names := strings.Fields(saneLine)
		s := NameTye{names[0], names[1]}
		slice = append(slice, s)
		if err == io.EOF {
			break
		}

	}
	for _, n := range slice {
		fmt.Printf("%s %s\n", n.fname, n.lname)
	}
}
