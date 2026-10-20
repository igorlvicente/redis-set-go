package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	statement, _ := reader.ReadString('\n')
	statementParser := regexp.MustCompile(`(?P<command>SET) (?P<key>[a-zA-Z]+) (?P<value>.*)`)
	match := statementParser.FindStringSubmatch(statement)

	parsedStatement := map[string]string{}
	for i, name := range match {
		if i != 0 && name != "" {
			parsedStatement[statementParser.SubexpNames()[i]] = name
		}
	}

	database := map[string]string{}

	if parsedStatement["command"] == "SET" {
		database[parsedStatement["key"]] = parsedStatement["value"]
	}

	fmt.Print(database)
}
