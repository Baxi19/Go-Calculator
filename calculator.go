package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type calc struct {}

func (calc) operate(input string, operator string) int {
	cleanInput := strings.Split(input, operator)
	operator1 := parsear(cleanInput[0])
	operator2 := parsear(cleanInput[1])

	switch operator {
	case "+":
		fmt.Println(operator1 + operator2)
		return (operator1 + operator2)
	case "-":
		fmt.Println(operator1 - operator2)
		return (operator1 - operator2)
	case "*":
		fmt.Println(operator1 * operator2)
		return (operator1 * operator2)
	case "/":
		fmt.Println(operator1 / operator2)
		return (operator1 / operator2)
	default:
		log.Println(operator, "operation is not supported!")
		return 0
	}
}

func parsear(entrada string) int {
	operator, _ := strconv.Atoi(entrada)
	return operator
}

func readInput() string {
	//get the user input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	read := readInput()
	operator := readInput()
	c := calc{}
	fmt.Println(c.operate(read, operator))
}
