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
	//get the user input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operador := "-"
	operation := scanner.Text()

	// using split, get values as list
	values := strings.Split(operation, operador)

	//Cast String as Int
	operator1, err1 := strconv.Atoi(values[0])
	operator2, err2 := strconv.Atoi(values[1])

	//check if exist any error
	if err1 != nil {
		//panic(err1)
		log.Println(err1.Error())
	} else if err2 != nil {
		//panic(err2)
		log.Println(err2.Error())
	} else {

		switch operador {
		case "+":
			fmt.Println(operator1 + operator2)
		case "-":
			fmt.Println(operator1 - operator2)
		case "*":
			fmt.Println(operator1 * operator2)
		case "/":
			fmt.Println(operator1 / operator2)
		default:
			log.Println(operador, "operation is not supported!")
		}

	}

}
