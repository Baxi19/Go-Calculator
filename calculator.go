package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//get the user input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operation := scanner.Text()

	// using split, get values as list
	values := strings.Split(operation, "+")

	//Cast String as Int
	operator1, _ := strconv.Atoi(values[0])
	operator2, _ := strconv.Atoi(values[1])

	fmt.Println(operator1 + operator2)
}
