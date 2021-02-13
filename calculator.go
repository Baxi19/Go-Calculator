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
	operator1, err1 := strconv.Atoi(values[0])
	operator2, err2 := strconv.Atoi(values[1])

	//check if exist any error
	if err1 != nil {
		panic(err1)
	} else if err2 != nil {
		panic(err2)
	} else {
		fmt.Println(operator1 + operator2)
	}

}
