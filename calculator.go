package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Basic Calculator")
	fmt.Println("================")
	fmt.Println("Enter expressions in the form: <number1> <operator> <number2>")
	fmt.Println("Supported operators: +, -, *, /")
	fmt.Println("Type 'exit' to quit the calculator.\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter expression: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Error: Invalid expression format. Use: <number1> <operator> <number2>")
			continue
		}

		num1, err1 := strconv.ParseFloat(parts[0], 64)
		operator := parts[1]
		num2, err2 := strconv.ParseFloat(parts[2], 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Error: Invalid numbers. Ensure you enter valid numeric values.")
			continue
		}

		switch operator {
		case "+":
			fmt.Printf("Result: %.2f\n", num1+num2)
		case "-":
			fmt.Printf("Result: %.2f\n", num1-num2)
		case "*":
			fmt.Printf("Result: %.2f\n", num1*num2)
		case "/":
			if num2 == 0 {
				fmt.Println("Error: Division by zero is not allowed.")
			} else {
				fmt.Printf("Result: %.2f\n", num1/num2)
			}
		default:
			fmt.Println("Error: Unsupported operator. Use +, -, *, or /.")
		}
	}
}
