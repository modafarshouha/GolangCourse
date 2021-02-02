package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name, address string
	infoMap := make(map[string]string)

	fmt.Println("Please enter a NAME :")
	fmt.Scanln(&name)
	infoMap["name"] = name

	fmt.Println("Please enter an ADDRESS :")
	inAdd := bufio.NewReader(os.Stdin)
	address, _ = inAdd.ReadString('\n')
	address = strings.TrimSuffix(address, "\r\n")
	infoMap["address"] = address

	jsonB, _ := json.Marshal(infoMap)
	printJSON := strings.Trim(string(jsonB), "{}")
	fmt.Println(printJSON)
}
