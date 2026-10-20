package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	database := map[string]string{}
	reader := bufio.NewReader(os.Stdin)
	key, _ := reader.ReadString('\n')
	value, _ := reader.ReadString('\n')
	database[key] = value
	fmt.Print(database)
}
