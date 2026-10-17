package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputFromTerminal, _ := reader.ReadString('\n')
	fmt.Print(inputFromTerminal)
}
