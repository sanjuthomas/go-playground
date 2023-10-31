package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	saneName := strings.Trim(name, "\n")

	fmt.Print("Enter your address: ")
	address, _ := reader.ReadString('\n')
	saneAddress := strings.Trim(address, "\n")

	var data map[string]string
	data = make(map[string]string)

	data["name"] = saneName
	data["address"] = saneAddress
	jsonData, err := json.Marshal(data)
	if err == nil {
		fmt.Println(string(jsonData))
	}
}
